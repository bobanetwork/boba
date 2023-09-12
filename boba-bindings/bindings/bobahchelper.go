// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ledgerwatch/erigon"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/accounts/abi"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = libcommon.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BobaHCHelperABI is the input ABI used to generate the binding from.
const BobaHCHelperABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"EndpointRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"EndpointUnregistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AddCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"AddPermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"CallOffchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"CheckPermittedCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Endpoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetLegacyResponse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_random\",\"type\":\"uint256\"}],\"name\":\"GetRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_auth\",\"type\":\"bytes32\"}],\"name\":\"RegisterEndpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"RemovePermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"session\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"myNum\",\"type\":\"uint256\"}],\"name\":\"SequentialRandom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"UnregisterEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"legacyWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prepaidCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BobaHCHelperBin is the compiled bytecode used for deploying new contracts.
var BobaHCHelperBin = "0x6101006040527342000000000000000000000000000000000000fd60805273deaddeaddeaddeaddeaddeaddeaddeaddead990160a05273420000000000000000000000000000000000002360c05273420000000000000000000000000000000000002060e05234801561007157600080fd5b50600180546001600160a01b0319163317905560805160a05160c05160e051612bb36100e060003960008181610a1301528181610a510152818161155301526115910152600081816103fa01528181610e920152610f85015260006118a801526000610efd0152612bb36000f3fe608060405234801561001057600080fd5b50600436106101355760003560e01c806387fc1b25116100b2578063c8b58d1511610081578063e2f9479211610066578063e2f947921461031c578063eb6598b51461032f578063f2fde38b1461034257600080fd5b8063c8b58d15146102d9578063d40c48b0146102fc57600080fd5b806387fc1b2514610227578063997e31cc1461023a578063a71adb3e14610295578063c1fd7e46146102b857600080fd5b80633042902711610109578063493d57d6116100ee578063493d57d6146101e157806362fccf6e146101f45780637e1214131461020757600080fd5b8063304290271461019e57806332be428f146101b157600080fd5b80622925261461013a5780630ceff204146101565780631730f39b1461016b578063248eaf621461017e575b600080fd5b61014360005481565b6040519081526020015b60405180910390f35b610169610164366004612172565b610355565b005b6101696101793660046121fd565b610426565b61014361018c366004612251565b60036020526000908152604090205481565b6101696101ac3660046121fd565b6105ec565b6101c46101bf36600461226c565b6107ad565b60408051931515845260208401929092529082015260600161014d565b6101436101ef3660046122ac565b610900565b6101696102023660046122d6565b610c35565b610143610215366004612251565b60026020526000908152604090205481565b610169610235366004612318565b610e0e565b610270610248366004612172565b60056020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161014d565b6102a86102a3366004612334565b610ef8565b604051901515815260200161014d565b6102cb6102c6366004612380565b61118f565b60405161014d929190612462565b6102a86102e7366004612251565b60046020526000908152604090205460ff1681565b61030f61030a36600461257f565b61138d565b60405161014d919061261e565b6102a861032a3660046121fd565b611779565b61016961033d36600461263f565b611890565b610169610350366004612251565b611a11565b6000548111156103c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c696420416d6f756e7400000000000000000000000000000000000060448201526064015b60405180910390fd5b806000808282546103d791906126be565b90915550506001546104239073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911683611a78565b50565b6000838360405160200161043b9291906126d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915073ffffffffffffffffffffffffffffffffffffffff166104fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016103bd565b60008181526005602052604090205473ffffffffffffffffffffffffffffffffffffffff16331461058a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016103bd565b600090815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff90941683526001938401909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790555050565b600083836040516020016106019291906126d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915073ffffffffffffffffffffffffffffffffffffffff166106c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016103bd565b60008181526005602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610750576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016103bd565b600090815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff9094168352600190930190522080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555050565b6040517f32be428f00000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166024830152603882018590526058820184905260788201839052600091829182916332be428f9183918290819081906098016040516020818303038152906040528051906020012090506108548133601e611b51565b90955093508415610879578380602001905181019061087391906126e5565b90935091505b848015610887575083516040145b6108ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c65640060448201526064016103bd565b50929a9099509197509095505050505050565b60008263ffffffff16600114610972576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f545552494e473a204765746820696e74657263657074206661696c757265000060448201526064016103bd565b6040517f493d57d600000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015263493d57d691600091908290603801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050601e7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1615610b915760007f0000000000000000000000000000000000000000000000000000000000000000905060008173ffffffffffffffffffffffffffffffffffffffff1663e24dfcde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610abf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae39190612709565b6040517f0a8c07b00000000000000000000000000000000000000000000000000000000081523360048201526024810182905290915060009073ffffffffffffffffffffffffffffffffffffffff841690630a8c07b0906044016020604051808303816000875af1158015610b5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b809190612722565b90508015610b8d57600093505b5050505b610b9c823383611b51565b9094509250838015610baf575082516020145b610c15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f48433a2047657452616e646f6d206661696c757265000000000000000000000060448201526064016103bd565b82806020019051810190610c299190612709565b98975050505050505050565b60008282604051602001610c4a9291906126d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915073ffffffffffffffffffffffffffffffffffffffff16610d0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016103bd565b60008181526005602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610d99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016103bd565b6000818152600560205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055517fddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a90610e0190859085903390612788565b60405180910390a1505050565b60008111610e78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c696420616d6f756e7400000000000000000000000000000000000060448201526064016103bd565b610eba73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084611dc4565b73ffffffffffffffffffffffffffffffffffffffff821660009081526002602052604081208054839290610eef9084906127c2565b90915550505050565b6000807f0000000000000000000000000000000000000000000000000000000000000000905060008585604051602001610f339291906126d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050610fae73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633306064611dc4565b6064600080828254610fc091906127c2565b90915550506040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f61f334d30000000000000000000000000000000000000000000000000000000017905291517fc1fd7e46000000000000000000000000000000000000000000000000000000008152909160609160009173ffffffffffffffffffffffffffffffffffffffff87169163c1fd7e4691611077918d918d918991016127da565b6000604051808303816000875af1158015611096573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526110dc919081019061280a565b925090508080156110ee575081516020145b80156110ff57508151602083012087145b1561117d576000848152600560205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000163390811790915590517fd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f88291611170918c918c91612788565b60405180910390a1611181565b5060005b9450505050505b9392505050565b60006060600063c1fd7e469050600087876040516020016111b19291906126d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012090503330146113015760008181526005602052604090205473ffffffffffffffffffffffffffffffffffffffff1661127b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f456e64706f696e74206e6f74207265676973746572656400000000000000000060448201526064016103bd565b600081815260056020908152604080832033845260010190915290205460ff16611301576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207065726d697474656400000000000000000060448201526064016103bd565b600082338a8a6040516020016113189291906126d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526113599392918b908b90602001612897565b60405160208183030381529060405280519060200120905061137d81336032611b51565b9450945050505094509492505050565b60608363ffffffff16600114806113ad57508363ffffffff166302000001145b611413576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f545552494e473a204765746820696e74657263657074206661696c757265000060448201526064016103bd565b3360009081526004602052604090205463d40c48b09060ff1680611435575060015b6114c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4c65676163792063616c6c6572206973206e6f7420696e2077686974656c697360448201527f740000000000000000000000000000000000000000000000000000000000000060648201526084016103bd565b60008133866040516020016114d6919061291d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526115159392918890602001612939565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050600060607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16156117535760007f0000000000000000000000000000000000000000000000000000000000000000905060008173ffffffffffffffffffffffffffffffffffffffff1663e24dfcde6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116239190612709565b6040517f0a8c07b00000000000000000000000000000000000000000000000000000000081523360048201526024810182905290915060009073ffffffffffffffffffffffffffffffffffffffff841690630a8c07b0906044016020604051808303816000875af115801561169c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116c09190612722565b90508061174f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f496e73756666696369656e742063726564697420696e206c656761637920426f60448201527f6261547572696e6743726564697420636f6e747261637400000000000000000060648201526084016103bd565b5050505b61175f83336000611b51565b90925090508161176e57600080fd5b979650505050505050565b600080848460405160200161178f9291906126d5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915073ffffffffffffffffffffffffffffffffffffffff16611851576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016103bd565b600090815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616845260010190915290205460ff1690509392505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461192f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c65720000000060448201526064016103bd565b60008481526006602052604090208054611948906129c6565b1590506119b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c7265616479206578697374730000000000000000000060448201526064016103bd565b8282826040516020016119c693929190612a19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152600086815260066020522090611a0a9082612a8c565b5050505050565b73ffffffffffffffffffffffffffffffffffffffff8116611a3157600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052611b4c9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611e28565b505050565b6000606073ffffffffffffffffffffffffffffffffffffffff84163014611c885773ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812054611ba290856127c2565b73ffffffffffffffffffffffffffffffffffffffff8616600090815260026020526040902054909150811115611c34576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e73756666696369656e74206372656469740000000000000000000000000060448201526064016103bd565b73ffffffffffffffffffffffffffffffffffffffff851660009081526002602052604081208054869290611c699084906126be565b9250508190555083600080828254611c8191906127c2565b9091555050505b60008581526006602052604081208054611ca1906129c6565b80601f0160208091040260200160405190810160405280929190818152602001828054611ccd906129c6565b8015611d1a5780601f10611cef57610100808354040283529160200191611d1a565b820191906000526020600020905b815481529060010190602001808311611cfd57829003601f168201915b505050505090506000815111611d8c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f48433a204d697373696e6720636163686520656e74727900000000000000000060448201526064016103bd565b6000868152600660205260408120611da391612124565b80806020019051810190611db7919061280a565b9250925050935093915050565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052611e229085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611aca565b50505050565b6000611e8a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611f349092919063ffffffff16565b805190915015611b4c5780806020019051810190611ea89190612722565b611b4c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016103bd565b6060611f438484600085611f4b565b949350505050565b606082471015611fdd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016103bd565b73ffffffffffffffffffffffffffffffffffffffff85163b61205b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103bd565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612084919061291d565b60006040518083038185875af1925050503d80600081146120c1576040519150601f19603f3d011682016040523d82523d6000602084013e6120c6565b606091505b509150915061176e828286606083156120e0575081611188565b8251156120f05782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bd919061261e565b508054612130906129c6565b6000825580601f10612140575050565b601f01602090049060005260206000209081019061042391905b8082111561216e576000815560010161215a565b5090565b60006020828403121561218457600080fd5b5035919050565b60008083601f84011261219d57600080fd5b50813567ffffffffffffffff8111156121b557600080fd5b6020830191508360208285010111156121cd57600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff811681146121f857600080fd5b919050565b60008060006040848603121561221257600080fd5b833567ffffffffffffffff81111561222957600080fd5b6122358682870161218b565b90945092506122489050602085016121d4565b90509250925092565b60006020828403121561226357600080fd5b611188826121d4565b60008060006060848603121561228157600080fd5b505081359360208301359350604090920135919050565b803563ffffffff811681146121f857600080fd5b600080604083850312156122bf57600080fd5b6122c883612298565b946020939093013593505050565b600080602083850312156122e957600080fd5b823567ffffffffffffffff81111561230057600080fd5b61230c8582860161218b565b90969095509350505050565b6000806040838503121561232b57600080fd5b6122c8836121d4565b60008060006040848603121561234957600080fd5b833567ffffffffffffffff81111561236057600080fd5b61236c8682870161218b565b909790965060209590950135949350505050565b6000806000806040858703121561239657600080fd5b843567ffffffffffffffff808211156123ae57600080fd5b6123ba8883890161218b565b909650945060208701359150808211156123d357600080fd5b506123e08782880161218b565b95989497509550505050565b60005b838110156124075781810151838201526020016123ef565b83811115611e225750506000910152565b600081518084526124308160208601602086016123ec565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8215158152604060208201526000611f436040830184612418565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156124f3576124f361247d565b604052919050565b600067ffffffffffffffff8211156125155761251561247d565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600061255461254f846124fb565b6124ac565b905082815283838301111561256857600080fd5b828260208301376000602084830101529392505050565b60008060006060848603121561259457600080fd5b61259d84612298565b9250602084013567ffffffffffffffff808211156125ba57600080fd5b818601915086601f8301126125ce57600080fd5b6125dd87833560208501612541565b935060408601359150808211156125f357600080fd5b508401601f8101861361260557600080fd5b61261486823560208401612541565b9150509250925092565b6020815260006111886020830184612418565b801515811461042357600080fd5b6000806000806060858703121561265557600080fd5b84359350602085013561266781612631565b9250604085013567ffffffffffffffff81111561268357600080fd5b6123e08782880161218b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156126d0576126d061268f565b500390565b8183823760009101908152919050565b600080604083850312156126f857600080fd5b505080516020909101519092909150565b60006020828403121561271b57600080fd5b5051919050565b60006020828403121561273457600080fd5b815161118881612631565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60408152600061279c60408301858761273f565b905073ffffffffffffffffffffffffffffffffffffffff83166020830152949350505050565b600082198211156127d5576127d561268f565b500190565b6040815260006127ee60408301858761273f565b82810360208401526128008185612418565b9695505050505050565b6000806040838503121561281d57600080fd5b825161282881612631565b602084015190925067ffffffffffffffff81111561284557600080fd5b8301601f8101851361285657600080fd5b805161286461254f826124fb565b81815286602083850101111561287957600080fd5b61288a8260208301602086016123ec565b8093505050509250929050565b7fffffffff000000000000000000000000000000000000000000000000000000008660e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b166004820152600084516128ff8160188501602089016123ec565b82018385601883013760009301601801928352509095945050505050565b6000825161292f8184602087016123ec565b9190910192915050565b7fffffffff000000000000000000000000000000000000000000000000000000008560e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008460601b166004820152600083516129a18160188501602088016123ec565b8351908301906129b88160188401602088016123ec565b016018019695505050505050565b600181811c908216806129da57607f821691505b602082108103612a13577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8315158152604060208201526000612a3560408301848661273f565b95945050505050565b601f821115611b4c57600081815260208120601f850160051c81016020861015612a655750805b601f850160051c820191505b81811015612a8457828155600101612a71565b505050505050565b815167ffffffffffffffff811115612aa657612aa661247d565b612aba81612ab484546129c6565b84612a3e565b602080601f831160018114612b0d5760008415612ad75750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555612a84565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015612b5a57888601518255948401946001909101908401612b3b565b5085821015612b9657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a"

