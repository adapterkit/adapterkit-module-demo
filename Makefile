test: generate
	go mod tidy
	go test -v .
	@echo "Done."

generate: demomod.pb.go
.PHONY: generate

demomod.pb.go: demomod.proto
	protoc --go_out=./ --go-grpc_out=./ $<

clean:
	rm -f demomod.pb.goa
