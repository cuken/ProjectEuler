/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// problemCmd represents the problem command
var problemCmd = &cobra.Command{
	Use:     "problem",
	Aliases: []string{"p"},
	Short:   "Interacts with the problems in Project Euler",
	Long: `Allows you to interact with problems on the Project Euler website
and your local file system. Includes commands for getting new problems
from the website as well as marking local files as solved.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(problemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// problemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// problemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
