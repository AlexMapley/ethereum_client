package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	fmt.Printf("\n\n%v\n\n", client)

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	fmt.Printf("\n\n%v\n\n", account)

	fmt.Println(account.Hex())        // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(account.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(account.Bytes())      // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(balance) // 25893180161173005034

	// blockNumber := big.NewInt(5532993)
	// balanceAt, err = client.BalanceAt(context.Background(), account, blockNumber)
	balanceAt, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(balance) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042

	fmt.Println(ethValue) // 25.729324269165216041

}
