// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

// reposCmd represents the repos command
var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Show all repositories",
	Long:  `Show all repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := range args {
			s, u, p := getServerCredential(args[i])
			hub, err := registry.New(fmt.Sprintf("https://%s", s), u, p, Verbose)
			if err != nil {
				log.Fatal(err)
			}
			repos, err := hub.Repositories()
			if err != nil {
				log.Fatal(err)
			}
			for _, repo := range repos {
				fmt.Printf("%s/%s\n", s, repo)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(reposCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reposCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reposCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
