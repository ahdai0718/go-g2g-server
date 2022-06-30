package platform

import (
	"ohdada/g2gserver/internal/pkg/pb"
)

// providerFactoryDefault .
type providerFactoryDefault struct{}

// Create .
func (factory *providerFactoryDefault) Create(platformProvider *pb.PlatformProvider) (provider Provider) {

	switch platformProvider.Name {
	default:
		provider = &ProviderDefault{}
	}

	return
}
