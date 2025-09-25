import json
import os
import requests
from typing import Dict, Any

def routescan_l2_transaction(event: Dict[str, Any], context: Any) -> Dict[str, Any]:
    """
    Lambda handler for getting L2 transaction data from Routescan API.

    Expected POST body:
    {
        "chainId": "56288",  # Chain ID for the boba bnb network
        "address": "0x...",  # Address to query
        "action": "txlist",  # Optional, defaults to txlist
        "startBlock": "41658388"  # Optional, starting block number
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

        # Construct API URL
        base_url = os.environ['ROUTESCAN_API_URL']
        api_key_token = os.environ['ROUTESCAN_API_KEY']
        api_url = f"{base_url}/mainnet/evm/{chain_id}/etherscan/api"
        # Make request to Routescan API
        params = {
            'module': 'account',
            'action': action,
            'address': address,
            'startBlock': start_block,
            'tag': 'latest',
            'apikey': api_key_token
        }

        response = requests.get(api_url, params=params)
        response.raise_for_status()  # Raise exception for non-200 status codes

        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*',
            },
            'body': response.text
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
