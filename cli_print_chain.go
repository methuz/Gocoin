package main

import (
	"encoding/hex"
	"fmt"
	"time"
)

func (cli *CLI) printChain() {
	bc := NewBlockchain("")
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()
		blockTime := time.Unix(block.Timestamp, 0)

		fmt.Printf("=============== Block %x ================\n", block.Hash)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Time: %s\n", blockTime.String())
		fmt.Printf("Transactions:\n")

		for txi, tx := range block.Transactions {
			txid := hex.EncodeToString(tx.ID)
			fmt.Printf("%d %s\n len %d:\n", txi, txid, len(tx.ID))
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
