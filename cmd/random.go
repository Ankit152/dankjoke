/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dank joke 😆",
	Long:  `This commad fetches you with a random joke on its own!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

// dank joke structure
type dankJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}
