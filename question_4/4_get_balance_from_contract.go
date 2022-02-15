package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	addrList := []string{
		"0x123d475e13aa54a43a7421d94caa4459da021c77",
		"0x0020c5222a24e4a96b720c06b803fb8d34adc0af",
		"0xfe808b079187cc460f47374580f5fb47c82b87a5",
	}

	client, _ := ethclient.Dial("https://bsc-dataseed.binance.org/")
	ca := common.HexToAddress("0x250b211EE44459dAd5Cd3bCa803dD6a7EcB5d46C")
	instance, err := NewToken(ca, client)
	if err!= nil {
		fmt.Printf("get the instance of contract err:%v", err)
		return
	}

	for i:=0; i<len(addrList); i++ {
		addr := addrList[i]
		account := common.HexToAddress(addr)
		bl, err2 := instance.BalanceOf(nil, account)
		if err2 != nil {
			fmt.Printf("get Balance from %s err, err=%v\n", addr, err2)
		}
		fmt.Printf("%s %d\n", addr, bl)
	}









	//balance, _ := client.BalanceAt(context.Background(), account, nil)
	//fmt.Println(balance)
	//
	//tokenBal := new(big.Float)
	//tokenBal.SetString(balance.String())
	//tokenValue := new(big.Float).Quo(tokenBal, big.NewFloat(math.Pow10(18)))
	//
	//fmt.Printf("Balance is %f \n", tokenValue)

}
