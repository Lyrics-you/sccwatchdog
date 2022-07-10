package cmd

import (
	"errors"
	"sccwatchdog/dog"

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
			dog.ShowAllDeployments(namespace)
		} else {
			dog.ShowDeployments(namespace, deployment)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	showCmd.PersistentFlags().StringVarP(&deployment, "deployment", "d", "", "scc deployment")
}
