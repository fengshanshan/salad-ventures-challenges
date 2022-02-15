//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';

contract Utility {
    struct TokenBalances {
        address token;
        uint256 balance;
    }

    function getBalances(address addr, address[] calldata tokens) public view returns(TokenBalances[] memory) {
        TokenBalances[] memory balances = new TokenBalances[](tokens.length);
        for (uint256 i = 0; i < tokens.length; i++) {
            balances[i] = TokenBalances(tokens[i], IERC20(tokens[i]).balanceOf(addr));
        }

        return balances;
    }
}
