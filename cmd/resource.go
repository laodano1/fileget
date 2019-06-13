// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"

	"strings"
)

var (
	typeVal string

)

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "monitored resource type node/pod",
	Long: `monitor all resource type node/pod. For example:
  isdkMonitor monitor resource

  isdkMonitor monitor resource -type node

  isdkMonitor monitor resource -type pod

isdkMonitor is a CLI tool to monitor kubernetes node/pod resource.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if typeVal == "" {
			fmt.Println("monitor all resources.")
			//fmt.Printf("%s\n%s\n", res.Nodes, res.Pods)
			res.MonitorKubeOutput("all")
		} else {
			res.Log.Printf("monitored resource type: %s\n", typeVal)
			switch strings.ToUpper(typeVal) {
			case "POD" :
				//for _, v := range res.Pods {
				//	fmt.Printf("\tpod: '%s' monitored.\n", v)
				//}
				res.MonitorKubeOutput("pod")
			case "NODE" :
				//for _, v := range res.Nodes {
				//	fmt.Printf("\tnode: '%s' monitored.\n", v)
				//}
				res.MonitorKubeOutput("node")
			default:
				res.Log.Println("bad type value.")
			}
		}
		res.TearDown()
	},
}

func init() {
	monitorCmd.AddCommand(resourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourceCmd.PersistentFlags().String("foo", "", "A help for foo")
	//resourceCmd.PersistentFlags().BoolVarP(&isAllResource, "all", "a", false, "monitor all resources, including node/pod.")
	resourceCmd.PersistentFlags().StringVarP(&typeVal, "type", "t", "", "monitor resource type, including node/pod.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
