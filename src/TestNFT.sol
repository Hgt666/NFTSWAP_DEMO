// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import  {ERC721} from "@openzeppelin/token/ERC721/ERC721.sol";
import {Ownable} from "@openzeppelin/access/Ownable.sol";
import {ERC721URIStorage} from "@openzeppelin/token/ERC721/extensions/ERC721URIStorage.sol";




/**
 * @title TestNFT
 * @dev 测试用 ERC721 NFT 合约，支持单铸造、批量铸造
 */
contract TestNFT is ERC721, ERC721URIStorage, Ownable {
    uint256 public tokenCounter;
    string public baseURI;

    constructor(
        string memory name,
        string memory symbol,
        string memory _baseURI
    ) ERC721(name, symbol) Ownable(msg.sender) {
        baseURI = _baseURI;
        tokenCounter = 0;
    }

    // 更新基础 URI
    function setBaseURI(string calldata newBaseURI) external onlyOwner {
        baseURI = newBaseURI;
    }

    // ========== 工具函数：uint256 转字符串（替代 .toString()） ==========
    function uint2Str(uint256 _num) internal pure returns (string memory) {
        if (_num == 0) return "0";
        uint256 temp = _num;
        uint256 digits;
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
        bytes memory buffer = new bytes(digits);
        while (_num != 0) {
            digits -= 1;
            buffer[digits] = bytes1(uint8(48 + _num % 10));
            _num /= 10;
        }
        return string(buffer);
    }

    // 单铸造：给调用者铸造 1 个 NFT
    function mint() external returns (uint256) {
        tokenCounter++;
        uint256 newTokenId = tokenCounter;
        _safeMint(msg.sender, newTokenId);
        // 调用自定义数值转字符串
        _setTokenURI(newTokenId, string(abi.encodePacked(baseURI, uint2Str(newTokenId))));
        return newTokenId;
    }

    // 批量铸造：一次铸造 count 个 NFT
    function batchMint(uint256 count) external returns (uint256[] memory) {
        require(count > 0 && count <= 100, "Invalid count");
        uint256[] memory tokenIds = new uint256[](count);

        for (uint256 i = 0; i < count; i++) {
            tokenCounter++;
            uint256 newTokenId = tokenCounter;
            _safeMint(msg.sender, newTokenId);
            _setTokenURI(newTokenId, string(abi.encodePacked(baseURI, uint2Str(newTokenId))));
            tokenIds[i] = newTokenId;
        }
        return tokenIds;
    }

    // 重写 tokenURI
    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    // 修复：同时重写两个父合约的 _burn
    // function _burn(uint256 tokenId) internal override(ERC721) {
    //     super._burn(tokenId);
    // }

    // 重写接口支持
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}