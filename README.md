# protobuffers-grpc

Enrollment system for students and courses using protobuffers and gRPC in go.

## Usage

For the purpose of learning process, database is created locally using docker container instance based on a Postgresql image.
Server reflection is active so you can monitor gRPC endpoints using tools like Postman.

### Run student server

This command sets up a student server for creating and retrieving a single student.

```
make runserver/student
```

### Run test server

This command sets up a test server for creating tests, questions, answers, student enrollment, taking tests and scores.

```
make runserver/test
```

### Run database server

Sets up a Postgres database running on port 5432. The database wipes out all data everytime it is initialized.

```
runserver/database
```

### Compile test protobuffer

This command compiles the test protobuffer to go code.

```
make compile/test
```

### Compile student protobuffer

This command compiles the student protobuffer to go code.

```
make compile/student
```

### Compile database docker image

This command compiles the docker image for the Postgres's database.

```
make compile/database
```

## Disclaimer

This code is an improved and augmented version of the "Curso de Go Avanzado: Protobuffers y gRPC" course on Platzi.