package model

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
