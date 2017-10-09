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

	"github.com/gaocegege/ali-clusterdata-analyzer/parser"
	"github.com/spf13/cobra"
)

var batchInstanceParser *parser.BatchInstanceParser

// batch-instanceCmd represents the batch-instance command
var batchInstanceCmd = &cobra.Command{
	Use:   "batch-instance",
	Short: "Batch instance related commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("!!!")
		batchInstances, err := batchInstanceParser.ParseFromFile()
		fmt.Println("!!!")
		if err != nil {
			log.Fatalln("Error occurred in batch-instance: %v", err)
		}
		GlobalAnalyzer.ImportBatchInstance(batchInstances)
		fmt.Println(GlobalAnalyzer)
	},
}

func init() {
	RootCmd.AddCommand(batchInstanceCmd)
	batchInstanceParser = parser.NewBatchInstanceParser(clusterDataDir)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// batch-instanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// batch-instanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
