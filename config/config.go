package config

//supported protocols
const (
	HTTPProto  = "http"
	HTTPSProto = "https"
)

type BackendService struct {
	UUID      string
	Endpoints []Endpoint
	Path      string
	Host      string
	Algorithm string
	Port      int
}

type Endpoint struct {
	IP   string
	Port int
}

type FrontendService struct {
	DefaultCert     *Certificate
	Name            string
	Port            int
	BackendServices []*BackendService
	Protocol        string
}

type LoadBalancerConfig struct {
	Name             string
	Scale            int
	FrontendServices []*FrontendService
	Config           string
}

type Certificate struct {
	Name  string
	Cert  string
	Key   string
	Fetch bool
}
