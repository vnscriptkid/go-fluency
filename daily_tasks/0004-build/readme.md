## Targets
- Learn how to build go program for different platforms

## Practice
- Use different options for GOOS and GOARCH
- Try running each built version on your local machine, what conclusion can you draw?
- Use file command to check the binary type
- How to check the current OS and Architecture of your machine
  - `go env GOOS`
  - `go env GOARCH`

## Notes
- 0004: Mach-O 64-bit executable x86_64
- 0004.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
- 0004: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, Go BuildID=PeY1wdumGHM2LRIVGQYP/Ez9BNMUhd7lN1nvpsD8v/1UEjwqx8Jv8tY-dBZNeu/p43jWuUDzA3WwrFfNLJ2, with debug_info, not stripped

- GOOS=linux GOARCH=arm go build -o 0004-linux-arm
- GOOS=linux GOARCH=arm64 go build -o 0004-linux-arm64
- GOOS=linux GOARCH=amd64 go build -o 0004-linux-amd64
- GOOS=linux go build -o 0004-linux
- GOOS=darwin go build -o 0004-darwin
- GOOS=darwin GOARCH=arm64 go build -o 0004-darwin-arm64