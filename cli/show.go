package cli

import (
	"sccwatchdog/dog"

	"github.com/desertbit/grumble"
)

func init() {
	App.AddCommand(&grumble.Command{
		Name: "show",
		Help: "Show information by namespace and deployments",
		Args: func(a *grumble.Args) {
			a.String("deployment", "sepcified deployment show infos", grumble.Default(""))
		},
		Flags: func(f *grumble.Flags) {
			f.String("d", "deployment", "", "deployment")
		},
		Run: func(c *grumble.Context) error {
			deployment := getValue(c, "deployment", "").(string)
			if deployment == "" {
				dog.ShowAllDeployments(namespace)
			} else {
				dog.ShowDeployments(namespace, deployment)
			}
			return nil
		},
	})
}
