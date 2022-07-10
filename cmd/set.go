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
		dog.SetDeploymentImage(namespace, deployment, container, image)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	setCmd.PersistentFlags().StringVarP(&deployment, "deployment", "d", "", "scc deployment")
	setCmd.PersistentFlags().StringVarP(&container, "container", "c", "", "deployment container")
	setCmd.PersistentFlags().StringVarP(&image, "image", "i", "", "container image")
}
