package example

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (provider ProviderExample) ResourcePower(ctx context.Context, request *pgrpc.ResourcePowerRequest) (*pgrpc.ResourcePowerReply, error) {
	return nil, fmt.Errorf("not implemented")
}
