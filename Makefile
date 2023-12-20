generate:
	protoc -Iproto --go_out=. --go_opt=module=github.com/KhetwalDevesh/project-totality --go-grpc_out=. --go-grpc_opt=module=github.com/KhetwalDevesh/project-totality proto/users.proto

test-generate1:
	find proto/*/v1 -name "*.proto" -exec \
	protoc -I{} \
	--go_out=stubs --go_opt=module=github.com/KhetwalDevesh/grpc-comm \
	--go-grpc_out=stubs --go-grpc_opt=module=github.com/KhetwalDevesh/grpc-comm \
	{} \;

test-generate2:
	find proto/ -name "*.proto" -exec \
	protoc -I{} \
	--go_out=stubs --go_opt=module=github.com/KhetwalDevesh/grpc-comm \
	--go-grpc_out=stubs --go-grpc_opt=module=github.com/KhetwalDevesh/grpc-comm \
	{} \;

test-generate3:
	protoc -Iproto --go_out=stubs --go_opt=module=github.com/my-merch/gateway \
      --go-grpc_out=stubs --go-grpc_opt=module=github.com/my-merch/gateway proto/users.proto
