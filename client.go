package slasher

import (
	"sync"

	"github.com/taeho-io/go-taeho/interceptor"
	"github.com/taeho-io/idl/gen/go/slasher"
	"google.golang.org/grpc"
)

const (
	ServiceURL = "slasher:80"
)

var (
	cm     = &sync.Mutex{}
	Client slasher.SlasherClient
)

func GetSlasherClient() slasher.SlasherClient {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client
	}

	// We don't need to error here, as this creates a pool and connections
	// will happen later
	conn, _ := grpc.Dial(
		ServiceURL,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			interceptor.ContextUnaryClientInterceptor(),
		),
	)

	cli := slasher.NewSlasherClient(conn)
	Client = cli
	return cli
}
