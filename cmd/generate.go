package cmd

import (
	"fmt"

	"math/rand"

	"github.com/spf13/cobra"
)

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long: `Generate random passwords with customizable options, like:
	
	password-gen-cli generate -l 12 -d -s`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCommand)

	generateCommand.Flags().IntP("length", "l", 16, "Length of generated password")
	generateCommand.Flags().BoolP("digits", "d", false, "Include digits in password")
	generateCommand.Flags().BoolP("special-chars", "s", false, "Include special characters in password")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charset += "1234567890"
	}

	if isSpecialChars {
		charset += "!@#$%^&*()_+{}:|<>?,.;"
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println(string(password))
}
