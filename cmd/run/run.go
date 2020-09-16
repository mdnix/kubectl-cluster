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

package run

import (
	"context"
	"fmt"
	"kubectl-cluster/pkg/config"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

const (
	KUBECTL = "kubectl"
)

var (
	targets []string
	tags    []string
	wg      *sync.WaitGroup
)

func New() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run a kubectl command",
		Long:  `Run a kubectl command on the specified cluster(s).`,
		RunE: func(cmd *cobra.Command, args []string) error {
			wg = new(sync.WaitGroup)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if len(args) == 0 {
				cmd.Help()
				fmt.Println("\nNo arguments passed")
				os.Exit(1)
			}

			if len(targets) > 0 && len(tags) > 0 {
				cmd.Help()
				fmt.Println("\n--targets and --tags flags can't be used together")
				os.Exit(1)
			}

			if len(targets) == 0 && len(tags) == 0 {
				runAll(ctx, args)
			}
			if len(targets) > 0 {
				runTargets(ctx, args)
			}
			if len(tags) > 0 {
				runTags(ctx, args)
			}
			return nil
		},
	}
	runCmd.Flags().StringSliceVar(&targets, "targets", []string{}, "Set the target cluster for execution of a command")
	runCmd.Flags().StringSliceVar(&tags, "tags", []string{}, "Run command on clusters containing the specified tags")
	return runCmd
}

func setKubeConfigFlag(args []string, i int) []string {
	args = append(args, "--kubeconfig", config.Conf.Clusters[i].Config)
	return args
}

func runAll(ctx context.Context, args []string) {
	for i := range config.Conf.Clusters {
		wg.Add(1)
		go runKubectl(ctx, config.Conf.Clusters[i].Name, args, i, wg)
	}
	wg.Wait()
}

func runTargets(ctx context.Context, args []string) {
	for _, target := range targets {
		for i := range config.Conf.Clusters {
			if config.Conf.Clusters[i].Name == target {
				wg.Add(1)
				go runKubectl(ctx, target, args, i, wg)
			}
		}
	}
	wg.Wait()
}

func runTags(ctx context.Context, args []string) {
	for _, tag := range tags {
		for i := range config.Conf.Clusters {
			if config.Conf.Clusters[i].Tags == tag {
				wg.Add(1)
				go runKubectl(ctx, config.Conf.Clusters[i].Name, args, i, wg)
			}
		}
	}
	wg.Wait()
}

func runKubectl(ctx context.Context, cluster string, args []string, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	args = setKubeConfigFlag(args, i)
	res, err := exec.CommandContext(ctx, KUBECTL, args...).CombinedOutput()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Results from: %s\n\n", cluster)
	fmt.Println(string(res))
}
