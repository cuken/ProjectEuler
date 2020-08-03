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
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	helper "github.com/cuken/ProjectEuler/EulerBot/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var force *bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a specific problem",
	Long: `Retrieves a specific file from the Project Euler website.
Will download the file and create a problem folder for you in the specified
problems directory.

Will not do anything if the folder already exists unless specified with the -f flag
to overwrite the contents.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least 1 problem number")
		}
		for _, num := range args {
			_, err := strconv.Atoi(num)
			if err != nil {
				return fmt.Errorf("The provided arg: %s is not a valid number", num)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			num, _ := strconv.Atoi(s)
			d := fmt.Sprintf("%s/%04d", viper.GetString("local.problemDirectory"), num)
			log.Printf("%s", d)
			if _, err := os.Stat(d); !os.IsNotExist(err) {
				// Folder already exists.
				if !*force {
					log.Printf("Problem %v already exists at %s\nYou can supply -f to overwrite existing files.", num, d)
					os.Exit(1)
				}
				os.RemoveAll(d)
			}
			_, _, err := helper.GenerateProblem(num, d)
			if err != nil {
				panic(err)
			}

		}
	},
}

func init() {
	problemCmd.AddCommand(getCmd)
	force = getCmd.Flags().BoolP("force", "f", false, "Overwrites existing files")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
