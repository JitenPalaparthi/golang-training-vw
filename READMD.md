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

-- go env

- GOROOT -> Root go path, installation 
- GOPATH -> all non std dependedncies are stored here
- GOBIN  -> all tools locally or from internet are installed in this directory.If not setup automatically $GOPATH/bin   

```bash 
go install golang.org/x/tools/gopls@latest
# to install your application, it installs into the directory that is set for GOBIN
go install .
```

- list of os and arch that go supports for cross compilation

```bash
go tool dist list
```

- GOOS 
- GOARCH

- to do cross compilation

```bash
GOOS=linux GOARCH=arm64 go build -o demo_linx_arm64 main.go
```

- To build a release binary

```bash
go build -ldflags="-s -w" -o demo_release .

GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o demo_mim_linux_arm64 main.go


```

- To pull package

```bash
go get github.com/JitenPalaparthi/shapes-pacakge-demo@v1.0.0
```