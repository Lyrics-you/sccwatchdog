package dog

import (
	"fmt"
	"sccwatchdog/utils"
	"strings"
)

func setContainerImage(namespace, deployment, container, image string) (string, error) {
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
	line := fmt.Sprintf(`set image -n %s deployment.apps/%s %s=%s`, namespace, deployment, container, image)
	cmdsplit := strings.Split(line, " ")
	out, err := utils.KubeCommand(cmdsplit)
	if err != nil {
		return string(out), err
	}
	if string(out) == "" {
		return fmt.Sprintf("%s(%s) image not changed", deployment, container), nil
	}
	return string(out), nil
}

func SetDeploymentImage(namespace, deployment, container, image string) {
	info, err := setContainerImage(namespace, deployment, container, image)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	log.Infof("%s", info)
	ShowDeployments(namespace, deployment)
}
