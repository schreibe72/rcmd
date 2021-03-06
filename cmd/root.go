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
	"os"

	"github.com/schreibe72/rcmd/azure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ServerCredentials struct {
	Username string
	Password string
}

var (
	cfgFile           string
	Servers           map[string]ServerCredentials
	Username          string
	Password          string
	Verbose           bool
	Azure             bool
	Version           string
	Githash           string
	SubscriptionNames []string
	registriesNames   []string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rcmd",
	Short: "A small command Tool for managing a docker registry",
	Long:  `A small command Tool for managing a docker registry`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rcmd.yaml)")
	RootCmd.PersistentFlags().StringVarP(&Username, "username", "U", "", "Username")
	RootCmd.PersistentFlags().StringVarP(&Password, "password", "O", "", "Password")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&Azure, "azure", "a", false, "get registry config from azure")
	RootCmd.PersistentFlags().StringArrayVarP(&SubscriptionNames, "subscription", "S", []string{}, "Use only this Subscriptions. No Value means: take them all")
	RootCmd.PersistentFlags().StringArrayVarP(&registriesNames, "registries", "R", []string{}, "Use only this registries. No Value means: take them all")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if Azure {
		initAzureConfig()
	} else {
		initConfigFile()
	}
}
func initAzureConfig() {
	Servers = map[string]ServerCredentials{}
	s, err := azure.GetSubscriptions(SubscriptionNames...)
	if err != nil {
		panic(err)
	}
	for _, id := range s.GetIDs() {
		registries, err := azure.GetContainerRegistries(id)
		if err != nil {
			panic(err)
		}
		for _, r := range registries {
			if len(registriesNames) == 0 || contains(registriesNames, r.LoginServer) || contains(registriesNames, r.Name) {
				Servers[r.LoginServer] = ServerCredentials{Username: r.Login, Password: r.Password}
			}
		}
	}
}

func initConfigFile() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".rcmd") // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if Verbose {
			log.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
	Servers = map[string]ServerCredentials{}
	viper.UnmarshalKey("Servers", &Servers)

	for r := range Servers {
		if len(registriesNames) > 0 && !contains(registriesNames, r) {
			delete(Servers, r)
		}
	}
}
