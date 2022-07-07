package cmd

import (
	"errors"
	"fmt"
	"sccwatchdog/utils"
	"strings"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information by namespace and deployments",
	Long: `
Show all deployments in namespace
eg.: swd show [-n <namespace>(default:"default")]
Show specified deployments "deploy1 deploy2" in namespace
eg.: swd show [-n <namespace>(default:"default")] -d "deploy1 deploy2"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Error(cmd, args, errors.New("unrecognized command"))
			return
		}
		if deployment == "" {
			showAllDeployments(namespace)
		} else {
			showDeployments(namespace, deployment)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	showCmd.PersistentFlags().StringVarP(&deployment, "depolyment", "d", "", "scc depolyment")
}

func showDeployments(namespace, depolyment string) {
	deploys := strings.Split(depolyment, " ")
	deploy, err := utils.GetDeploymentInfos(namespace, deploys)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	for _, d := range deploy {
		fmt.Printf("[%s/%s: (%s) @ %s]\n", d.Namespace, d.Name, d.Image, d.LastUpdateTime)
	}
}

func showAllDeployments(namespace string) {
	deploy, err := utils.GetAllDeploymentsInfos(namespace)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	for _, d := range deploy {
		fmt.Printf("[%s/%s: (%s) @ %s]\n", d.Namespace, d.Name, d.Image, d.LastUpdateTime)
	}
}