// DeployBobaHCHelper deploys a new Ethereum contract, binding an instance of BobaHCHelper to it.
func DeployBobaHCHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (libcommon.Address, types.Transaction, *BobaHCHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(BobaHCHelperABI))
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, libcommon.FromHex(BobaHCHelperBin), backend)
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}
	return address, tx, &BobaHCHelper{BobaHCHelperCaller: BobaHCHelperCaller{contract: contract}, BobaHCHelperTransactor: BobaHCHelperTransactor{contract: contract}, BobaHCHelperFilterer: BobaHCHelperFilterer{contract: contract}}, nil
}

// BobaHCHelper is an auto generated Go binding around an Ethereum contract.
type BobaHCHelper struct {
	BobaHCHelperCaller     // Read-only binding to the contract
	BobaHCHelperTransactor // Write-only binding to the contract
	BobaHCHelperFilterer   // Log filterer for contract events
}

// BobaHCHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type BobaHCHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BobaHCHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BobaHCHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BobaHCHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BobaHCHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BobaHCHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BobaHCHelperSession struct {
	Contract     *BobaHCHelper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BobaHCHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BobaHCHelperCallerSession struct {
	Contract *BobaHCHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BobaHCHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BobaHCHelperTransactorSession struct {
	Contract     *BobaHCHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BobaHCHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type BobaHCHelperRaw struct {
	Contract *BobaHCHelper // Generic contract binding to access the raw methods on
}

// BobaHCHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BobaHCHelperCallerRaw struct {
	Contract *BobaHCHelperCaller // Generic read-only contract binding to access the raw methods on
}

// BobaHCHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BobaHCHelperTransactorRaw struct {
	Contract *BobaHCHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBobaHCHelper creates a new instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelper(address libcommon.Address, backend bind.ContractBackend) (*BobaHCHelper, error) {
	contract, err := bindBobaHCHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelper{BobaHCHelperCaller: BobaHCHelperCaller{contract: contract}, BobaHCHelperTransactor: BobaHCHelperTransactor{contract: contract}, BobaHCHelperFilterer: BobaHCHelperFilterer{contract: contract}}, nil
}

// NewBobaHCHelperCaller creates a new read-only instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelperCaller(address libcommon.Address, caller bind.ContractCaller) (*BobaHCHelperCaller, error) {
	contract, err := bindBobaHCHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperCaller{contract: contract}, nil
}

