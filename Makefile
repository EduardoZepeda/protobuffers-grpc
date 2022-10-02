.PHONY: runserver/student
runserver/student:
	@echo 'Running student server'
	go run server-student/main.go

.PHONY: runserver/test
runserver/test:
	@echo 'Running test server'
	go run server-test/main.go

.PHONY: runserver/database
runserver/database:
	@echo 'Running docker container'
	docker run -p 54321:5432 go-grpc-db

.PHONY: compile/test
compile/test:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative testpb/test.proto

.PHONY: compile/student
compile/student:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative studentpb/student.proto

.PHONY: compile/database
compile/database:
	@echo 'Compiling docker container'
	docker build database/ -t go-grpc-db