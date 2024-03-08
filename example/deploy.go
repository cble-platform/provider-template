package example

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (provider ProviderExample) DeployResource(ctx context.Context, request *pgrpc.DeployResourceRequest) (*pgrpc.DeployResourceReply, error) {
	return nil, fmt.Errorf("not implemented")
}
