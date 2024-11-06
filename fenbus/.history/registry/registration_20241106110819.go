package registry

type Registration struct {
	Servicename Servicename
	ServiceURL  string
}

type Servicename string

const (
	LogService = Servicename("LogService")
)
