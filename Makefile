dependency:
	go mod tidy && go mod vendor

generate:
	go generate -x ./...
