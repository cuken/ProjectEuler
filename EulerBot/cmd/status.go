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
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Returns the current problem and the start time / duration",
	Long: `Reads the config file to retrieve the current problem, returns the time
the problem was started, and the duration since the start.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := viper.GetString("problem.current")
		s := viper.GetString("problem.startTime")
		start, _ := time.Parse(time.RFC1123, s)
		diff := time.Since(start)
		fmt.Printf("You are currently working on Problem: %v\nYou started it on: %v\nTotal Duration: %v\n", p, start, diff)
	},
}

func init() {
	problemCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
