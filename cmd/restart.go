package cmd

import (
	"errors"
	"sccwatchdog/dog"

	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart by namespace and deployment",
	Long: `
eg.: swd restart -n <namespace> -d <deployment> [-c <container>(default:first container,else <deployment>)]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Error(cmd, args, errors.New("unrecognized command"))
			return
		}
		restartDeployment(namespace, deployment, container)
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
	restartCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	restartCmd.PersistentFlags().StringVarP(&deployment, "depolyment", "d", "", "scc depolyment")
	restartCmd.PersistentFlags().StringVarP(&container, "container", "c", "", "depolyment container")
}

func restartDeployment(namespace, deployment, container string) {
	info, err := dog.RestartDeployment(namespace, deployment, container)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	log.Infof("%s", info)
}
