// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
)

const L1FeeVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L1FeeVault.sol:L1FeeVault\",\"label\":\"totalProcessed\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"}],\"types\":{\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L1FeeVaultStorageLayout = new(solc.StorageLayout)

var L1FeeVaultDeployedBin = "0x6080604052600436106100695760003560e01c806384411d651161004357806384411d651461010c578063d0e12f9014610130578063d3e5792b1461017157600080fd5b80630d9019e1146100755780633ccfd60b146100d357806354fd4d50146100ea57600080fd5b3661007057005b600080fd5b34801561008157600080fd5b506100a97f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100df57600080fd5b506100e86101a5565b005b3480156100f657600080fd5b506100ff610591565b6040516100ca91906107eb565b34801561011857600080fd5b5061012260005481565b6040519081526020016100ca565b34801561013c57600080fd5b506101647f000000000000000000000000000000000000000000000000000000000000000081565b6040516100ca919061086f565b34801561017d57600080fd5b506101227f000000000000000000000000000000000000000000000000000000000000000081565b7f0000000000000000000000000000000000000000000000000000000000000000471015610280576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b60004790508060008082825461029691906108b2565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a17f38e04cbeb8c10f8f568618aa75be0f10b6729b8b4237743b4de20cbcde2839ee817f0000000000000000000000000000000000000000000000000000000000000000337f000000000000000000000000000000000000000000000000000000000000000060405161038494939291906108ca565b60405180910390a160017f000000000000000000000000000000000000000000000000000000000000000060018111156103c0576103c0610805565b036104d95760007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d806000811461043f576040519150601f19603f3d011682016040523d82523d6000602084013e610444565b606091505b50509050806104d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4665655661756c743a206661696c656420746f2073656e642045544820746f2060448201527f4c322066656520726563697069656e74000000000000000000000000000000006064820152608401610277565b5050565b604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000109163e11013dd91849161055c917f0000000000000000000000000000000000000000000000000000000000000000916188b89160040161090b565b6000604051808303818588803b15801561057557600080fd5b505af1158015610589573d6000803e3d6000fd5b505050505050565b60606105bc7f0000000000000000000000000000000000000000000000000000000000000000610634565b6105e57f0000000000000000000000000000000000000000000000000000000000000000610634565b61060e7f0000000000000000000000000000000000000000000000000000000000000000610634565b60405160200161062093929190610946565b604051602081830303815290604052905090565b60608160000361067757505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156106a1578061068b816109bc565b915061069a9050600a83610a23565b915061067b565b60008167ffffffffffffffff8111156106bc576106bc610a37565b6040519080825280601f01601f1916602001820160405280156106e6576020820181803683370190505b5090505b8415610769576106fb600183610a66565b9150610708600a86610a7d565b6107139060306108b2565b60f81b81838151811061072857610728610a91565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610762600a86610a23565b94506106ea565b949350505050565b60005b8381101561078c578181015183820152602001610774565b8381111561079b576000848401525b50505050565b600081518084526107b9816020860160208601610771565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006107fe60208301846107a1565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811061086b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b6020810161087d8284610834565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156108c5576108c5610883565b500190565b84815273ffffffffffffffffffffffffffffffffffffffff848116602083015283166040820152608081016109026060830184610834565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff8316602082015260606040820152600061090260608301846107a1565b60008451610958818460208901610771565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610994816001850160208a01610771565b600192019182015283516109af816002840160208801610771565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036109ed576109ed610883565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610a3257610a326109f4565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610a7857610a78610883565b500390565b600082610a8c57610a8c6109f4565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1FeeVaultStorageLayoutJSON), L1FeeVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1FeeVault"] = L1FeeVaultStorageLayout
	deployedBytecodes["L1FeeVault"] = L1FeeVaultDeployedBin
}
