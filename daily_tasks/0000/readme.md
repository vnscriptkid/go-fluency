## Exercise
- Create new module under `$GOPATH/src/github.com/vnscriptkid/hello-go`
  - `mkdir -p $GOPATH/src/github.com/vnscriptkid/hello-go`
- Go to the module folder
  - `cd $GOPATH/src/github.com/vnscriptkid/hello-go`
- Init module
  - `go mod init github.com/vnscriptkid/hello-go`
  - OR: `go mod init`
- Implement simple hello world:
```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```
- Run the program
  - `go run .`
- Build the program
  - `go build -o hello`
- Run the binary
  - `./hello`
- Push to git-hub
  - `git init`
  - `git add .`
  - `git commit -m "Initial commit"`
  - `git remote add origin

- Install lib: `go install -v github.com/vnscriptkid/hello-go@latest`
- Check src code: `ls $GOPATH/pkg/mod`
- Check binary: `ls $GOPATH/bin`

- Version the lib
  - `git tag v1.0.0`
  - `git log --oneline`
  - `git push origin v1.0.0`

- Install again: `go install -v github.com/vnscriptkid/hello-go@v1.0.0`
- Check src code:
```bash
ls $GOPATH/pkg/mod/github.com/vnscriptkid 
hello-go@v0.0.0-20231220152744-b1bac253a63f
hello-go@v0.0.1
```
- Run the binary: `hello-go`
  - `$GOPATH/bin/hello-go`
  - Confirm new code has been built for the new version