// NewBobaHCHelperTransactor creates a new write-only instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelperTransactor(address libcommon.Address, transactor bind.ContractTransactor) (*BobaHCHelperTransactor, error) {
	contract, err := bindBobaHCHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperTransactor{contract: contract}, nil
}

// NewBobaHCHelperFilterer creates a new log filterer instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelperFilterer(address libcommon.Address, filterer bind.ContractFilterer) (*BobaHCHelperFilterer, error) {
	contract, err := bindBobaHCHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperFilterer{contract: contract}, nil
}

// bindBobaHCHelper binds a generic wrapper to an already deployed contract.
func bindBobaHCHelper(address libcommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BobaHCHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BobaHCHelper *BobaHCHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BobaHCHelper.Contract.BobaHCHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BobaHCHelper *BobaHCHelperRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _BobaHCHelper.Contract.BobaHCHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BobaHCHelper *BobaHCHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _BobaHCHelper.Contract.BobaHCHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BobaHCHelper *BobaHCHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BobaHCHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BobaHCHelper *BobaHCHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _BobaHCHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BobaHCHelper *BobaHCHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _BobaHCHelper.Contract.contract.Transact(opts, method, params...)
}

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xe2f94792.
//
// Solidity: function CheckPermittedCaller(string _url, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCaller) CheckPermittedCaller(opts *bind.CallOpts, _url string, _callerAddress libcommon.Address) (bool, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "CheckPermittedCaller", _url, _callerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xe2f94792.
//
// Solidity: function CheckPermittedCaller(string _url, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperSession) CheckPermittedCaller(_url string, _callerAddress libcommon.Address) (bool, error) {
	return _BobaHCHelper.Contract.CheckPermittedCaller(&_BobaHCHelper.CallOpts, _url, _callerAddress)
}

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xe2f94792.
//
// Solidity: function CheckPermittedCaller(string _url, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCallerSession) CheckPermittedCaller(_url string, _callerAddress libcommon.Address) (bool, error) {
	return _BobaHCHelper.Contract.CheckPermittedCaller(&_BobaHCHelper.CallOpts, _url, _callerAddress)
}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner)
func (_BobaHCHelper *BobaHCHelperCaller) Endpoints(opts *bind.CallOpts, arg0 [32]byte) (libcommon.Address, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "Endpoints", arg0)

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner)
func (_BobaHCHelper *BobaHCHelperSession) Endpoints(arg0 [32]byte) (libcommon.Address, error) {
	return _BobaHCHelper.Contract.Endpoints(&_BobaHCHelper.CallOpts, arg0)
}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner)
func (_BobaHCHelper *BobaHCHelperCallerSession) Endpoints(arg0 [32]byte) (libcommon.Address, error) {
	return _BobaHCHelper.Contract.Endpoints(&_BobaHCHelper.CallOpts, arg0)
}

