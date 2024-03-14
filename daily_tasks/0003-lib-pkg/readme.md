## Targets
- Implement utilities like: `ls`, `pwd`
- Structure program into lib package and main package
  - `lib` package contains exported functions
  - `main` package contains `main` function that calls exported functions

## Exercise
- Init a new module
- Create a `lib` package
- Design API for `ls` and `pwd` utilities
  - `(result, error)` pattern for return value
  - Focus on high level design first
  - Leave implementation details for later
- Expecting we can run:
```bash
./util -n ls
./util -n pwd
```
- Implement main program
  - Read -n flag as command name
  - Do simple validation to make sure this is passed
  - Use `switch` to call corresponding function given command name
  - Print error if command name is not recognized (default case)
  - Print result
- Make sure program run as expected
- Implement `ls` and `pwd` functions
  - `ls` should list all files in current directory
  - `pwd` should print current working directory
- Last but not least
  - Try `go install` at root of project
  - Look at `$GOPATH/bin` to see if newly built binary is there
  - Try go up one level and run `./bin/0003 -n ls`. What conclusion can you draw?
    - pwd prints current working directory where command is run (not where binary is located)

## Quizz

## Extra
- Reading open source: https://github.com/alexellis/derek/