package example

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (provider ProviderExample) RetrieveData(ctx context.Context, request *pgrpc.RetrieveDataRequest) (*pgrpc.RetrieveDataReply, error) {
	return nil, fmt.Errorf("not implemented")
}
