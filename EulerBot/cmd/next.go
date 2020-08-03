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
	"io/ioutil"
	"log"
	"time"

	helper "github.com/cuken/ProjectEuler/EulerBot/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Complets the current problem and fetches the next one",
	Long:  `Sets the completion time for the current problem, increments the current problem and retrieves the next problem.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := viper.GetInt("problem.current")
		p++
		d := fmt.Sprintf("%s/%04d", viper.GetString("local.problemDirectory"), p)
		log.Printf("Great job! Moving onto problem %v", p)
		t, txt, err := helper.GenerateProblem(p, d)
		if err != nil {
			log.Fatalf("Error when attemping to retrieve the next problem: %s", err)
		}
		tNow := time.Now()
		s := tNow.Format(time.RFC1123)
		viper.Set("problem.startTime", s)
		viper.Set("problem.current", p)

		probTitle := []byte(fmt.Sprintf("Currently working on Problem #%v - %s", p, t))
		probLink := []byte(fmt.Sprintf("%s - https://projecteuler.net/problem=%v", txt, p))
		startTime := []byte(fmt.Sprintf("Problem Started: %s", s))

		streamDir := viper.GetString("local.streamDirectory")

		_ = ioutil.WriteFile(fmt.Sprintf("%s/problemTitle.txt", streamDir), probTitle, 0777)
		_ = ioutil.WriteFile(fmt.Sprintf("%s/problemRunner.txt", streamDir), probLink, 0777)
		_ = ioutil.WriteFile(fmt.Sprintf("%s/problemStartTime.txt", streamDir), startTime, 0777)

		err = viper.WriteConfig()
		if err != nil {
			log.Printf("Error when attempting to write start time to config: %s", err)
		}

	},
}

func init() {
	problemCmd.AddCommand(nextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
