## Targets
- Create simple CLI app that allows getting input from user through CLI with no external libs
- Simple input validation

## Exercise
- Init go module
- Get hash lib from "github.com/alexellis/hmac"
- Read user input from CLI using `flag.StringVar` and `flag.Parse` and store in local variables
  - message
  - secret
- Do simple validation to make sure message and secret are not empty
- Sign message using `hmac` lib
- Print out the result using string interpolation
  - Make sure to use correct format for the result (especially the hash part)

## QnA
- Special env vars: GOPATH
  - What does this keep? TODO: experiment
- Special path: `/usr/local/go/bin`. What does this keep? TODO: experiment
- `go install lib@version` vs `go get lib@version`
  - DIFF
    - `go install` will install the lib to `$GOPATH/pkg/mod`
    - `go get` will install the lib to `$GOPATH/src`
  - Use cases
    - `go install` is used to install a lib that you want to use in your project
    - `go get` is used to install a lib that you want to read the source code
- Differentiate:
  - `$GOPATH/pkg` 
  - `$GOPATH/src`
  - `$GOPATH/bin`
- How to make binaries installed by `go install` available globally?
  - Add `$GOPATH/bin` to `$PATH`
- Differentiate:
  - `/usr/local/go/bin`: where go is installed
  - `$GOPATH/bin`: where binaries are installed by `go install`

- go get
  - use case: running inside a project to install lib as dependency
    - `go get` will clone the lib to `$GOPATH/pkg/mod`
    - update go.mod, go.sum
    - `/Users/thanhnguyen/code/go/pkg/mod/github.com/alexellis/hash-browns@v1.3.0`

- Differences:
  - go get inside go module
  - go get outside go module
  - go install outside go module
    - Clone src code into `/Users/thanhnguyen/code/go/pkg/mod/github.com/alexellis/hash-browns@v1.3.0`
    - Build binary into `/Users/thanhnguyen/code/go/bin/hash-browns`
  - go install inside go module (same as above)

- What does `go clean -modcache` do?
  - Clean the module cache at `$GOPATH/pkg/mod`
  - Global cache for all projects, not just the current project

- Given `go install lib@___`, what could be ___ ?
  - `latest`
  - `v1.0.0`
  - `commit-hash`