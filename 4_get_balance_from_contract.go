package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"regexp"
)

func main() {

	//rinkeby
	addr := "0xbfDC6603DC5938D9d75b580c92b280183D4db020"
	//mainnet
	//addr := "0xC35A32D77Fc4e7832A7A75b4388BCDEa3c237eCd"

	fmt.Printf("Balance of %s\n\n", addr)

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("Address is valid: %v\n", re.MatchString(addr))

	//client, _ := ethclient.Dial("https://mainnet.infura.io/v3/816cbf2267454e52ac7633d5429dc5d5")
	client, _ := ethclient.Dial("https://rinkeby.infura.io/v3/816cbf2267454e52ac7633d5429dc5d5")

	account := common.HexToAddress(addr)
	balance, _ := client.BalanceAt(context.Background(), account, nil)
	fmt.Println(balance)

	ethBal := new(big.Float)
	ethBal.SetString(balance.String())
	ethValue := new(big.Float).Quo(ethBal, big.NewFloat(math.Pow10(18)))

	fmt.Printf("Balance is %d Wei %f Eth\n", balance, ethValue)

}
