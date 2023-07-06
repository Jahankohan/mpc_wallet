// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Forwarder

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ForwarderMetaData contains all meta data concerning the Forwarder contract.
var ForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"MetaTransactionExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"META_TRANSACTION_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"META_TRANSACTION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"calculatedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstLayerAbi\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first_layer_digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"functionSig\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrievedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"second_layer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b507f1db874843672c48c2b3f8b32abc9aedf3c4132bc5c1094d62797110a8f762d546002819055507fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6600381905550620000706200018060201b60201c565b60048190555030600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060800160405280605281526020016200201f60529139604051602001620000e3919062000206565b60405160208183030381529060405280519060200120600254600354600454600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516020016200013b9594939291906200029a565b604051602081830303815290604052600690816200015a919062000567565b5060066040516200016c9190620006e8565b604051809103902060018190555062000701565b6000804690508091505090565b600081519050919050565b600081905092915050565b60005b83811015620001c3578082015181840152602081019050620001a6565b60008484015250505050565b6000620001dc826200018d565b620001e8818562000198565b9350620001fa818560208601620001a3565b80840191505092915050565b6000620002148284620001cf565b915081905092915050565b6000819050919050565b62000234816200021f565b82525050565b6000819050919050565b6200024f816200023a565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002828262000255565b9050919050565b620002948162000275565b82525050565b600060a082019050620002b1600083018862000229565b620002c0602083018762000229565b620002cf604083018662000229565b620002de606083018562000244565b620002ed608083018462000289565b9695505050505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200037957607f821691505b6020821081036200038f576200038e62000331565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003f97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003ba565b620004058683620003ba565b95508019841693508086168417925050509392505050565b6000819050919050565b600062000448620004426200043c846200023a565b6200041d565b6200023a565b9050919050565b6000819050919050565b620004648362000427565b6200047c62000473826200044f565b848454620003c7565b825550505050565b600090565b6200049362000484565b620004a081848462000459565b505050565b5b81811015620004c857620004bc60008262000489565b600181019050620004a6565b5050565b601f8211156200051757620004e18162000395565b620004ec84620003aa565b81016020851015620004fc578190505b620005146200050b85620003aa565b830182620004a5565b50505b505050565b600082821c905092915050565b60006200053c600019846008026200051c565b1980831691505092915050565b600062000557838362000529565b9150826002028217905092915050565b6200057282620002f7565b67ffffffffffffffff8111156200058e576200058d62000302565b5b6200059a825462000360565b620005a7828285620004cc565b600060209050601f831160018114620005df5760008415620005ca578287015190505b620005d6858262000549565b86555062000646565b601f198416620005ef8662000395565b60005b828110156200061957848901518255600182019150602085019450602081019050620005f2565b8683101562000639578489015162000635601f89168262000529565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60008154620006688162000360565b6200067481866200064e565b94506001821660008114620006925760018114620006a857620006df565b60ff1983168652811515820286019350620006df565b620006b38562000395565b60005b83811015620006d757815481890152600182019150602081019050620006b6565b838801955050505b50505092915050565b6000620006f6828462000659565b915081905092915050565b61190e80620007116000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c8063aef61c53116100b8578063e5ea988b1161007c578063e5ea988b14610320578063eea8ebe31461033e578063f172a4ce1461035a578063f698da2514610378578063f6b4dfb414610396578063fa2c322d146103b457610137565b8063aef61c531461028a578063b76ed6e3146102a8578063bbb433c4146102c6578063c7977be7146102e4578063e1b11da41461030257610137565b806352a82b65116100ff57806352a82b65146101e45780637ecebe00146102025780639c96203b14610232578063a9cae65314610250578063adc879e91461026c57610137565b806313e481291461013c578063152b5c0f1461015a5780631b3012a3146101785780632d0335ab146101965780633408e470146101c6575b600080fd5b6101446103d2565b6040516101519190610c68565b60405180910390f35b6101626103f8565b60405161016f9190610c9c565b60405180910390f35b6101806103fe565b60405161018d9190610d47565b60405180910390f35b6101b060048036038101906101ab9190610da9565b61048c565b6040516101bd9190610def565b60405180910390f35b6101ce6104d4565b6040516101db9190610def565b60405180910390f35b6101ec6104e1565b6040516101f99190610c9c565b60405180910390f35b61021c60048036038101906102179190610da9565b6104e7565b6040516102299190610def565b60405180910390f35b61023a6104ff565b6040516102479190610c9c565b60405180910390f35b61026a60048036038101906102659190610fa4565b610505565b005b610274610849565b6040516102819190610def565b60405180910390f35b61029261084f565b60405161029f9190610c9c565b60405180910390f35b6102b0610855565b6040516102bd9190610d47565b60405180910390f35b6102ce6108e3565b6040516102db91906110a2565b60405180910390f35b6102ec6108ff565b6040516102f99190610c9c565b60405180910390f35b61030a610941565b60405161031791906110a2565b60405180910390f35b61032861095d565b6040516103359190610d47565b60405180910390f35b610358600480360381019061035391906110c4565b6109eb565b005b610362610bb3565b60405161036f9190610c9c565b60405180910390f35b610380610bb9565b60405161038d9190610c9c565b60405180910390f35b61039e610bbf565b6040516103ab9190610c68565b60405180910390f35b6103bc610be5565b6040516103c99190610c9c565b60405180910390f35b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b6006805461040b9061118a565b80601f01602080910402602001604051908101604052809291908181526020018280546104379061118a565b80156104845780601f1061045957610100808354040283529160200191610484565b820191906000526020600020905b81548152906001019060200180831161046757829003601f168201915b505050505081565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000804690508091505090565b60095481565b60006020528060005260406000206000915090505481565b600a5481565b6001546040518060800160405280604381526020016118446043913960405160200161053191906111f7565b604051602081830303815290604052805190602001206000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548887805190602001206040516020016105a2949392919061120e565b604051602081830303815290604052805190602001206040516020016105c99291906112c0565b6040516020818303038152906040528051906020012060098190555060016009548285856040516000815260200160405260405161060a9493929190611306565b6020604051602081039080840390855afa15801561062c573d6000803e3d6000fd5b50505060206040510351600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610706576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106fd90611397565b60405180910390fd5b6000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190610755906113e6565b919050555060008573ffffffffffffffffffffffffffffffffffffffff1685604051610781919061146a565b600060405180830381855af49150503d80600081146107bc576040519150601f19603f3d011682016040523d82523d6000602084013e6107c1565b606091505b5050905080610805576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fc906114cd565b60405180910390fd5b7f5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b873387604051610838939291906114ed565b60405180910390a150505050505050565b60045481565b600b5481565b600780546108629061118a565b80601f016020809104026020016040519081016040528092919081815260200182805461088e9061118a565b80156108db5780601f106108b0576101008083540402835291602001916108db565b820191906000526020600020905b8154815290600101906020018083116108be57829003601f168201915b505050505081565b6040518060800160405280604381526020016118446043913981565b6040518060800160405280605281526020016118876052913960405160200161092891906111f7565b6040516020818303038152906040528051906020012081565b6040518060800160405280605281526020016118876052913981565b6008805461096a9061118a565b80601f01602080910402602001604051908101604052809291908181526020018280546109969061118a565b80156109e35780601f106109b8576101008083540402835291602001916109e3565b820191906000526020600020905b8154815290600101906020018083116109c657829003601f168201915b505050505081565b8380519060200120600b8190555060405180608001604052806043815260200161184460439139604051602001610a2291906111f7565b604051602081830303815290604052805190602001206000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205486600b54604051602001610a8e949392919061120e565b60405160208183030381529060405260089081610aab91906116d7565b506008604051610abb919061182c565b6040518091039020600a81905550600154600a54604051602001610ae09291906112c0565b60405160208183030381529060405260079081610afd91906116d7565b506007604051610b0d919061182c565b6040518091039020600981905550600160095482858560405160008152602001604052604051610b409493929190611306565b6020604051602081039080840390855afa158015610b62573d6000803e3d6000fd5b50505060206040510351600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b60025481565b60015481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60405180608001604052806043815260200161184460439139604051602001610c0e91906111f7565b6040516020818303038152906040528051906020012081565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c5282610c27565b9050919050565b610c6281610c47565b82525050565b6000602082019050610c7d6000830184610c59565b92915050565b6000819050919050565b610c9681610c83565b82525050565b6000602082019050610cb16000830184610c8d565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610cf1578082015181840152602081019050610cd6565b60008484015250505050565b6000601f19601f8301169050919050565b6000610d1982610cb7565b610d238185610cc2565b9350610d33818560208601610cd3565b610d3c81610cfd565b840191505092915050565b60006020820190508181036000830152610d618184610d0e565b905092915050565b6000604051905090565b600080fd5b600080fd5b610d8681610c47565b8114610d9157600080fd5b50565b600081359050610da381610d7d565b92915050565b600060208284031215610dbf57610dbe610d73565b5b6000610dcd84828501610d94565b91505092915050565b6000819050919050565b610de981610dd6565b82525050565b6000602082019050610e046000830184610de0565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e4c82610cfd565b810181811067ffffffffffffffff82111715610e6b57610e6a610e14565b5b80604052505050565b6000610e7e610d69565b9050610e8a8282610e43565b919050565b600067ffffffffffffffff821115610eaa57610ea9610e14565b5b610eb382610cfd565b9050602081019050919050565b82818337600083830152505050565b6000610ee2610edd84610e8f565b610e74565b905082815260208101848484011115610efe57610efd610e0f565b5b610f09848285610ec0565b509392505050565b600082601f830112610f2657610f25610e0a565b5b8135610f36848260208601610ecf565b91505092915050565b610f4881610c83565b8114610f5357600080fd5b50565b600081359050610f6581610f3f565b92915050565b600060ff82169050919050565b610f8181610f6b565b8114610f8c57600080fd5b50565b600081359050610f9e81610f78565b92915050565b60008060008060008060c08789031215610fc157610fc0610d73565b5b6000610fcf89828a01610d94565b9650506020610fe089828a01610d94565b955050604087013567ffffffffffffffff81111561100157611000610d78565b5b61100d89828a01610f11565b945050606061101e89828a01610f56565b935050608061102f89828a01610f56565b92505060a061104089828a01610f8f565b9150509295509295509295565b600081519050919050565b600082825260208201905092915050565b60006110748261104d565b61107e8185611058565b935061108e818560208601610cd3565b61109781610cfd565b840191505092915050565b600060208201905081810360008301526110bc8184611069565b905092915050565b600080600080600060a086880312156110e0576110df610d73565b5b60006110ee88828901610d94565b955050602086013567ffffffffffffffff81111561110f5761110e610d78565b5b61111b88828901610f11565b945050604061112c88828901610f56565b935050606061113d88828901610f56565b925050608061114e88828901610f8f565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806111a257607f821691505b6020821081036111b5576111b461115b565b5b50919050565b600081905092915050565b60006111d18261104d565b6111db81856111bb565b93506111eb818560208601610cd3565b80840191505092915050565b600061120382846111c6565b915081905092915050565b60006080820190506112236000830187610c8d565b6112306020830186610de0565b61123d6040830185610c59565b61124a6060830184610c8d565b95945050505050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b60006112896002836111bb565b915061129482611253565b600282019050919050565b6000819050919050565b6112ba6112b582610c83565b61129f565b82525050565b60006112cb8261127c565b91506112d782856112a9565b6020820191506112e782846112a9565b6020820191508190509392505050565b61130081610f6b565b82525050565b600060808201905061131b6000830187610c8d565b61132860208301866112f7565b6113356040830185610c8d565b6113426060830184610c8d565b95945050505050565b7f5369676e61747572657320646f206e6f74206d61746368000000000000000000600082015250565b6000611381601783611058565b915061138c8261134b565b602082019050919050565b600060208201905081810360008301526113b081611374565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006113f182610dd6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611423576114226113b7565b5b600182019050919050565b600081905092915050565b600061144482610cb7565b61144e818561142e565b935061145e818560208601610cd3565b80840191505092915050565b60006114768284611439565b915081905092915050565b7f46756e6374696f6e2063616c6c206e6f74207375636365737366756c00000000600082015250565b60006114b7601c83611058565b91506114c282611481565b602082019050919050565b600060208201905081810360008301526114e6816114aa565b9050919050565b60006060820190506115026000830186610c59565b61150f6020830185610c59565b81810360408301526115218184610d0e565b9050949350505050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261158d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611550565b6115978683611550565b95508019841693508086168417925050509392505050565b6000819050919050565b60006115d46115cf6115ca84610dd6565b6115af565b610dd6565b9050919050565b6000819050919050565b6115ee836115b9565b6116026115fa826115db565b84845461155d565b825550505050565b600090565b61161761160a565b6116228184846115e5565b505050565b5b818110156116465761163b60008261160f565b600181019050611628565b5050565b601f82111561168b5761165c8161152b565b61166584611540565b81016020851015611674578190505b61168861168085611540565b830182611627565b50505b505050565b600082821c905092915050565b60006116ae60001984600802611690565b1980831691505092915050565b60006116c7838361169d565b9150826002028217905092915050565b6116e082610cb7565b67ffffffffffffffff8111156116f9576116f8610e14565b5b611703825461118a565b61170e82828561164a565b600060209050601f831160018114611741576000841561172f578287015190505b61173985826116bb565b8655506117a1565b601f19841661174f8661152b565b60005b8281101561177757848901518255600182019150602085019450602081019050611752565b868310156117945784890151611790601f89168261169d565b8355505b6001600288020188555050505b505050505050565b600081546117b68161118a565b6117c0818661142e565b945060018216600081146117db57600181146117f057611823565b60ff1983168652811515820286019350611823565b6117f98561152b565b60005b8381101561181b578154818901526001820191506020810190506117fc565b838801955050505b50505092915050565b600061183882846117a9565b91508190509291505056fe4d6574615472616e73616374696f6e2875696e74323536206e6f6e63652c616464726573732066726f6d2c62797465732066756e6374696f6e5369676e617475726529454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429a2646970667358221220b060235ba3c2ca2eb78103d962c1d55e1c2059e9c8390a567375fd88b9ce373f64736f6c63430008120033454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c75696e7432353620636861696e49642c6164647265737320766572696679696e67436f6e747261637429",
}

// ForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ForwarderMetaData.ABI instead.
var ForwarderABI = ForwarderMetaData.ABI

// ForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ForwarderMetaData.Bin instead.
var ForwarderBin = ForwarderMetaData.Bin

// DeployForwarder deploys a new Ethereum contract, binding an instance of Forwarder to it.
func DeployForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Forwarder, error) {
	parsed, err := ForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// Forwarder is an auto generated Go binding around an Ethereum contract.
type Forwarder struct {
	ForwarderCaller     // Read-only binding to the contract
	ForwarderTransactor // Write-only binding to the contract
	ForwarderFilterer   // Log filterer for contract events
}

// ForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderSession struct {
	Contract     *Forwarder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderCallerSession struct {
	Contract *ForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderTransactorSession struct {
	Contract     *ForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderRaw struct {
	Contract *Forwarder // Generic contract binding to access the raw methods on
}

// ForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderCallerRaw struct {
	Contract *ForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderTransactorRaw struct {
	Contract *ForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarder creates a new instance of Forwarder, bound to a specific deployed contract.
func NewForwarder(address common.Address, backend bind.ContractBackend) (*Forwarder, error) {
	contract, err := bindForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Forwarder{ForwarderCaller: ForwarderCaller{contract: contract}, ForwarderTransactor: ForwarderTransactor{contract: contract}, ForwarderFilterer: ForwarderFilterer{contract: contract}}, nil
}

// NewForwarderCaller creates a new read-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderCaller(address common.Address, caller bind.ContractCaller) (*ForwarderCaller, error) {
	contract, err := bindForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderCaller{contract: contract}, nil
}

// NewForwarderTransactor creates a new write-only instance of Forwarder, bound to a specific deployed contract.
func NewForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderTransactor, error) {
	contract, err := bindForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderTransactor{contract: contract}, nil
}

// NewForwarderFilterer creates a new log filterer instance of Forwarder, bound to a specific deployed contract.
func NewForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ForwarderFilterer, error) {
	contract, err := bindForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForwarderFilterer{contract: contract}, nil
}

// bindForwarder binds a generic wrapper to an already deployed contract.
func bindForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForwarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.ForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.ForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Forwarder *ForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Forwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Forwarder *ForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Forwarder *ForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Forwarder.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Forwarder *ForwarderCaller) EIP712DOMAIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "EIP712_DOMAIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Forwarder *ForwarderSession) EIP712DOMAIN() (string, error) {
	return _Forwarder.Contract.EIP712DOMAIN(&_Forwarder.CallOpts)
}

