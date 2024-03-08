package example

import (
	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

type ProviderExample struct {
	pgrpc.DefaultProviderServer
}

const (
	name        = "provider-template"
	description = "Just an example provider"
	author      = "YOUR NAME <github.com/YOUR-USERNAME>"
	version     = "v0.0.0"
)

var CONFIG *ProviderExampleConfig

func (provider *ProviderExample) Name() string {
	return name
}

func (provider *ProviderExample) Description() string {
	return description
}

func (provider *ProviderExample) Author() string {
	return author
}

func (provider *ProviderExample) Version() string {
	return version
}
