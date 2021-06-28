package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/console/prompt"
	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/wallet"
	"github.com/spf13/cobra"
)

const flagAddress = "address"

func walletCmd() *cobra.Command {
	var walletCmds = &cobra.Command{
		Use:   "wallet",
		Short: "Manage your wallet, accounts and keys",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	walletCmds.AddCommand(walletNewAccountCmd())
	walletCmds.AddCommand(walletPrintPrivKeyCmd())

	return walletCmds
}

func walletNewAccountCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new-account",
		Short: "Creates a new account with a new set of a elliptic-curve Private + Public keys.",
		Run: func(cmd *cobra.Command, args []string) {
			password := getPassPhrase("Please enter a password to encrypt the new wallet:", true)

			acc, err := wallet.NewKeystoreAccount(password)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("New account created: %s\n", acc)
			fmt.Printf("Saved in: %s\n", wallet.GetWalletDir())
		},
	}

	return cmd
}

func walletPrintPrivKeyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "print-pk",
		Short: "Unlocks keystore file and prints the Private + Public keys.",
		Run: func(cmd *cobra.Command, args []string) {
			address, _ := cmd.Flags().GetString(flagAddress)
			password := getPassPhrase("Please enter a password to decrypt the wallet:", false)

			key, err := wallet.GetKeyForAddress(common.NewAddress(address), password)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println(key.PrivateKey.D)
		},
	}

	cmd.Flags().String(flagAddress, "", "Address for which the private key needs to be retrieved")
	cmd.MarkFlagRequired(flagAddress)

	return cmd
}

func getPassPhrase(text string, confirmation bool) string {
	if text != "" {
		fmt.Println(text)
	}
	password, err := prompt.Stdin.PromptPassword("Password: ")
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	if confirmation {
		confirm, err := prompt.Stdin.PromptPassword("Repeat password: ")
		if err != nil {
			log.Fatalf("Failed to read password confirmation: %v", err)
		}
		if password != confirm {
			log.Fatalf("Passwords do not match")
		}
	}
	return password
}
