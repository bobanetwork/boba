// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
)

const L2GovernanceERC20StorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"},{\"astId\":1005,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_nonces\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_address,t_struct(Counter)1013_storage)\"},{\"astId\":1006,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_PERMIT_TYPEHASH_DEPRECATED_SLOT\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_bytes32\"},{\"astId\":1007,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_delegates\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":1008,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_checkpoints\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_mapping(t_address,t_array(t_struct(Checkpoint)1012_storage)dyn_storage)\"},{\"astId\":1009,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"_totalSupplyCheckpoints\",\"offset\":0,\"slot\":\"9\",\"type\":\"t_array(t_struct(Checkpoint)1012_storage)dyn_storage\"},{\"astId\":1010,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"l1Token\",\"offset\":0,\"slot\":\"10\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/boba/L2GovernanceERC20.sol:L2GovernanceERC20\",\"label\":\"l2Bridge\",\"offset\":0,\"slot\":\"11\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(Checkpoint)1012_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct ERC20Votes.Checkpoint[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(Checkpoint)1012_storage\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_mapping(t_address,t_array(t_struct(Checkpoint)1012_storage)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct ERC20Votes.Checkpoint[])\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_array(t_struct(Checkpoint)1012_storage)dyn_storage\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_struct(Counter)1013_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct Counters.Counter)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(Counter)1013_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(Checkpoint)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ERC20Votes.Checkpoint\",\"numberOfBytes\":\"32\"},\"t_struct(Counter)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Counters.Counter\",\"numberOfBytes\":\"32\"},\"t_uint224\":{\"encoding\":\"inplace\",\"label\":\"uint224\",\"numberOfBytes\":\"28\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"}}}"

var L2GovernanceERC20StorageLayout = new(solc.StorageLayout)

