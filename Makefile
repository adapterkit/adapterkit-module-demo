generate:
	protoc --go_out=./ demomod.proto
.PHONY: generate