// LegacyWhitelist is a free data retrieval call binding the contract method 0xc8b58d15.
//
// Solidity: function legacyWhitelist(address ) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCaller) LegacyWhitelist(opts *bind.CallOpts, arg0 libcommon.Address) (bool, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "legacyWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LegacyWhitelist is a free data retrieval call binding the contract method 0xc8b58d15.
//
// Solidity: function legacyWhitelist(address ) view returns(bool)
func (_BobaHCHelper *BobaHCHelperSession) LegacyWhitelist(arg0 libcommon.Address) (bool, error) {
	return _BobaHCHelper.Contract.LegacyWhitelist(&_BobaHCHelper.CallOpts, arg0)
}

// LegacyWhitelist is a free data retrieval call binding the contract method 0xc8b58d15.
//
// Solidity: function legacyWhitelist(address ) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCallerSession) LegacyWhitelist(arg0 libcommon.Address) (bool, error) {
	return _BobaHCHelper.Contract.LegacyWhitelist(&_BobaHCHelper.CallOpts, arg0)
}

// OwnerRevenue is a free data retrieval call binding the contract method 0x00292526.
//
// Solidity: function ownerRevenue() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCaller) OwnerRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "ownerRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerRevenue is a free data retrieval call binding the contract method 0x00292526.
//
// Solidity: function ownerRevenue() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) OwnerRevenue() (*big.Int, error) {
	return _BobaHCHelper.Contract.OwnerRevenue(&_BobaHCHelper.CallOpts)
}

