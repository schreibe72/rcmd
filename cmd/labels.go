// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	"github.com/schreibe72/rcmd/registry"
	"github.com/spf13/cobra"
)

// labelsCmd represents the labels command
var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "Show all Labels for a Repo Tag",
	Long:  `Show all Labels for a Repo Tag`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, tag, err := splitRepoTag(args...)
		if err != nil {
			log.Fatal(err)
		}
		hub, err := registry.New(Server, Username, Password, Verbose)
		if err != nil {
			log.Fatal(err)
		}
		labels, err := hub.Labels(repo, tag)
		for k, v := range labels {
			fmt.Printf("\t%s => %s\n", k, v)
		}
	},
}

func init() {
	RootCmd.AddCommand(labelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// labelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
