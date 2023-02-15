check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger-generate: check-swagger
	GO111MODULE=on go mod vendor && GO111MODULE=off swagger generate spec -o ./docs/swagger.json --scan-models

swagger-serve: check-swagger
	swagger serve -F=swagger docs/swagger.json