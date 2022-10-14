package demomod_test

import (
	"context"
	"fmt"
	"testing"

	demomod "moul.io/adapterkit-module-demo"
)

func TestPackage(t *testing.T) {
	var _ demomod.DemomodSvcServer = &demomod.Service{}
}

func ExampleDemomodSvcSum_basic() {
	svc := demomod.Service{}
	ctx := context.TODO()
	ret, err := svc.Sum(ctx, &demomod.Sum_Request{A: 42, B: 1000})
	fmt.Println(ret, err)

	// Output:
	// c:1042 <nil>
}
