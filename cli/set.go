package cli

import (
	"sccwatchdog/dog"

	"github.com/desertbit/grumble"
)

func init() {
	App.AddCommand(&grumble.Command{
		Name: "set",
		Help: "Set image by namespace and deployment",
		Args: func(a *grumble.Args) {
			a.String("deployment", "scc deployment", grumble.Default(""))
			a.String("container", "deployment container", grumble.Default(""))
			a.String("image", "container image", grumble.Default(""))
		},
		Flags: func(f *grumble.Flags) {
			f.String("d", "deployment", "", "scc deployment")
			f.String("c", "container", "", "deployment container")
			f.String("i", "image", "", "container image")
		},
		Run: func(c *grumble.Context) error {
			deployment := getValue(c, "deployment", "").(string)
			container := getValue(c, "container", "").(string)
			image := getValue(c, "image", "").(string)
			if container != "" && image == "" {
				container, image = image, container
			}
			dog.SetDeploymentImage(namespace, deployment, container, image)
			return nil
		},
	})
}
