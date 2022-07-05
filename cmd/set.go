package cmd

import (
	"errors"
	"sccwatchdog/dog"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set image by namespace and deployment",
	Long: `
eg.: swd set -n <namespace> -d <deployment> [-c <container>(default:first container,else <deployment>)] -i <image>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Error(cmd, args, errors.New("unrecognized command"))
			return
		}
		setDeploymentImage(namespace, deployment, container, image)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	setCmd.PersistentFlags().StringVarP(&deployment, "depolyment", "d", "", "scc depolyment")
	setCmd.PersistentFlags().StringVarP(&container, "container", "c", "", "depolyment container")
	setCmd.PersistentFlags().StringVarP(&image, "image", "i", "", "container image")
}

func setDeploymentImage(namespace, deployment, container, image string) {
	info, err := dog.SetDeploymentImage(namespace, deployment, container, image)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	log.Infof("%s", info)
}
