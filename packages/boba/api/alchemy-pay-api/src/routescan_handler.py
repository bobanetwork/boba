import json
import os
import requests
from typing import Dict, Any, List

def filter_and_merge_transaction_data(txlist_data: List[Dict], cross_tx_data: List[Dict]) -> List[Dict]:
    """
    Filter transactions with methodId '0xf60ed84c' / 'payAndWithdraw' and merge with cross-chain data.
    Match transactions based on hash === srcTxHash.
    """
    print("\n=== Debug: Transaction Processing ===")
    print(f"Total transactions from txlist: {len(txlist_data)}")

    # First filter transactions with the specific methodId
    filtered_txs = [
        tx for tx in txlist_data
        if tx.get('methodId', '').lower() == '0xf60ed84c'
    ]
    print(f"Filtered transactions with methodId '0xf60ed84c': {len(filtered_txs)}")
    if filtered_txs:
        print("Sample filtered transaction:")
        print(json.dumps(filtered_txs[0], indent=2))

    print(f"\nTotal cross-chain transactions: {len(cross_tx_data)}")

    # Convert cross_tx_data to a dict for O(1) lookup
    cross_tx_map = {
        tx['srcTxHash'].lower(): tx
        for tx in cross_tx_data
        if 'srcTxHash' in tx
    }
    print(f"\nCross-chain transactions after mapping: {len(cross_tx_map)}")

    merged_txs = []
    for tx in filtered_txs:
        tx_hash = tx.get('hash', '').lower()
        cross_tx = cross_tx_map.get(tx_hash, {})

        if cross_tx:
            print(f"\nFound matching cross-chain tx for hash: {tx_hash}")

        merged_tx = {
            **tx,
            'status': cross_tx.get('status', 'unknown'),
            'messageNonce': cross_tx.get('data', {}).get('messageNonce'),
            'messageHash': cross_tx.get('data', {}).get('messageHash'),
            'dstTxHash': cross_tx.get('dstTxHash'),
            'dstBlockNumber': cross_tx.get('dstBlockNumber')
        }
        merged_txs.append(merged_tx)

    print(f"\nFinal merged transactions: {len(merged_txs)}")
    return merged_txs

def routescan_l2_transaction(event: Dict[str, Any], context: Any) -> Dict[str, Any]:
    """
    Lambda handler for getting L2 transaction data from Routescan API.

    Expected POST body:
    {
        "chainId": "56288",  # Chain ID for the boba bnb network
        "address": "0x...",  # Address to query
        "action": "txlist",  # Optional, defaults to txlist
        "startBlock": "41658388",  # Optional, starting block number
        "limit": 50  # Optional, defaults to 50
    }
    """
    try:
        # Parse request body
        body = json.loads(event.get('body', '{}'))

        # Validate required parameters
        if not body.get('chainId'):
            return {
                'statusCode': 400,
                'body': json.dumps({'error': 'chainId is required'})
            }

        if not body.get('address'):
            return {
                'statusCode': 400,
                'body': json.dumps({'error': 'address is required'})
            }

        # Get parameters with defaults
        chain_id = body['chainId']
        address = body['address']
        action = body.get('action', 'txlist')
        start_block = body.get('startBlock', '0')
        limit = body.get('limit', 50)

        # Get base URL and API key from environment
        base_url = os.environ['ROUTESCAN_API_URL']
        api_key_token = os.environ['ROUTESCAN_API_KEY']

        # 1. Get transaction list from etherscan-compatible API
        txlist_url = f"{base_url}/mainnet/evm/{chain_id}/etherscan/api"
        txlist_params = {
            'module': 'account',
            'action': action,
            'address': address,
            'startBlock': start_block,
            'tag': 'latest',
            'apikey': api_key_token,
        }
        txlist_response = requests.get(txlist_url, params=txlist_params)
        txlist_response.raise_for_status()
        txlist_data = json.loads(txlist_response.text)

        # 2. Get cross-chain transaction data
        cross_tx_url = f"{base_url}/mainnet/evm/cross-transactions/messages"
        cross_tx_params = {
            'chainId': chain_id,
            'srcChainIds': chain_id,
            'dstChainIds': '56',  # BSC mainnet
            'sort': 'desc',
            'limit': limit,
            'apikey': api_key_token
        }
        cross_tx_response = requests.get(cross_tx_url, params=cross_tx_params)
        cross_tx_response.raise_for_status()
        cross_tx_data = json.loads(cross_tx_response.text)

        # 3. Filter and merge the data
        if txlist_data.get('status') == '1' and txlist_data.get('result'):
            merged_txs = filter_and_merge_transaction_data(
                txlist_data['result'],
                cross_tx_data.get('items', [])
            )
            response_data = {
                'status': '1',
                'message': 'OK',
                'result': merged_txs
            }
        else:
            # If no transactions found, return original response
            response_data = txlist_data

        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*',
            },
            'body': json.dumps(response_data)
        }

    except requests.exceptions.RequestException as e:
        return {
            'statusCode': 500,
            'body': json.dumps({
                'error': 'Failed to fetch data from Routescan API',
                'details': str(e)
            })
        }
    except json.JSONDecodeError:
        return {
            'statusCode': 400,
            'body': json.dumps({'error': 'Invalid JSON in request body'})
        }
    except Exception as e:
        return {
            'statusCode': 500,
            'headers': {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            'body': json.dumps({
                'error': 'Internal server error',
                'details': str(e)
            })
        }
