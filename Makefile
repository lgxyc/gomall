.PHONY: gen-demo-proto
gen-demo-proto:
	@cd  demo/demo_proto && \
	cwgo server -I ../../idl \
	--type rpc  \
	--module github.com/lgxyc/gomall/demo/demo_proto \
	--service demo_proto \
	--idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift &&
	cwgo server 
	--type rpc\
	--module github.com/lgxyc/gomall/demo/demo_thrift \
	--service demo_thrift \
	--idl ../../idl/echo.thrift
.PHONY: demo_proto
demo_proto:
	@cd demo/demo_proto && go work use . && go run .