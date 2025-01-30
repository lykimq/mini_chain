package main

import (

	"bufio"
	"fmt"
	"os"
	"strings"

	"mini_chain/internal/blockchain"
)

func main() {

	// Initialize the blockchain
	chain := blockchain.NewBlockchain()

	// Create scanner for reading user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a comand (add/list/validate/quit):")
		scanner.Scan()

		command := strings.ToLower(scanner.Text())
		switch command {
		case "add":
			fmt.Print("Enter data for the new block: ")
			scanner.Scan()
			data := scanner.Text()

			if err:= chain.AddBlock(data); err != nil {
				fmt.Println("Error adding block:", err)
			} else {
				fmt.Println("Block added successfully")
			}

		case "list":
			chain.PrintChain()

		case "validate":
			err := chain.IsValid()
			if err != nil {
				fmt.Println("Blockchain is invalid:", err)
			} else {
				fmt.Println("Blockchain is valid")
			}

		case "quit":
			fmt.Println("Exiting the program")
			return

		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}