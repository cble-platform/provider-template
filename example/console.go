package example

import (
	"context"
	"fmt"

	pgrpc "github.com/cble-platform/cble-provider-grpc/pkg/provider"
)

func (provider ProviderExample) GetConsole(ctx context.Context, request *pgrpc.GetConsoleRequest) (*pgrpc.GetConsoleReply, error) {
	return nil, fmt.Errorf("not implemented")
}