// EIP712DOMAIN is a free data retrieval call binding the contract method 0xe1b11da4.
//
// Solidity: function EIP712_DOMAIN() view returns(string)
func (_Forwarder *ForwarderCallerSession) EIP712DOMAIN() (string, error) {
	return _Forwarder.Contract.EIP712DOMAIN(&_Forwarder.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc7977be7.
//
// Solidity: function EIP712_DOMAIN_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "EIP712_DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc7977be7.
//
// Solidity: function EIP712_DOMAIN_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPEHASH(&_Forwarder.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc7977be7.
//
// Solidity: function EIP712_DOMAIN_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Forwarder.Contract.EIP712DOMAINTYPEHASH(&_Forwarder.CallOpts)
}

// METATRANSACTIONTYPE is a free data retrieval call binding the contract method 0xbbb433c4.
//
// Solidity: function META_TRANSACTION_TYPE() view returns(string)
func (_Forwarder *ForwarderCaller) METATRANSACTIONTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "META_TRANSACTION_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// METATRANSACTIONTYPE is a free data retrieval call binding the contract method 0xbbb433c4.
//
// Solidity: function META_TRANSACTION_TYPE() view returns(string)
func (_Forwarder *ForwarderSession) METATRANSACTIONTYPE() (string, error) {
	return _Forwarder.Contract.METATRANSACTIONTYPE(&_Forwarder.CallOpts)
}

// METATRANSACTIONTYPE is a free data retrieval call binding the contract method 0xbbb433c4.
//
// Solidity: function META_TRANSACTION_TYPE() view returns(string)
func (_Forwarder *ForwarderCallerSession) METATRANSACTIONTYPE() (string, error) {
	return _Forwarder.Contract.METATRANSACTIONTYPE(&_Forwarder.CallOpts)
}

// METATRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0xfa2c322d.
//
// Solidity: function META_TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderCaller) METATRANSACTIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "META_TRANSACTION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// METATRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0xfa2c322d.
//
// Solidity: function META_TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderSession) METATRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Forwarder.Contract.METATRANSACTIONTYPEHASH(&_Forwarder.CallOpts)
}

