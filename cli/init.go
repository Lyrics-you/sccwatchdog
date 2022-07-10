package cli

import (
	"fmt"
	"sccwatchdog/model"

	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

var (
	namespace string
)

var App = grumble.New(&grumble.Config{
	Name:                  "sccwatchdog",
	Description:           "sccwatchdog is a scc deployment watcher system.",
	HistoryFile:           "/tmp/watchdog.hist",
	Prompt:                "watchdog » ",
	PromptColor:           color.New(color.FgBlue, color.Bold),
	HelpHeadlineColor:     color.New(color.FgBlue),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})

func init() {
	version := model.Historys[len(model.Historys)-1]
	App.SetPrintASCIILogo(func(a *grumble.App) {
		println(model.LOGO)
		println(model.EMOJI["sccwatchdog"], version.Version)
	})
	namespace = "default"
	App.SetPrompt(fmt.Sprintf("watchdog [%s] » ", namespace))
}
