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

package list

import (
	"fmt"
	"kubectl-cluster/pkg/config"

	"github.com/spf13/cobra"
)

const (
	KUBECTL = "kubectl"
)

func New() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List available clusters",
		Long:  `List all availble clusters which are defined in the config`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Available clusters: ")
			for i := range config.Conf.Clusters {
				fmt.Printf("\t %s\n", config.Conf.Clusters[i].Name)
			}
		},
	}
	return listCmd
}
