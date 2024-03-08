package example

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (provider ProviderExample) Configure(ctx context.Context, request *pgrpc.ConfigureRequest) (*pgrpc.ConfigureReply, error) {
	return nil, fmt.Errorf("not implemented")
}
