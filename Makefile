
run: compile
	./cmd/compiled.exe

compile:
	go build -o cmd/compiled.exe
cmdtes:
	cmd/generate_proto.sh

groto:
	protoc --proto_path=./proto --go_out=. --go-grpc_out=. ./proto/$(target)/*.proto

spes-test:
	go test -run ^TestUserAdd$ github.com/private-project-pp/product-rpc-service/usecase/users_administration -v

gen-mock-all: gen-mock-repo

gen-mock-repo:
	cmd/generate_mocked_repo.sh

