package registry

type Registration struct {
	Servicename      Servicename
	ServiceURL       string
	RegistryService  []Servicename
	ServiceUpdateURL string
}

type Servicename string

const (
	LogService = Servicename("LogService")
)

type patchEntry struct {
	Name Servicename
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
