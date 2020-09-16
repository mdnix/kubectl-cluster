/*
Copyright © 2020 Marco De Luca

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
	"kubectl-cluster/cmd/list"
	"kubectl-cluster/cmd/run"
	"kubectl-cluster/cmd/version"
	"kubectl-cluster/pkg/config"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile         string
	rootCmdExamples = `
	# List available clusters
	kubectl cluster list

	# Get Pods of all clusters using the current context
	kubectl cluster run get pods 

	# Get Pods of all clusters overriding the current namespace of the context
	# NOTE: when using flags for regular kubectl commands "--" has to be added to signify the end of command options
	kubectl cluster run -- get pods --all-namespaces

	# Get Pods of specified clusters using the current context
	kubectl cluster run --targets okd,minikube get pods 

	# Get Pods of clusters containing the specified tag in the config
	kubectl cluster run --tags test get pods 
`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-cluster",
	Short: "kubectl cluster plugin",
	Long: `The kubectl-cluster plugin lets you run kubectl commands across a number of specified clusters. 
It is possible to run any kubectl command since this plugin is basically a wrapper around kubectl.`,
	Version: version.Version,
	Example: rootCmdExamples,
}

func addSubCommand() {
	rootCmd.AddCommand(
		version.New(),
		run.New(),
		list.New())
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "", "config file (default is $HOME/.clusters)")
	addSubCommand()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from flag.
		viper.SetConfigType("yaml")
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfgPath := fmt.Sprintf("%s/%s", home, ".clusters")

		// Search config in home directory with name ".clusters".
		viper.SetConfigType("yaml")
		viper.SetConfigFile(cfgPath)
	}
	// Read the config file
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
		config.Conf = &config.Config{}
		err = viper.Unmarshal(config.Conf)
	} else {
		fmt.Println("Unable to read config file: ", err)
		os.Exit(1)
	}
}
