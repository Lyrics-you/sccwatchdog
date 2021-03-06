package model

type WatchDog struct {
	Namespace  string
	Deployment string
}

type Deployment struct {
	Name           string
	Namespace      string
	LastUpdateTime string
	Image          string
	Message        string
}

type GetInfo struct {
	Info  string
	Error error
}

type GetList struct {
	List  []string
	Error error
}

type History struct {
	Version     string
	Description string
}

type Image struct {
	Name    string
	Version string
}
