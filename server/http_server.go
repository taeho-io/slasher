package server

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/taeho-io/idl/gen/go/slasher"
	slasherClient "github.com/taeho-io/slasher"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func ServeHTTP(addr string, _ Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := slasher.RegisterSlasherHandlerFromEndpoint(
		ctx,
		mux,
		slasherClient.ServiceURL,
		opts,
	); err != nil {
		return err
	}

	return http.ListenAndServe(addr, mux)
}
