export REPO=github.com/lgxyc/gomall

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd  demo/demo_proto && \
	cwgo server -I ../../idl \
	--type rpc  \
	--module ${REPO}/demo/demo_proto \
	--service demo_proto \
	--idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && \
	cwgo server -I ../../idl \
	--type rpc \
	--module ${REPO}/demo/demo_thrift \
	--service demo_thrift \
	--idl ../../idl/echo.thrift


.PHONY: app-frontend-fix
app-frontend-fix:
	@cd app/frontend && golanci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m


.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/home.proto \
	--service frontend \
	-module ${REPO}/app/frontend \
	-I ../../idl
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/auth_page.proto \
	--service frontend \
	-module ${REPO}/app/frontend \
	-I ../../idl
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/product_page.proto \
	--service frontend \
	-module ${REPO}/app/frontend \
	-I ../../idl
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/category_page.proto \
	--service frontend \
	-module ${REPO}/app/frontend \
	-I ../../idl
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/cart_page.proto \
	--service frontend \
	-module ${REPO}/app/frontend \
	-I ../../idl
	@cd app/frontend && \
	cwgo server --type HTTP \
	--idl ../../idl/frontend/checkout_page.proto \
	--service frontend \
	-module ${REPO}/app/frontend \
	-I ../../idl


.PHONY: gen-user
gen-user:
	@cd  rpc_gen &&  cwgo client --type rpc \
	-I ../idl \
	--idl ../idl/user.proto \
	--module ${REPO}/rpc_gen \
	--service user 
	@cd app/user && \
	cwgo server --type rpc \
	-I ../../idl \
	--idl ../../idl/user.proto \
	--module ${REPO}/app/user \
	--service user \
	--pass "-use ${REPO}/rpc_gen/kitex_gen"

.PHONY: gen-product
gen-product:
	@cd  rpc_gen && \
	cwgo client --type rpc \
	-I ../idl \
	--idl ../idl/product.proto \
	--module ${REPO}/rpc_gen \
	--service product 
	@cd app/product && \
	cwgo server --type rpc \
	-I ../../idl \
	--idl ../../idl/product.proto \
	--module ${REPO}/app/product \
	--service product \
	--pass "-use ${REPO}/rpc_gen/kitex_gen"

.PHONY: gen-cart
gen-cart:
	@cd  rpc_gen && \
	cwgo client --type rpc \
	-I ../idl \
	--idl ../idl/cart.proto \
	--module ${REPO}/rpc_gen \
	--service cart 
	@cd app/cart && \
	cwgo server --type rpc \
	-I ../../idl \
	--idl ../../idl/cart.proto \
	--module ${REPO}/app/cart \
	--service cart \
	--pass "-use ${REPO}/rpc_gen/kitex_gen"

.PHONY: gen-payment
gen-payment:
	@cd  rpc_gen && \
	cwgo client --type rpc \
	-I ../idl \
	--idl ../idl/payment.proto \
	--module ${REPO}/rpc_gen \
	--service payment 
	@cd app/payment && \
	cwgo server --type rpc \
	-I ../../idl \
	--idl ../../idl/payment.proto \
	--module ${REPO}/app/payment \
	--service payment \
	--pass "-use ${REPO}/rpc_gen/kitex_gen"

.PHONY: gen-checkout
gen-checkout:
	@cd  rpc_gen && \
	cwgo client --type rpc \
	-I ../idl \
	--idl ../idl/checkout.proto \
	--module ${REPO}/rpc_gen \
	--service checkout 
	@cd app/checkout && \
	cwgo server --type rpc \
	-I ../../idl \
	--idl ../../idl/checkout.proto \
	--module ${REPO}/app/checkout \
	--service checkout \
	--pass "-use ${REPO}/rpc_gen/kitex_gen"