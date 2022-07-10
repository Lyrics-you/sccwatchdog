package cli

import (
	"fmt"

	"github.com/desertbit/grumble"
)

func init() {
	App.AddCommand(&grumble.Command{
		Name:    "select",
		Help:    "Select the namespace",
		Aliases: []string{"into"},
		Args: func(a *grumble.Args) {
			a.String("namespace", "sepcified namespace", grumble.Default("default"))
		},
		Run: func(c *grumble.Context) error {
			namespace = c.Args.String("namespace")
			c.App.SetPrompt(fmt.Sprintf("watchdog [%s] » ", namespace))
			return nil
		},
	})
}
