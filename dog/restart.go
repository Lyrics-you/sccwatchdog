package dog

import (
	"encoding/json"
	"fmt"
	"sccwatchdog/utils"
	"strings"
	"time"
)

type restart struct {
	Spec spec `json:"spec"`
}
type spec struct {
	Template template `json:"template"`
}
type template struct {
	Spec tspec `json:"spec"`
}
type tspec struct {
	Containers []containers `json:"containers"`
}
type containers struct {
	Name string `json:"name"`
	Env  []env  `json:"env"`
}
type env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func RestartDeployment(namespace, deployment, container string) (string, error) {
	if namespace == "" {
		namespace = "default"
	}
	if container == "" {
		containers, err := utils.GetDeploymentContainers(namespace, deployment)
		if err != nil {
			return "", err
		}
		container = containers[0]
		if container == "" {
			container = deployment
		}
	}
	// kubectl patch deployment -n mergemultx mergemultx1
	// -p '{"spec":{"template":{"spec":{"containers":[{"name":"mergemultx","env":[{"name":"RESTART_","value":"'$(date +%s)'"}]}]}}}}'
	// add useless env variable
	line := fmt.Sprintf(`patch deployment -n %s %s -p`, namespace, deployment)
	cmdsplit := strings.Split(line, " ")
	time := fmt.Sprintf("%d", time.Now().Unix())
	r := restart{Spec: spec{Template: template{Spec: tspec{Containers: []containers{{Name: container, Env: []env{{Name: "SWD_RESTART", Value: time}}}}}}}}
	pattern, _ := json.Marshal(r)
	cmdsplit = append(cmdsplit, string(pattern))
	out, err := utils.KubeCommand(cmdsplit)
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
