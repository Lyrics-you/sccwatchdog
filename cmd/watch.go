package cmd

import (
	"errors"
	"sccwatchdog/dog"
	"sccwatchdog/utils"

	"github.com/spf13/cobra"
)

var (
	second int
)
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch changed and restarted by namespace and deployments ",
	Long: `
Watch all deployments in namespace by t second
eg.: swd watch [-n <namespace>(default:"default")] [-s t]
Watch specified deployments "deploy1 deploy2" in namespace by t second
eg.: swd watch [-n <namespace>(default:"default")] -d "deploy1 deploy2" [-s t]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			Error(cmd, args, errors.New("unrecognized command"))
			return
		}
		if deployment == "" {
			watchAllDeplolyments(namespace, second)
		} else {
			watchDeplolyments(namespace, deployment, second)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	watchCmd.PersistentFlags().StringVarP(&deployment, "depolyment", "d", "", "scc depolyment")
	watchCmd.PersistentFlags().IntVarP(&second, "second", "s", 0, "times interval (second)")
}

func watchDeplolyments(namespace, depolyment string, s int) {
	deploys, err := utils.GetDeploymentInfos(namespace, depolyment)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	dog.WatchStart(deploys, s)
}

func watchAllDeplolyments(namespaces string, s int) {
	deploys, err := utils.GetAllDeploymentsInfos(namespace)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	dog.WatchAllStart(deploys, namespaces, s)
}
