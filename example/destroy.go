package example

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (provider ProviderExample) DestroyResource(ctx context.Context, request *pgrpc.DestroyResourceRequest) (*pgrpc.DestroyResourceReply, error) {
	return nil, fmt.Errorf("not implemented")
}
