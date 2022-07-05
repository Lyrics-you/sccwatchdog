package cmd

import (
	"errors"
	"sccwatchdog/logger"

	"github.com/spf13/cobra"
)

var (
	log        = logger.Logger()
	namespace  string
	deployment string
	container  string
	image      string
	version    bool
)

var rootCmd = &cobra.Command{
	Use:   "swd",
	Short: "sccwatchdog is a scc deployment watcher system.",
	Long: `
Sccwatchdog is a free and open source scc deployment watcher system,
designed to watch scc deployment image version and updatetime with speed and efficiency.
You can also restart deployment and set container image by it.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Error(cmd, args, errors.New("unrecognized command"))
			return
		}
		showVersion()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "sccwatchdog version")
}
func Execute() {
	rootCmd.Execute()
}
