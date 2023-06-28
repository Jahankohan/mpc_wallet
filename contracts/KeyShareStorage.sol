// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

contract KeyShareStorage {
    address private owner;
    bool private paused;

    mapping(address => bool) private acl;
    mapping(bytes32 => bytes) private shares;

    event ShareStored(bytes32 indexed shareId, address indexed by);
    event ShareUpdated(bytes32 indexed shareId, address indexed by);
    event ShareDeleted(bytes32 indexed shareId, address indexed by);
    event Paused(address indexed by);
    event Unpaused(address indexed by);

    constructor() {
        owner = msg.sender;
        acl[owner] = true;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not authorized");
        _;
    }

    modifier onlyAuthorized() {
        require(acl[msg.sender], "Not authorized");
        _;
    }

    modifier whenNotPaused() {
        require(!paused, "Contract is paused");
        _;
    }

    modifier whenPaused() {
        require(paused, "Contract is not paused");
        _;
    }

    function addAuthorizedAddress(address addr) public onlyOwner {
        acl[addr] = true;
    }

    function removeAuthorizedAddress(address addr) public onlyOwner {
        acl[addr] = false;
    }

    function storeShare(bytes32 shareId, bytes memory share) public onlyAuthorized whenNotPaused {
        require(shares[shareId].length == 0, "Share ID already exists");
        shares[shareId] = share;
        emit ShareStored(shareId, msg.sender);
    }

    function getShare(bytes32 shareId) public view onlyAuthorized whenNotPaused returns (bytes memory) {
        return shares[shareId];
    }

    function updateShare(bytes32 shareId, bytes memory newShare) public onlyAuthorized whenNotPaused {
        shares[shareId] = newShare;
        emit ShareUpdated(shareId, msg.sender);
    }

    function deleteShare(bytes32 shareId) public onlyAuthorized whenNotPaused {
        delete shares[shareId];
        emit ShareDeleted(shareId, msg.sender);
    }

    function pause() public onlyOwner whenNotPaused {
        paused = true;
        emit Paused(msg.sender);
    }

    function unpause() public onlyOwner whenPaused {
        paused = false;
        emit Unpaused(msg.sender);
    }
}
