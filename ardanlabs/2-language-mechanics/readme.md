# Language mechanicss

## Keywords
- Precision types: int32, int64
- Non-precision types: int, uint
- Wordsize: 32, 64
- Same architecture: size(int) == size(pointer) == size(wordsize)
- Assign 2 structs of different types: Must have same fields and types
- Assign untyped to typed: don't need to cast
- Stack
    - Each goroutine (app thread) manages 1 stack, initial size is 2048 bytes
    - Return?
- Escape analysis
    - Func varialbe should be on stack or heap
    - Assumption: On return, stack frame is not cleaned up and will be reused next call?!?
    - Define as struct and return the pointer to that struct => on heap (why? if on stack, the next call would override shared memory)
    - Define as struct and return a copy => on stack (why? caller would have copy of struct, even next call it get overridden, it's fine)
- Stack growth
- Garbage Collection
- Constants
- IOTA

## Q
- When to use
```go
var x Person
x := Person{}
```
- Size of struct and why?
```go
type example struct {
    flag    bool
    counter int16
    pi float32 
}
```
- Print size of struct
```go
fmt.Println(unsafe.Sizeof(example{}))
```
