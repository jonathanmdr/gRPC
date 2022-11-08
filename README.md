# gRPC
Example of gRPC with Go

## Compiling `.proto` files:
---
> `--go_out` specifies out directory for generation all entities "_messages of `.proto` file_"
>
> `--go-grpc_out` specifies out directory for generation all gRPC code "_services of `.proto` file_"
>
> `latest param` specifies directory for get `.proto` file
```sh
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

## Installing client gRPC:
---
> The client used for tests is named [evans](https://github.com/ktr0731/evans)
```sh
# MacOS installation using brew
brew tap ktr0731/evans
brew install evans
```

## Creating SQLite database:
---
```sh
# For connect on database
sqlite3 db.sqlite
```
```sql
-- For create table categories
create table categories(
    id string primary key,
    name string,
    description string
);

-- For create table courses
create table courses(
    id string primary key,
    name string,
    description string,
    category_id string,
    foreign key (category_id) references categories(id)
);
```

## Running server:
---
```sh
go run cmd/grpcServer/main.go
```

## Running client:
---
> By default the evans client listens on TCP port 50051, you don't need to specify if you are using this port
```sh
evans -r repl
```

## Select gRPC service on client:
---
> Needs to start client before selecting gRPC service
>
> `first param` is the gRPC service name
```sh
# If your function uses stream mode, for stop this use 'Ctrl + D'
service CategoryService
```

## Calling gRPC service function on client:
---
> Needs to start client before selecting gRPC service
>
> `first param` is the function name
```sh
call CreateCategory
```
