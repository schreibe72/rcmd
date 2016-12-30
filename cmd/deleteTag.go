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
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// deleteTagCmd represents the deleteTag command
var deleteTagCmd = &cobra.Command{
	Use:   "deleteTag",
	Short: "Delete a repositry tag",
	Long:  `Delete a repositry tag`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Missing Repo:Tag")
		}
		splitparts := strings.Split(args[0], ":")
		if len(splitparts) < 2 {
			log.Fatal("Wrong Format: Repo:Tag")
		}
		tag := splitparts[len(splitparts)-1]
		reposlice := splitparts[0 : len(splitparts)-1]
		repo := strings.Join(reposlice, ":")
		connect()
		digest, err := hub.ManifestDigest(repo, tag)
		if err != nil {
			log.Fatal(err)
		}
		err = hub.DeleteManifest(repo, digest)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteTagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteTagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteTagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
