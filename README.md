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