var L2GovernanceERC20DeployedBin = "0x608060405234801561001057600080fd5b50600436106101e55760003560e01c80637ecebe001161010f578063b4b5ea57116100a2578063d505accf11610071578063d505accf146104ba578063d5abeb01146104cd578063dd62ed3e1461050d578063f1127ed81461055357600080fd5b8063b4b5ea571461046c578063c01e1bd61461047f578063c3cda5201461049f578063ce665dd8146104b257600080fd5b80639dc29fac116100de5780639dc29fac14610413578063a457c2d714610426578063a9059cbb14610439578063ae1f6aaf1461044c57600080fd5b80637ecebe00146103d25780638e539e8c146103e557806395d89b41146103f85780639ab24eb01461040057600080fd5b806339509351116101875780635c19a95c116101565780635c19a95c146103315780636fcfff451461034457806370a082311461036c578063782d6fe1146103a257600080fd5b806339509351146102985780633a46b1a8146102ab57806340c10f19146102be578063587cde1e146102d357600080fd5b806318160ddd116101c357806318160ddd1461023a57806323b872dd1461024c578063313ce5671461025f5780633644e5151461029057600080fd5b806301ffc9a7146101ea57806306fdde0314610212578063095ea7b314610227575b600080fd5b6101fd6101f83660046126c2565b6105a5565b60405190151581526020015b60405180910390f35b61021a610665565b6040516102099190612704565b6101fd6102353660046127a0565b6106f7565b6002545b604051908152602001610209565b6101fd61025a3660046127ca565b61070f565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610209565b61023e610733565b6101fd6102a63660046127a0565b610742565b61023e6102b93660046127a0565b61078e565b6102d16102cc3660046127a0565b610834565b005b61030c6102e1366004612806565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600760205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610209565b6102d161033f366004612806565b610913565b610357610352366004612806565b610920565b60405163ffffffff9091168152602001610209565b61023e61037a366004612806565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6103b56103b03660046127a0565b610955565b6040516bffffffffffffffffffffffff9091168152602001610209565b61023e6103e0366004612806565b610969565b61023e6103f3366004612821565b610994565b61021a610a0a565b61023e61040e366004612806565b610a19565b6102d16104213660046127a0565b610ae4565b6101fd6104343660046127a0565b610bb7565b6101fd6104473660046127a0565b610c88565b600b5461030c9073ffffffffffffffffffffffffffffffffffffffff1681565b6103b561047a366004612806565b610c96565b600a5461030c9073ffffffffffffffffffffffffffffffffffffffff1681565b6102d16104ad36600461284b565b610ca4565b61023e600081565b6102d16104c83660046128a3565b610e1b565b6104e06b019d971e4fe8401e7400000081565b6040517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9091168152602001610209565b61023e61051b36600461290d565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b610566610561366004612940565b610fda565b60408051825163ffffffff1681526020928301517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169281019290925201610209565b60007f01ffc9a7a5cef8baa21ed3c5c0d7e23accb804b619e9333b597f47a0d84076e27f1d1d8b63000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000084167f01ffc9a700000000000000000000000000000000000000000000000000000000148061065d57507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b949350505050565b60606003805461067490612980565b80601f01602080910402602001604051908101604052809291908181526020018280546106a090612980565b80156106ed5780601f106106c2576101008083540402835291602001916106ed565b820191906000526020600020905b8154815290600101906020018083116106d057829003601f168201915b5050505050905090565b600033610705818585611080565b5060019392505050565b60003361071d858285611233565b61072885858561130a565b506001949350505050565b600061073d6115c3565b905090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061070590829086906107899087906129fc565b611080565b60004382106107fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4552433230566f7465733a20626c6f636b206e6f7420796574206d696e65640060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260086020526040902061082d90836116f7565b9392505050565b600b5473ffffffffffffffffffffffffffffffffffffffff1633146108b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f6e6c79204c32204272696467652063616e206d696e7420616e64206275726e60448201526064016107f5565b6108bf82826117eb565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161090791815260200190565b60405180910390a25050565b61091d33826117f9565b50565b73ffffffffffffffffffffffffffffffffffffffff811660009081526008602052604081205461094f90611897565b92915050565b600061082d610964848461078e565b611931565b73ffffffffffffffffffffffffffffffffffffffff811660009081526005602052604081205461094f565b60004382106109ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4552433230566f7465733a20626c6f636b206e6f7420796574206d696e65640060448201526064016107f5565b61094f6009836116f7565b60606004805461067490612980565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600860205260408120548015610abc5773ffffffffffffffffffffffffffffffffffffffff83166000908152600860205260409020610a75600183612a14565b81548110610a8557610a85612a2b565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610abf565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b600b5473ffffffffffffffffffffffffffffffffffffffff163314610b65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f6e6c79204c32204272696467652063616e206d696e7420616e64206275726e60448201526064016107f5565b610b6f82826119cf565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161090791815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610c7b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016107f5565b6107288286868403611080565b60003361070581858561130a565b600061094f61096483610a19565b83421115610d0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4552433230566f7465733a207369676e6174757265206578706972656400000060448201526064016107f5565b604080517fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf602082015273ffffffffffffffffffffffffffffffffffffffff8816918101919091526060810186905260808101859052600090610d9590610d8d9060a001604051602081830303815290604052805190602001206119d9565b858585611a42565b9050610da081611a6a565b8614610e08576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433230566f7465733a20696e76616c6964206e6f6e63650000000000000060448201526064016107f5565b610e1281886117f9565b50505050505050565b83421115610e85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016107f5565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610eb48c611a6a565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610f1c826119d9565b90506000610f2c82878787611a42565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610fc3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016107f5565b610fce8a8a8a611080565b50505050505050505050565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff83166000908152600860205260409020805463ffffffff841690811061102b5761102b612a2b565b60009182526020918290206040805180820190915291015463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16918101919091529392505050565b73ffffffffffffffffffffffffffffffffffffffff8316611122576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff82166111c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461130457818110156112f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016107f5565b6113048484848403611080565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166113ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff8216611450576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015611506576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff80851660009081526020819052604080822085850390559185168152908120805484929061154a9084906129fc565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516115b091815260200190565b60405180910390a3611304848484611aa4565b60003073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561162957507f000000000000000000000000000000000000000000000000000000000000000046145b1561165357507f000000000000000000000000000000000000000000000000000000000000000090565b50604080517f00000000000000000000000000000000000000000000000000000000000000006020808301919091527f0000000000000000000000000000000000000000000000000000000000000000828401527f000000000000000000000000000000000000000000000000000000000000000060608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b600061170381836129fc565b835490925060005b8181101561176857600061171f8284611aaf565b90508486828154811061173457611734612a2b565b60009182526020909120015463ffffffff16111561175457809250611762565b61175f8160016129fc565b91505b5061170b565b81156117c1578461177a600184612a14565b8154811061178a5761178a612a2b565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166117c4565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1695945050505050565b6117f58282611aca565b5050565b73ffffffffffffffffffffffffffffffffffffffff8281166000818152600760208181526040808420805485845282862054949093528787167fffffffffffffffffffffffff00000000000000000000000000000000000000008416811790915590519190951694919391928592917f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9190a4611304828483611b80565b600063ffffffff82111561192d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f322062697473000000000000000000000000000000000000000000000000000060648201526084016107f5565b5090565b60006bffffffffffffffffffffffff82111561192d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203960448201527f362062697473000000000000000000000000000000000000000000000000000060648201526084016107f5565b6117f58282611d25565b600061094f6119e66115c3565b836040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b6000806000611a5387878787611d3d565b91509150611a6081611e55565b5095945050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526005602052604090208054600181018255905b50919050565b505050565b611a9f8383836120a9565b6000611abe6002848418612a5a565b61082d908484166129fc565b611ad482826120e8565b6002546b019d971e4fe8401e740000001015611b72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433230566f7465733a20746f74616c20737570706c79207269736b73206f60448201527f766572666c6f77696e6720766f7465730000000000000000000000000000000060648201526084016107f5565b61130460096122108361221c565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015611bbc5750600081115b15611a9f5773ffffffffffffffffffffffffffffffffffffffff831615611c715773ffffffffffffffffffffffffffffffffffffffff831660009081526008602052604081208190611c11906124148561221c565b915091508473ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7248383604051611c66929190918252602082015260400190565b60405180910390a250505b73ffffffffffffffffffffffffffffffffffffffff821615611a9f5773ffffffffffffffffffffffffffffffffffffffff821660009081526008602052604081208190611cc1906122108561221c565b915091508373ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7248383604051611d16929190918252602082015260400190565b60405180910390a25050505050565b611d2f8282612420565b61130460096124148361221c565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611d745750600090506003611e4c565b8460ff16601b14158015611d8c57508460ff16601c14155b15611d9d5750600090506004611e4c565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611df1573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611e4557600060019250925050611e4c565b9150600090505b94509492505050565b6000816004811115611e6957611e69612a95565b03611e715750565b6001816004811115611e8557611e85612a95565b03611eec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016107f5565b6002816004811115611f0057611f00612a95565b03611f67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016107f5565b6003816004811115611f7b57611f7b612a95565b03612008576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016107f5565b600481600481111561201c5761201c612a95565b0361091d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff838116600090815260076020526040808220548584168352912054611a9f92918216911683611b80565b73ffffffffffffffffffffffffffffffffffffffff8216612165576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016107f5565b806002600082825461217791906129fc565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040812080548392906121b19084906129fc565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a36117f560008383611aa4565b600061082d82846129fc565b82546000908190801561227c5785612235600183612a14565b8154811061224557612245612a2b565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661227f565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1692506122ad83858763ffffffff16565b91506000811180156122f557506122c56000436129fc565b866122d1600184612a14565b815481106122e1576122e1612a2b565b60009182526020909120015463ffffffff16145b1561237f5761230382612614565b8661230f600184612a14565b8154811061231f5761231f612a2b565b9060005260206000200160000160046101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16021790555061240b565b8560405180604001604052806123a060004361239b91906129fc565b611897565b63ffffffff1681526020016123b485612614565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b50935093915050565b600061082d8284612a14565b73ffffffffffffffffffffffffffffffffffffffff82166124c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015612579576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016107f5565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604081208383039055600280548492906125b5908490612a14565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3611a9f83600084611aa4565b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff82111561192d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f323420626974730000000000000000000000000000000000000000000000000060648201526084016107f5565b6000602082840312156126d457600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461082d57600080fd5b600060208083528351808285015260005b8181101561273157858101830151858201604001528201612715565b81811115612743576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461279b57600080fd5b919050565b600080604083850312156127b357600080fd5b6127bc83612777565b946020939093013593505050565b6000806000606084860312156127df57600080fd5b6127e884612777565b92506127f660208501612777565b9150604084013590509250925092565b60006020828403121561281857600080fd5b61082d82612777565b60006020828403121561283357600080fd5b5035919050565b803560ff8116811461279b57600080fd5b60008060008060008060c0878903121561286457600080fd5b61286d87612777565b955060208701359450604087013593506128896060880161283a565b92506080870135915060a087013590509295509295509295565b600080600080600080600060e0888a0312156128be57600080fd5b6128c788612777565b96506128d560208901612777565b955060408801359450606088013593506128f16080890161283a565b925060a0880135915060c0880135905092959891949750929550565b6000806040838503121561292057600080fd5b61292983612777565b915061293760208401612777565b90509250929050565b6000806040838503121561295357600080fd5b61295c83612777565b9150602083013563ffffffff8116811461297557600080fd5b809150509250929050565b600181811c9082168061299457607f821691505b602082108103611a99577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115612a0f57612a0f6129cd565b500190565b600082821015612a2657612a266129cd565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082612a90577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2GovernanceERC20StorageLayoutJSON), L2GovernanceERC20StorageLayout); err != nil {
		panic(err)
	}

	layouts["L2GovernanceERC20"] = L2GovernanceERC20StorageLayout
	deployedBytecodes["L2GovernanceERC20"] = L2GovernanceERC20DeployedBin
}
