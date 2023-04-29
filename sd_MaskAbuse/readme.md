## Setup
- Run server:
    - `go run .`

- Client 1
    - `telnet localhost 8888`
    - `/nick jack`
    - `/join room1`
    - `/msg hi everyone`

- Client 1
    - `telnet localhost 8888`
    - `/nick marry`
    - `/join room1`
    - `/msg hi jack`

## TODO
- Measure perf btw naive way and trie way