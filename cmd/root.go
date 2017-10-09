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
	"os"

	"github.com/gaocegege/ali-clusterdata-analyzer/analyzer"
	"github.com/spf13/cobra"

	"github.com/gaocegege/ali-clusterdata-analyzer/parser"
)

var (
	clusterDataDir string
	GlobalAnalyzer *analyzer.Analyzer
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ali-clusterdata-analyzer",
	Short: "Analyzer for alibaba cluster data",
	Long:  `For your research paper`,
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
	cobra.OnInitialize(initParser)
	cobra.OnInitialize(initAnalyzer)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ali-clusterdata-analyzer.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.PersistentFlags().StringVar(&clusterDataDir, "data-dir", "", "Alibaba cluster data directory")
}

func initParser() {
	if clusterDataDir == "" {
		log.Fatalln("The directory must be specified")
	}
	batchInstanceParser = parser.NewBatchInstanceParser(clusterDataDir)
}

func initAnalyzer() {
	GlobalAnalyzer = analyzer.NewAnalyzer()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" { // enable ability to specify config file via flag
	// 	viper.SetConfigFile(cfgFile)
	// }

	// viper.SetConfigName(".ali-clusterdata-analyzer") // name of config file (without extension)
	// viper.AddConfigPath("$HOME")  // adding home directory as first search path
	// viper.AutomaticEnv()          // read in environment variables that match

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}
