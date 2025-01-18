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
	@cd demo/demo_thrift && \
	cwgo server -I ../../idl \
	--type rpc \
	--module github.com/lgxyc/gomall/demo/demo_thrift \
	--service demo_thrift \
	--idl ../../idl/echo.thrift


.PHONY: app-frontend-fix
app-frontend-fix:
	@cd app/frontend && golanci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/auth.proto \
	--service frontend \
	-module github.com/lgxyc/gomall/app/frontend \
	-I ../../idl

.PHONY: gen-frontend-home
gen-frontend-home:
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/home.proto \
	--service frontend \
	-module github.com/lgxyc/gomall/app/frontend \
	-I ../../idl

.PHONY: gen-rpc-gen_user-client
gen-rpc-gen-user-client:
	@cd  rpc_gen && \
	cwgo client -I ../idl \
	--type rpc  \
	--module github.com/lgxyc/gomall/rpc_gen \
	--service user \
	--idl ../idl/user.proto
	
.PHONY: gen-rpc-gen_user-server
gen-rpc-gen-user-server:
	@cd app/user && \
	cwgo server --type rpc \
	--idl ../../idl/user.proto \
	--service user \
	--module github.com/lgxyc/gomall/app/user \
	-I ../../idl \
	--pass "-use github.com/lgxyc/gomall/rpc_gen/kitex_gen"

	