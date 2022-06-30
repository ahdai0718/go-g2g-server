package platform

import (
	"ohdada/g2gserver/internal/pkg/pb"
)

var (
	defaultProviderFactorySimpleFactory = &providerFactorySimpleFactory{}
)

// CreateProvider .
func CreateProvider(platformProvider *pb.PlatformProvider) Provider {
	factory := defaultProviderFactorySimpleFactory.Create(platformProvider.FactoryName)
	provider := factory.Create(platformProvider)
	return provider
}
