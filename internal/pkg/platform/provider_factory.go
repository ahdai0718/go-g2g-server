package platform

import "ohdada/g2gserver/internal/pkg/pb"

// ProviderFactory .
type ProviderFactory interface {
	Create(platformProvider *pb.PlatformProvider) Provider
}

type providerFactorySimpleFactory struct{}

func (factory *providerFactorySimpleFactory) Create(name string) (providerFactory ProviderFactory) {

	switch name {
	default:
		providerFactory = &providerFactoryDefault{}
	}

	return
}
