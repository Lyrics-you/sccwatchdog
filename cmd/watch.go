package cmd

import (
	"errors"
	"sccwatchdog/dog"
	"sccwatchdog/utils"
	"strings"

	"github.com/spf13/cobra"
)

var (
	second int
	except string
)
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch changed and restarted by namespace and deployments ",
	Long: `
Watch all deployments in namespace by t second
eg.: swd watch [-n <namespace>(default:"default")] [-s t]
Watch specified deployments "deploy1 deploy2" in namespace by t second
eg.: swd watch [-n <namespace>(default:"default")] [-d "deploy1 deploy2"] [-s t]
Not Watch expect deployments "deploy2" in namespace by t second
eg.: swd watch [-n <namespace>(default:"default")] [-d "deploy1 deploy2"] [-e "deploy2"] [-s t]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			Error(cmd, args, errors.New("unrecognized command"))
			return
		}
		if deployment == "" {
			watchAllDeplolyments(namespace, except, second)
		} else {
			watchDeplolyments(namespace, deployment, except, second)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	watchCmd.PersistentFlags().StringVarP(&deployment, "depolyment", "d", "", "scc depolyment")
	watchCmd.PersistentFlags().StringVarP(&except, "except", "e", "", "expect depolyment")
	watchCmd.PersistentFlags().IntVarP(&second, "second", "s", 0, "times interval (second)")
}

func watchDeplolyments(namespace, depolyment, except string, s int) {
	realDeploys := []string{}
	if except != "" {
		for _, d := range strings.Split(depolyment, " ") {
			if !strings.Contains(except, d) {
				realDeploys = append(realDeploys, d)
			}
		}
	} else {
		realDeploys = strings.Split(depolyment, " ")
	}
	deploys, err := utils.GetDeploymentInfos(namespace, realDeploys)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	dog.WatchStart(deploys, s)
}

func watchAllDeplolyments(namespaces, except string, s int) {
	deploys, err := utils.GetAllDeploymentsInfos(namespace)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	dog.WatchAllStart(deploys, namespaces, except, s)
}
