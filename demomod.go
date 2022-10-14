package demomod

import context "context"

type Service struct {
	UnimplementedDemomodSvcServer
}

func (svc Service) Sum(ctx context.Context, req *Sum_Request) (*Sum_Result, error) {
	a := req.GetA()
	b := req.GetB()
	sum := a + b
	ret := Sum_Result{C: sum}
	return &ret, nil
}

func (svc *Service) EchoStream(stream DemomodSvc_EchoStreamServer) error {
	panic("not implemented")
}

func (svc *Service) SayHello(ctx context.Context, _ *Empty) (*HelloResult, error) {
	ret := HelloResult{Msg: "hello!"}
	return &ret, nil
}
