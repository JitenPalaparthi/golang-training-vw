## Go module 

```bash

go mod init demo

```

- To run 

```bash

go run main.go

# Two files of main package 

go run main.go greet.go

go run .

```


- To buold 

```bash
go build -o demo main.go greet.go hello.go 

go buold . 

go build -o demo .

```