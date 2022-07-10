package cli

import (
	"fmt"
	"sccwatchdog/dog"

	"github.com/desertbit/grumble"
)

var (
	deployment string
	except     string
	second     int
)

func init() {
	App.AddCommand(&grumble.Command{
		Name: "watch",
		Help: "Watch changed and restarted by namespace and deployments",
		Args: func(a *grumble.Args) {
			a.String("deployment", "scc deployment", grumble.Default(""))
			a.String("except", "except deployment", grumble.Default(""))
			a.Int("second", "times interval (second)", grumble.Default(0))
		},
		Flags: func(f *grumble.Flags) {
			f.String("d", "deployment", "", "deployment")
			f.String("e", "except", "", "except deployment")
			f.Int("s", "second", 0, "times interval (second)")
		},
		Run: func(c *grumble.Context) error {
			deployment := getValue(c, "deployment", "").(string)
			except := getValue(c, "except", "").(string)
			second := getValue(c, "second", 0).(int)
			fmt.Println(deployment, except, second)
			if deployment == "" {
				dog.WatchAllDeplolyments(namespace, except, second)
			} else {
				dog.WatchDeplolyments(namespace, deployment, except, second)
			}
			return nil
		},
	})
}
