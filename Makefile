check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger-generate: check-swagger
	GO111MODULE=off swagger generate spec -o ./docs/swagger.json --scan-models

swagger-serve: check-swagger
	swagger serve -F=swagger docs/swagger.json

run:
	go run cmd/api/main.go

run-docker-mysql:
	docker run --name golang-training-ground-mysql -e MYSQL_ROOT_PASSWORD=test -p 3306:3306 -d mysql:latest

migrate-up:
	migrate -path internal/championship/infrastructure/database/migrations/ -database "mysql://root:test@tcp(localhost:3306)/golang_playground" -verbose up

migrate-down:
	migrate -path internal/championship/infrastructure/database/migrations/ -database "mysql://root:test@tcp(localhost:3306)/golang_playground" -verbose down

migrate-drop:
	migrate -path internal/championship/infrastructure/database/migrations/ -database "mysql://root:test@tcp(localhost:3306)/golang_playground" -verbose drop