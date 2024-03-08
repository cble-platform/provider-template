package example

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (provider ProviderExample) ExtractResourceMetadata(ctx context.Context, request *pgrpc.ExtractResourceMetadataRequest) (*pgrpc.ExtractResourceMetadataReply, error) {
	return nil, fmt.Errorf("not implemented")
}
