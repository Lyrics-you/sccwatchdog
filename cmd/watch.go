package cmd

import (
	"errors"
	"sccwatchdog/dog"

	"github.com/spf13/cobra"
)

var (
	second int
	except string
)
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch changed and restarted by namespace and deployments",
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
			dog.WatchAllDeplolyments(namespace, except, second)
		} else {
			dog.WatchDeplolyments(namespace, deployment, except, second)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "scc namespace")
	watchCmd.PersistentFlags().StringVarP(&deployment, "deployment", "d", "", "scc deployment")
	watchCmd.PersistentFlags().StringVarP(&except, "except", "e", "", "expect deployment")
	watchCmd.PersistentFlags().IntVarP(&second, "second", "s", 0, "times interval (second)")
}
