package utils

const (
	// all
	ItemsDeploymentName                   = "{.items[*].metadata.name}"
	ItemsDeploymentNamespace              = "{.items[*].metadata.namespace}"
	ItemsDeploymentImage                  = "{.items[*].spec.template.spec.containers[*].image}"
	ItemsDeploymentLastUpdateTime         = "{.items[*].status.conditions[-1].lastUpdateTime}"
	ItemsDeploymentMessage                = "{.items[*].status.conditions[-1].message}"
	ItemsDeploymentImageAndLastUpdateTime = "{range.items[*]}[{.spec.template.spec.containers[*].image},{.status.conditions[-1].lastUpdateTime}]{end}"
	ItemsDeploymentAllInfos               = "{range.items[*]}[{.metadata.name},{.spec.template.spec.containers[*].image},{.status.conditions[-1].lastUpdateTime}]{end}"
	// single
	DeploymentImage                  = "{.spec.template.spec.containers[*].image}"
	DeploymentContainerName          = "{.spec.template.spec.containers[*].name}"
	DeploymentLastUpdateTime         = "{.status.conditions[-1].lastUpdateTime}"
	DeploymentMessage                = "{.status.conditions[-1].message}"
	DeploymentImageAndLastUpdateTime = "{.spec.template.spec.containers[*].image},{.status.conditions[-1].lastUpdateTime}"
)