// OwnerRevenue is a free data retrieval call binding the contract method 0x00292526.
//
// Solidity: function ownerRevenue() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCallerSession) OwnerRevenue() (*big.Int, error) {
	return _BobaHCHelper.Contract.OwnerRevenue(&_BobaHCHelper.CallOpts)
}

// PendingCharge is a free data retrieval call binding the contract method 0x248eaf62.
//
// Solidity: function pendingCharge(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCaller) PendingCharge(opts *bind.CallOpts, arg0 libcommon.Address) (*big.Int, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "pendingCharge", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCharge is a free data retrieval call binding the contract method 0x248eaf62.
//
// Solidity: function pendingCharge(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) PendingCharge(arg0 libcommon.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PendingCharge(&_BobaHCHelper.CallOpts, arg0)
}

// PendingCharge is a free data retrieval call binding the contract method 0x248eaf62.
//
// Solidity: function pendingCharge(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCallerSession) PendingCharge(arg0 libcommon.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PendingCharge(&_BobaHCHelper.CallOpts, arg0)
}

// PrepaidCredit is a free data retrieval call binding the contract method 0x7e121413.
//
// Solidity: function prepaidCredit(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCaller) PrepaidCredit(opts *bind.CallOpts, arg0 libcommon.Address) (*big.Int, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "prepaidCredit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrepaidCredit is a free data retrieval call binding the contract method 0x7e121413.
//
// Solidity: function prepaidCredit(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) PrepaidCredit(arg0 libcommon.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PrepaidCredit(&_BobaHCHelper.CallOpts, arg0)
}

