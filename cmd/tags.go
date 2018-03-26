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
	"strings"

	"github.com/schreibe72/rcmd/registry"
	"github.com/spf13/cobra"
)

var (
	stringFlag  bool
	descFlag    bool
	intFlag     bool
	versionFlag bool
	sortLabels  string
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Show all tags of a repository",
	Long:  `Show all tags of a repository`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		c, err := getSortCriteria()
		if err != nil {
			return fmt.Errorf("Cannot Use More Than One Sortflags. Do Not Use -i -V -s Together")
		}
		if descFlag && c == "" {
			return fmt.Errorf("Cannot Output in Decent Order. You Need A Sort Flag")
		}
		if sortLabels == "" && c != "" {
			return fmt.Errorf("Cannot Sort Because Missing Sort Label. Use -l")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Missing Repo")
		}
		repo := args[0]
		hub, err := registry.New(Server, Username, Password, Verbose)
		if err != nil {
			log.Fatal(err)
		}
		var tags []string
		if stringFlag || intFlag || versionFlag {
			criteria, _ := getSortCriteria()
			tags, err = hub.SortedTagsByLabel(repo, strings.Split(sortLabels, ","), descFlag, criteria)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			tags, err = hub.Tags(repo)
			if err != nil {
				log.Fatal(err)
			}
		}

		for _, tag := range tags {
			fmt.Printf("%s:%s\n", repo, tag)
		}
	},
}

func getSortCriteria() (string, error) {
	i := 0
	result := ""
	if intFlag {
		i++
		result = "int"
	}
	if versionFlag {
		i++
		result = "version"
	}
	if stringFlag {
		i++
		result = "string"
	}
	if i > 1 {
		return "", fmt.Errorf("Too Many Flags")
	}
	return result, nil
}

func init() {
	RootCmd.AddCommand(tagsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tagsCmd.Flags().BoolVarP(&stringFlag, "sort", "s", false, "Sort label by string")
	tagsCmd.Flags().BoolVarP(&intFlag, "int", "i", false, "Sort label by integer")
	tagsCmd.Flags().BoolVarP(&versionFlag, "version", "V", false, "Sort label by version (X.Y.Z)")
	tagsCmd.Flags().BoolVarP(&descFlag, "desc", "d", false, "decent sort order")
	tagsCmd.Flags().StringVarP(&sortLabels, "labels", "l", "", "SortLabel, use Komma to separate more labels")

}