// METATRANSACTIONTYPEHASH is a free data retrieval call binding the contract method 0xfa2c322d.
//
// Solidity: function META_TRANSACTION_TYPEHASH() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) METATRANSACTIONTYPEHASH() ([32]byte, error) {
	return _Forwarder.Contract.METATRANSACTIONTYPEHASH(&_Forwarder.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_Forwarder *ForwarderCaller) ChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_Forwarder *ForwarderSession) ChainID() (*big.Int, error) {
	return _Forwarder.Contract.ChainID(&_Forwarder.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint256)
func (_Forwarder *ForwarderCallerSession) ChainID() (*big.Int, error) {
	return _Forwarder.Contract.ChainID(&_Forwarder.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_Forwarder *ForwarderCaller) ContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "contractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_Forwarder *ForwarderSession) ContractAddress() (common.Address, error) {
	return _Forwarder.Contract.ContractAddress(&_Forwarder.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() view returns(address)
func (_Forwarder *ForwarderCallerSession) ContractAddress() (common.Address, error) {
	return _Forwarder.Contract.ContractAddress(&_Forwarder.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() view returns(bytes)
func (_Forwarder *ForwarderCaller) DataHash(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "dataHash")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() view returns(bytes)
func (_Forwarder *ForwarderSession) DataHash() ([]byte, error) {
	return _Forwarder.Contract.DataHash(&_Forwarder.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0x1b3012a3.
//
// Solidity: function dataHash() view returns(bytes)
func (_Forwarder *ForwarderCallerSession) DataHash() ([]byte, error) {
	return _Forwarder.Contract.DataHash(&_Forwarder.CallOpts)
}

// Digest is a free data retrieval call binding the contract method 0x52a82b65.
//
// Solidity: function digest() view returns(bytes32)
func (_Forwarder *ForwarderCaller) Digest(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "digest")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Digest is a free data retrieval call binding the contract method 0x52a82b65.
//
// Solidity: function digest() view returns(bytes32)
func (_Forwarder *ForwarderSession) Digest() ([32]byte, error) {
	return _Forwarder.Contract.Digest(&_Forwarder.CallOpts)
}

// Digest is a free data retrieval call binding the contract method 0x52a82b65.
//
// Solidity: function digest() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) Digest() ([32]byte, error) {
	return _Forwarder.Contract.Digest(&_Forwarder.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Forwarder *ForwarderCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Forwarder *ForwarderSession) DomainSeparator() ([32]byte, error) {
	return _Forwarder.Contract.DomainSeparator(&_Forwarder.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) DomainSeparator() ([32]byte, error) {
	return _Forwarder.Contract.DomainSeparator(&_Forwarder.CallOpts)
}

// FirstLayerAbi is a free data retrieval call binding the contract method 0xe5ea988b.
//
// Solidity: function firstLayerAbi() view returns(bytes)
func (_Forwarder *ForwarderCaller) FirstLayerAbi(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "firstLayerAbi")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FirstLayerAbi is a free data retrieval call binding the contract method 0xe5ea988b.
//
// Solidity: function firstLayerAbi() view returns(bytes)
func (_Forwarder *ForwarderSession) FirstLayerAbi() ([]byte, error) {
	return _Forwarder.Contract.FirstLayerAbi(&_Forwarder.CallOpts)
}

// FirstLayerAbi is a free data retrieval call binding the contract method 0xe5ea988b.
//
// Solidity: function firstLayerAbi() view returns(bytes)
func (_Forwarder *ForwarderCallerSession) FirstLayerAbi() ([]byte, error) {
	return _Forwarder.Contract.FirstLayerAbi(&_Forwarder.CallOpts)
}

// FirstLayerDigest is a free data retrieval call binding the contract method 0x9c96203b.
//
// Solidity: function first_layer_digest() view returns(bytes32)
func (_Forwarder *ForwarderCaller) FirstLayerDigest(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "first_layer_digest")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FirstLayerDigest is a free data retrieval call binding the contract method 0x9c96203b.
//
// Solidity: function first_layer_digest() view returns(bytes32)
func (_Forwarder *ForwarderSession) FirstLayerDigest() ([32]byte, error) {
	return _Forwarder.Contract.FirstLayerDigest(&_Forwarder.CallOpts)
}

// FirstLayerDigest is a free data retrieval call binding the contract method 0x9c96203b.
//
// Solidity: function first_layer_digest() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) FirstLayerDigest() ([32]byte, error) {
	return _Forwarder.Contract.FirstLayerDigest(&_Forwarder.CallOpts)
}

// FunctionSig is a free data retrieval call binding the contract method 0xaef61c53.
//
// Solidity: function functionSig() view returns(bytes32)
func (_Forwarder *ForwarderCaller) FunctionSig(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "functionSig")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FunctionSig is a free data retrieval call binding the contract method 0xaef61c53.
//
// Solidity: function functionSig() view returns(bytes32)
func (_Forwarder *ForwarderSession) FunctionSig() ([32]byte, error) {
	return _Forwarder.Contract.FunctionSig(&_Forwarder.CallOpts)
}

// FunctionSig is a free data retrieval call binding the contract method 0xaef61c53.
//
// Solidity: function functionSig() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) FunctionSig() ([32]byte, error) {
	return _Forwarder.Contract.FunctionSig(&_Forwarder.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Forwarder *ForwarderCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Forwarder *ForwarderSession) GetChainId() (*big.Int, error) {
	return _Forwarder.Contract.GetChainId(&_Forwarder.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_Forwarder *ForwarderCallerSession) GetChainId() (*big.Int, error) {
	return _Forwarder.Contract.GetChainId(&_Forwarder.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256)
func (_Forwarder *ForwarderCaller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256)
func (_Forwarder *ForwarderSession) GetNonce(user common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256)
func (_Forwarder *ForwarderCallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _Forwarder.Contract.GetNonce(&_Forwarder.CallOpts, user)
}

// NameHash is a free data retrieval call binding the contract method 0xf172a4ce.
//
// Solidity: function nameHash() view returns(bytes32)
func (_Forwarder *ForwarderCaller) NameHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "nameHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NameHash is a free data retrieval call binding the contract method 0xf172a4ce.
//
// Solidity: function nameHash() view returns(bytes32)
func (_Forwarder *ForwarderSession) NameHash() ([32]byte, error) {
	return _Forwarder.Contract.NameHash(&_Forwarder.CallOpts)
}

// NameHash is a free data retrieval call binding the contract method 0xf172a4ce.
//
// Solidity: function nameHash() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) NameHash() ([32]byte, error) {
	return _Forwarder.Contract.NameHash(&_Forwarder.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Forwarder *ForwarderCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Forwarder *ForwarderSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Forwarder.Contract.Nonces(&_Forwarder.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Forwarder *ForwarderCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Forwarder.Contract.Nonces(&_Forwarder.CallOpts, arg0)
}

// RetrievedAddress is a free data retrieval call binding the contract method 0x13e48129.
//
// Solidity: function retrievedAddress() view returns(address)
func (_Forwarder *ForwarderCaller) RetrievedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "retrievedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RetrievedAddress is a free data retrieval call binding the contract method 0x13e48129.
//
// Solidity: function retrievedAddress() view returns(address)
func (_Forwarder *ForwarderSession) RetrievedAddress() (common.Address, error) {
	return _Forwarder.Contract.RetrievedAddress(&_Forwarder.CallOpts)
}

// RetrievedAddress is a free data retrieval call binding the contract method 0x13e48129.
//
// Solidity: function retrievedAddress() view returns(address)
func (_Forwarder *ForwarderCallerSession) RetrievedAddress() (common.Address, error) {
	return _Forwarder.Contract.RetrievedAddress(&_Forwarder.CallOpts)
}

// SecondLayer is a free data retrieval call binding the contract method 0xb76ed6e3.
//
// Solidity: function second_layer() view returns(bytes)
func (_Forwarder *ForwarderCaller) SecondLayer(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "second_layer")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SecondLayer is a free data retrieval call binding the contract method 0xb76ed6e3.
//
// Solidity: function second_layer() view returns(bytes)
func (_Forwarder *ForwarderSession) SecondLayer() ([]byte, error) {
	return _Forwarder.Contract.SecondLayer(&_Forwarder.CallOpts)
}

// SecondLayer is a free data retrieval call binding the contract method 0xb76ed6e3.
//
// Solidity: function second_layer() view returns(bytes)
func (_Forwarder *ForwarderCallerSession) SecondLayer() ([]byte, error) {
	return _Forwarder.Contract.SecondLayer(&_Forwarder.CallOpts)
}

// VersionHash is a free data retrieval call binding the contract method 0x152b5c0f.
//
// Solidity: function versionHash() view returns(bytes32)
func (_Forwarder *ForwarderCaller) VersionHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Forwarder.contract.Call(opts, &out, "versionHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VersionHash is a free data retrieval call binding the contract method 0x152b5c0f.
//
// Solidity: function versionHash() view returns(bytes32)
func (_Forwarder *ForwarderSession) VersionHash() ([32]byte, error) {
	return _Forwarder.Contract.VersionHash(&_Forwarder.CallOpts)
}

// VersionHash is a free data retrieval call binding the contract method 0x152b5c0f.
//
// Solidity: function versionHash() view returns(bytes32)
func (_Forwarder *ForwarderCallerSession) VersionHash() ([32]byte, error) {
	return _Forwarder.Contract.VersionHash(&_Forwarder.CallOpts)
}

// CalculatedAddress is a paid mutator transaction binding the contract method 0xeea8ebe3.
//
// Solidity: function calculatedAddress(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns()
func (_Forwarder *ForwarderTransactor) CalculatedAddress(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "calculatedAddress", userAddress, functionSignature, sigR, sigS, sigV)
}

// CalculatedAddress is a paid mutator transaction binding the contract method 0xeea8ebe3.
//
// Solidity: function calculatedAddress(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns()
func (_Forwarder *ForwarderSession) CalculatedAddress(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _Forwarder.Contract.CalculatedAddress(&_Forwarder.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// CalculatedAddress is a paid mutator transaction binding the contract method 0xeea8ebe3.
//
// Solidity: function calculatedAddress(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns()
func (_Forwarder *ForwarderTransactorSession) CalculatedAddress(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _Forwarder.Contract.CalculatedAddress(&_Forwarder.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0xa9cae653.
//
// Solidity: function executeMetaTransaction(address userAddress, address targetContract, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns()
func (_Forwarder *ForwarderTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, targetContract common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _Forwarder.contract.Transact(opts, "executeMetaTransaction", userAddress, targetContract, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0xa9cae653.
//
// Solidity: function executeMetaTransaction(address userAddress, address targetContract, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns()
func (_Forwarder *ForwarderSession) ExecuteMetaTransaction(userAddress common.Address, targetContract common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteMetaTransaction(&_Forwarder.TransactOpts, userAddress, targetContract, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0xa9cae653.
//
// Solidity: function executeMetaTransaction(address userAddress, address targetContract, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) returns()
func (_Forwarder *ForwarderTransactorSession) ExecuteMetaTransaction(userAddress common.Address, targetContract common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _Forwarder.Contract.ExecuteMetaTransaction(&_Forwarder.TransactOpts, userAddress, targetContract, functionSignature, sigR, sigS, sigV)
}

// ForwarderMetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the Forwarder contract.
type ForwarderMetaTransactionExecutedIterator struct {
	Event *ForwarderMetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *ForwarderMetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForwarderMetaTransactionExecuted)
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
		it.Event = new(ForwarderMetaTransactionExecuted)
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
func (it *ForwarderMetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForwarderMetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForwarderMetaTransactionExecuted represents a MetaTransactionExecuted event raised by the Forwarder contract.
type ForwarderMetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_Forwarder *ForwarderFilterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*ForwarderMetaTransactionExecutedIterator, error) {

	logs, sub, err := _Forwarder.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &ForwarderMetaTransactionExecutedIterator{contract: _Forwarder.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_Forwarder *ForwarderFilterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *ForwarderMetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _Forwarder.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForwarderMetaTransactionExecuted)
				if err := _Forwarder.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_Forwarder *ForwarderFilterer) ParseMetaTransactionExecuted(log types.Log) (*ForwarderMetaTransactionExecuted, error) {
	event := new(ForwarderMetaTransactionExecuted)
	if err := _Forwarder.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