// PrepaidCredit is a free data retrieval call binding the contract method 0x7e121413.
//
// Solidity: function prepaidCredit(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCallerSession) PrepaidCredit(arg0 libcommon.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PrepaidCredit(&_BobaHCHelper.CallOpts, arg0)
}

// AddCredit is a paid mutator transaction binding the contract method 0x87fc1b25.
//
// Solidity: function AddCredit(address _caller, uint256 _amount) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) AddCredit(opts *bind.TransactOpts, _caller libcommon.Address, _amount *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "AddCredit", _caller, _amount)
}

// AddCredit is a paid mutator transaction binding the contract method 0x87fc1b25.
//
// Solidity: function AddCredit(address _caller, uint256 _amount) returns()
func (_BobaHCHelper *BobaHCHelperSession) AddCredit(_caller libcommon.Address, _amount *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.AddCredit(&_BobaHCHelper.TransactOpts, _caller, _amount)
}

// AddCredit is a paid mutator transaction binding the contract method 0x87fc1b25.
//
// Solidity: function AddCredit(address _caller, uint256 _amount) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) AddCredit(_caller libcommon.Address, _amount *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.AddCredit(&_BobaHCHelper.TransactOpts, _caller, _amount)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x1730f39b.
//
// Solidity: function AddPermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) AddPermittedCaller(opts *bind.TransactOpts, _url string, _callerAddress libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "AddPermittedCaller", _url, _callerAddress)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x1730f39b.
//
// Solidity: function AddPermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperSession) AddPermittedCaller(_url string, _callerAddress libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.Contract.AddPermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x1730f39b.
//
// Solidity: function AddPermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) AddPermittedCaller(_url string, _callerAddress libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.Contract.AddPermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// CallOffchain is a paid mutator transaction binding the contract method 0xc1fd7e46.
//
// Solidity: function CallOffchain(string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) CallOffchain(opts *bind.TransactOpts, _url string, _payload []byte) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "CallOffchain", _url, _payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0xc1fd7e46.
//
// Solidity: function CallOffchain(string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperSession) CallOffchain(_url string, _payload []byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.CallOffchain(&_BobaHCHelper.TransactOpts, _url, _payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0xc1fd7e46.
//
// Solidity: function CallOffchain(string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) CallOffchain(_url string, _payload []byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.CallOffchain(&_BobaHCHelper.TransactOpts, _url, _payload)
}

// GetLegacyResponse is a paid mutator transaction binding the contract method 0xd40c48b0.
//
// Solidity: function GetLegacyResponse(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) GetLegacyResponse(opts *bind.TransactOpts, rType uint32, _url string, _payload []byte) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetLegacyResponse", rType, _url, _payload)
}

// GetLegacyResponse is a paid mutator transaction binding the contract method 0xd40c48b0.
//
// Solidity: function GetLegacyResponse(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperSession) GetLegacyResponse(rType uint32, _url string, _payload []byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.GetLegacyResponse(&_BobaHCHelper.TransactOpts, rType, _url, _payload)
}

// GetLegacyResponse is a paid mutator transaction binding the contract method 0xd40c48b0.
//
// Solidity: function GetLegacyResponse(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetLegacyResponse(rType uint32, _url string, _payload []byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.GetLegacyResponse(&_BobaHCHelper.TransactOpts, rType, _url, _payload)
}

// GetRandom is a paid mutator transaction binding the contract method 0x493d57d6.
//
// Solidity: function GetRandom(uint32 rType, uint256 _random) returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactor) GetRandom(opts *bind.TransactOpts, rType uint32, _random *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetRandom", rType, _random)
}

// GetRandom is a paid mutator transaction binding the contract method 0x493d57d6.
//
// Solidity: function GetRandom(uint32 rType, uint256 _random) returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) GetRandom(rType uint32, _random *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.GetRandom(&_BobaHCHelper.TransactOpts, rType, _random)
}

// GetRandom is a paid mutator transaction binding the contract method 0x493d57d6.
//
// Solidity: function GetRandom(uint32 rType, uint256 _random) returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetRandom(rType uint32, _random *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.GetRandom(&_BobaHCHelper.TransactOpts, rType, _random)
}

