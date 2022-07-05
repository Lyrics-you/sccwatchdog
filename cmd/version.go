package cmd

import (
	"errors"
	"fmt"
	"sccwatchdog/model"

	"github.com/spf13/cobra"
)

var (
	isDesc bool
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version subcommand show swd version info",
	Long: `
eg.: swd version [-d]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Error(cmd, args, errors.New("unrecognized command"))
			return
		}
		if !isDesc {
			showVersion()
		} else {
			showDesciption()
		}
	},
}

func printTag(tag, value string) {
	fmt.Printf("%-13s : %s\n", tag, value)
}

func showVersion() {
	version := model.Historys[len(model.Historys)-1]
	printTag("Name", fmt.Sprintf("SCCWatchDog%s", model.EMOJI["sccwatchdog"]))
	printTag("Version", version.Version)
	printTag("Email", "Leyuan.Jia@Outlook.com")
}

func showDesciption() {
	version := model.Historys[len(model.Historys)-1]
	printTag("Name", fmt.Sprintf("SCCWatchDog%s", model.EMOJI["sccwatchdog"]))
	printTag("Version", version.Version)
	printTag("Description", version.Description)
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().BoolVarP(&isDesc, "description", "d", false, "history description")
}
