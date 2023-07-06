// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract Forwarder {
    mapping(address => uint256) public nonces;

    // EIP712
    string private constant EIP712_DOMAIN = "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)";
    string private constant META_TRANSACTION_TYPE = "MetaTransaction(uint256 nonce,address from,bytes functionSignature)";
    bytes32 private constant EIP712_DOMAIN_TYPEHASH = keccak256(abi.encodePacked(EIP712_DOMAIN));
    bytes32 private constant META_TRANSACTION_TYPEHASH = keccak256(abi.encodePacked(META_TRANSACTION_TYPE));

    bytes32 private domainSeparator;

    event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature);

    constructor() {
        domainSeparator = keccak256(abi.encode(
            EIP712_DOMAIN_TYPEHASH,
            keccak256("MyApp"), // name
            keccak256("1"),     // version
            getChainId(),
            address(this)
        ));
    }

    function executeMetaTransaction(
        address userAddress,
        bytes memory functionSignature,
        bytes32 sigR,
        bytes32 sigS,
        uint8 sigV
    ) public {
        bytes32 digest = keccak256(abi.encodePacked(
            "\x19\x01",
            domainSeparator,
            keccak256(abi.encode(
                META_TRANSACTION_TYPEHASH,
                nonces[userAddress],
                userAddress,
                keccak256(functionSignature)
            ))
        ));

        require(userAddress == ecrecover(digest, sigV, sigR, sigS), "Signatures do not match");

        nonces[userAddress]++;

        // Delegate call to the target contract
        (bool success,) = address(this).delegatecall(functionSignature);
        require(success, "Function call not successful");

        emit MetaTransactionExecuted(userAddress, msg.sender, functionSignature);
    }

    function getChainId() internal view returns (uint256) {
        uint256 chainId;
        assembly { chainId := chainid() }
        return chainId;
    }

    function getNonce(address user) external view returns(uint256) {
        return nonces[user];
    }
}