// PutResponse is a paid mutator transaction binding the contract method 0xeb6598b5.
//
// Solidity: function PutResponse(bytes32 _cacheKey, bool _success, bytes _returndata) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) PutResponse(opts *bind.TransactOpts, _cacheKey [32]byte, _success bool, _returndata []byte) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "PutResponse", _cacheKey, _success, _returndata)
}

// PutResponse is a paid mutator transaction binding the contract method 0xeb6598b5.
//
// Solidity: function PutResponse(bytes32 _cacheKey, bool _success, bytes _returndata) returns()
func (_BobaHCHelper *BobaHCHelperSession) PutResponse(_cacheKey [32]byte, _success bool, _returndata []byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, _cacheKey, _success, _returndata)
}

// PutResponse is a paid mutator transaction binding the contract method 0xeb6598b5.
//
// Solidity: function PutResponse(bytes32 _cacheKey, bool _success, bytes _returndata) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) PutResponse(_cacheKey [32]byte, _success bool, _returndata []byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, _cacheKey, _success, _returndata)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0xa71adb3e.
//
// Solidity: function RegisterEndpoint(string _url, bytes32 _auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperTransactor) RegisterEndpoint(opts *bind.TransactOpts, _url string, _auth [32]byte) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "RegisterEndpoint", _url, _auth)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0xa71adb3e.
//
// Solidity: function RegisterEndpoint(string _url, bytes32 _auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperSession) RegisterEndpoint(_url string, _auth [32]byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.RegisterEndpoint(&_BobaHCHelper.TransactOpts, _url, _auth)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0xa71adb3e.
//
// Solidity: function RegisterEndpoint(string _url, bytes32 _auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperTransactorSession) RegisterEndpoint(_url string, _auth [32]byte) (types.Transaction, error) {
	return _BobaHCHelper.Contract.RegisterEndpoint(&_BobaHCHelper.TransactOpts, _url, _auth)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x30429027.
//
// Solidity: function RemovePermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) RemovePermittedCaller(opts *bind.TransactOpts, _url string, _callerAddress libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "RemovePermittedCaller", _url, _callerAddress)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x30429027.
//
// Solidity: function RemovePermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperSession) RemovePermittedCaller(_url string, _callerAddress libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.Contract.RemovePermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x30429027.
//
// Solidity: function RemovePermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) RemovePermittedCaller(_url string, _callerAddress libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.Contract.RemovePermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// SequentialRandom is a paid mutator transaction binding the contract method 0x32be428f.
//
// Solidity: function SequentialRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bool, bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperTransactor) SequentialRandom(opts *bind.TransactOpts, session [32]byte, nextHash [32]byte, myNum *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "SequentialRandom", session, nextHash, myNum)
}

// SequentialRandom is a paid mutator transaction binding the contract method 0x32be428f.
//
// Solidity: function SequentialRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bool, bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperSession) SequentialRandom(session [32]byte, nextHash [32]byte, myNum *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.SequentialRandom(&_BobaHCHelper.TransactOpts, session, nextHash, myNum)
}

