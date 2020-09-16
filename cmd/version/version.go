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

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	BuildDate string
	GitBranch string
	GitHash   string
	Version   string
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display detailed version information",
		Long:  `Displays information about the version of this software.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("BuildDate: %s\nGitBranch: %s\nGitHash: %s\nVersion: %s\n", BuildDate, GitBranch, GitHash, Version)
		},
	}
}
