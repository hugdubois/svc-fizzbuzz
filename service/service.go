// service package is the simple fizzbuzz service
package service

var (
	name    = "svc-fizzbuzz"
	version = "latest" // injected with -ldflags in Makefile
)

// Service is the fizzbuzz service
type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewService return new fizzbuzz service
func NewService() *Service {
	return &Service{
		Name:    name,
		Version: version,
	}
}
