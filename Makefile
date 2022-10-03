## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## runserver/student: Run the student server in port 5060
.PHONY: runserver/student
runserver/student:
	@echo 'Running student server'
	go run server-student/main.go

## runserver/test: Run the test server in port 5061
.PHONY: runserver/test
runserver/test:
	@echo 'Running test server'
	go run server-test/main.go

## runserver/database: Run the docker database server on port 54321
.PHONY: runserver/database
runserver/database:
	@echo 'Running docker container'
	docker run -p 54321:5432 go-grpc-db

## compile/test: Compile test.proto into go code
.PHONY: compile/test
compile/test:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative testpb/test.proto

## compile/student: Compile student.proto into go code
.PHONY: compile/student
compile/student:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative studentpb/student.proto

## compile/database: Build docker Postgres' image
.PHONY: compile/database
compile/database:
	@echo 'Compiling docker container'
	docker build database/ -t go-grpc-db