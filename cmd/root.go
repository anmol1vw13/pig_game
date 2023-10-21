/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/anmol1vw13/pig_game/game"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pig",
	Run: func(cmd *cobra.Command, args []string) { 
		game.Run(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
	rootCmd.SetUsageFunc(nil)
	rootCmd.SetUsageTemplate(
		`pig - A command line tool to simulate a game of pig. It is a two-player game played with a 6-sided die.

Usage:
	pig [strategy] [strategy]

Args:
	strategy   The number between 1 to 100

Description:
	This command line application accepts two numbers between 1 to 100 as a positional argument. Strategies for player 1 and player 2, and performs the game of pig simulation on it. If no strategies are provided, then it will return an error. 

	If the number is out of range, the application will exit with an error message. Otherwise, it will perform the simulation and output the result.

Example usage:
	$ pig 10 15
	Result: Holding at   10 vs Holding at   15: wins: 3/10 (30.0%), losses: 7/10 (70.0%)
`,
	)
}


