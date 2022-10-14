package demomod_test

import (
	"testing"

	demomod "moul.io/adapterkit-module-demo"
)

func TestPackage(t *testing.T) {
	var svc demomod.DemomodSvcServer = &demomod.Service{}
	_ = svc
}
