package cli

import (
	"sccwatchdog/dog"

	"github.com/desertbit/grumble"
)

func init() {
	App.AddCommand(&grumble.Command{
		Name: "restart",
		Help: "Restart by namespace and deployment",
		Args: func(a *grumble.Args) {
			a.String("deployment", "sepcified deployment show infos", grumble.Default(""))
			a.String("container", "deployment container", grumble.Default(""))
		},
		Flags: func(f *grumble.Flags) {
			f.String("d", "deployment", "", "deployment")
			f.String("c", "container", "", "deployment container")
		},
		Run: func(c *grumble.Context) error {
			deployment := getValue(c, "deployment", "").(string)
			container := getValue(c, "container", "").(string)
			dog.RestartDeployment(namespace, deployment, container)
			return nil
		},
	})
}
