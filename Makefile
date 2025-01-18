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

.PHONY: gen-link-fix
demo-link-fix:
	@cd demo/demo_proto && golanci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && \
	cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/lgxyc/gomall/app/frontend -I ../../idl

	