## Redis
- Sorted set
    - zadd key score "member"
    - `zadd rate-limiter:user123 100 "100"`
    - Get all members of a sorted set
    - `zrange rate-limiter:user123 0 -1`
    - Get all members of a sorted set with scores
    - `zrange rate-limiter:user123 0 -1 withscores`
    - Get members by score range
    - `zrangebyscore rate-limiter:user123 -inf 100`

## Why rate limiting?
- Prevent abuse (DoS, brute force, etc.)
- Protect limited resources (API, etc.)
- By dropping requests coming from a user that has exceeded a limit
    - Block requests before they reach the application code
    - There's still a cost to process the request to determine if it should be dropped. But it's much lower than processing the request in the application code.

## Refs:
- https://youtu.be/TCTkkVoY3-I