package demomod_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	demomod "moul.io/adapterkit-module-demo"
)

func TestInterface(t *testing.T) {
	var _ demomod.DemomodSvcServer = demomod.New("")
}

func ExampleDemomodSvcSum_basic() {
	svc := demomod.New("")
	ctx := context.TODO()
	ret, err := svc.Sum(ctx, &demomod.Sum_Request{A: 42, B: 1000})
	fmt.Println(ret, err)

	// Output: c:1042 <nil>
}

func ExampleDemomodSvcSayHello_basic() {
	svc := demomod.New("foo")
	ctx := context.TODO()
	ret, err := svc.SayHello(ctx, nil)
	fmt.Println(ret, err)

	// Output: msg:"hello! foo" <nil>
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	service := demomod.New("test")
	demomod.RegisterDemomodSvcServer(server, service)
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func ExampleDemomodSvcSum_grpc() {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := demomod.NewDemomodSvcClient(conn)
	req := &demomod.Sum_Request{A: 42, B: 1000}
	ret, err := client.Sum(ctx, req)
	fmt.Println(ret, err)

	// Output: c:1042 <nil>
}