// SequentialRandom is a paid mutator transaction binding the contract method 0x32be428f.
//
// Solidity: function SequentialRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bool, bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperTransactorSession) SequentialRandom(session [32]byte, nextHash [32]byte, myNum *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.SequentialRandom(&_BobaHCHelper.TransactOpts, session, nextHash, myNum)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x62fccf6e.
//
// Solidity: function UnregisterEndpoint(string _url) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) UnregisterEndpoint(opts *bind.TransactOpts, _url string) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "UnregisterEndpoint", _url)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x62fccf6e.
//
// Solidity: function UnregisterEndpoint(string _url) returns()
func (_BobaHCHelper *BobaHCHelperSession) UnregisterEndpoint(_url string) (types.Transaction, error) {
	return _BobaHCHelper.Contract.UnregisterEndpoint(&_BobaHCHelper.TransactOpts, _url)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x62fccf6e.
//
// Solidity: function UnregisterEndpoint(string _url) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) UnregisterEndpoint(_url string) (types.Transaction, error) {
	return _BobaHCHelper.Contract.UnregisterEndpoint(&_BobaHCHelper.TransactOpts, _url)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_BobaHCHelper *BobaHCHelperSession) TransferOwnership(_newOwner libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.Contract.TransferOwnership(&_BobaHCHelper.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) TransferOwnership(_newOwner libcommon.Address) (types.Transaction, error) {
	return _BobaHCHelper.Contract.TransferOwnership(&_BobaHCHelper.TransactOpts, _newOwner)
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x0ceff204.
//
// Solidity: function withdrawRevenue(uint256 _withdrawAmount) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) WithdrawRevenue(opts *bind.TransactOpts, _withdrawAmount *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "withdrawRevenue", _withdrawAmount)
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x0ceff204.
//
// Solidity: function withdrawRevenue(uint256 _withdrawAmount) returns()
func (_BobaHCHelper *BobaHCHelperSession) WithdrawRevenue(_withdrawAmount *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.WithdrawRevenue(&_BobaHCHelper.TransactOpts, _withdrawAmount)
}

// WithdrawRevenue is a paid mutator transaction binding the contract method 0x0ceff204.
//
// Solidity: function withdrawRevenue(uint256 _withdrawAmount) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) WithdrawRevenue(_withdrawAmount *big.Int) (types.Transaction, error) {
	return _BobaHCHelper.Contract.WithdrawRevenue(&_BobaHCHelper.TransactOpts, _withdrawAmount)
}

// BobaHCHelperEndpointRegisteredIterator is returned from FilterEndpointRegistered and is used to iterate over the raw logs and unpacked data for EndpointRegistered events raised by the BobaHCHelper contract.
type BobaHCHelperEndpointRegisteredIterator struct {
	Event *BobaHCHelperEndpointRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BobaHCHelperEndpointRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BobaHCHelperEndpointRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BobaHCHelperEndpointRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BobaHCHelperEndpointRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BobaHCHelperEndpointRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BobaHCHelperEndpointRegistered represents a EndpointRegistered event raised by the BobaHCHelper contract.
type BobaHCHelperEndpointRegistered struct {
	URL   string
	Owner libcommon.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEndpointRegistered is a free log retrieval operation binding the contract event 0xd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f882.
//
// Solidity: event EndpointRegistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterEndpointRegistered(opts *bind.FilterOpts) (*BobaHCHelperEndpointRegisteredIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "EndpointRegistered")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperEndpointRegisteredIterator{contract: _BobaHCHelper.contract, event: "EndpointRegistered", logs: logs, sub: sub}, nil
}

// WatchEndpointRegistered is a free log subscription operation binding the contract event 0xd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f882.
//
// Solidity: event EndpointRegistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) WatchEndpointRegistered(opts *bind.WatchOpts, sink chan<- *BobaHCHelperEndpointRegistered) (event.Subscription, error) {

	logs, sub, err := _BobaHCHelper.contract.WatchLogs(opts, "EndpointRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BobaHCHelperEndpointRegistered)
				if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEndpointRegistered is a log parse operation binding the contract event 0xd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f882.
//
// Solidity: event EndpointRegistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseEndpointRegistered(log types.Log) (*BobaHCHelperEndpointRegistered, error) {
	event := new(BobaHCHelperEndpointRegistered)
	if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BobaHCHelperEndpointUnregisteredIterator is returned from FilterEndpointUnregistered and is used to iterate over the raw logs and unpacked data for EndpointUnregistered events raised by the BobaHCHelper contract.
type BobaHCHelperEndpointUnregisteredIterator struct {
	Event *BobaHCHelperEndpointUnregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BobaHCHelperEndpointUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BobaHCHelperEndpointUnregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BobaHCHelperEndpointUnregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BobaHCHelperEndpointUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BobaHCHelperEndpointUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BobaHCHelperEndpointUnregistered represents a EndpointUnregistered event raised by the BobaHCHelper contract.
type BobaHCHelperEndpointUnregistered struct {
	URL   string
	Owner libcommon.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEndpointUnregistered is a free log retrieval operation binding the contract event 0xddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a.
//
// Solidity: event EndpointUnregistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterEndpointUnregistered(opts *bind.FilterOpts) (*BobaHCHelperEndpointUnregisteredIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "EndpointUnregistered")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperEndpointUnregisteredIterator{contract: _BobaHCHelper.contract, event: "EndpointUnregistered", logs: logs, sub: sub}, nil
}

// WatchEndpointUnregistered is a free log subscription operation binding the contract event 0xddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a.
//
// Solidity: event EndpointUnregistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) WatchEndpointUnregistered(opts *bind.WatchOpts, sink chan<- *BobaHCHelperEndpointUnregistered) (event.Subscription, error) {

	logs, sub, err := _BobaHCHelper.contract.WatchLogs(opts, "EndpointUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BobaHCHelperEndpointUnregistered)
				if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointUnregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEndpointUnregistered is a log parse operation binding the contract event 0xddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a.
//
// Solidity: event EndpointUnregistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseEndpointUnregistered(log types.Log) (*BobaHCHelperEndpointUnregistered, error) {
	event := new(BobaHCHelperEndpointUnregistered)
	if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
