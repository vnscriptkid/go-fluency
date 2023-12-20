# Tiner Feed System Design

## Requirements
- Find users nearby (geopartial database)
- Find users who share same interests (profile, auth svc)
- Won't see a user twice (set, bloomfilter)
- Swipe left/right with low letency (pre-gen, cache)

## Services
- Location tracker
    - Knows where user located on earth based on long/lat
        - Users proactively send location each 30s
    - Can answer who are the users nearby
    - DB: Redis
        - Load (write/reads) can be high -> use case for sharding

## Geopartial playground
- Hanoi: Long 105.804817 / Lat 21.028511
    - `GEOADD locations 105.804817 21.028511 "hanoi"`
- HCM: Long 106.660172 / Lat 10.762622
    - `GEOADD locations 106.660172 10.762622 "hcm"`
- Get the distance between two locations
    - `GETDIST locations "hanoi" "hcm" km`
- Get the locations within a certain radius of a given location
    - `GEORADIUS locations 105.804817 21.028511 1000 km`

## Mongodb sharding
- https://github.com/minhhungit/mongodb-cluster-docker-compose/tree/master

## TODOs
- Measure space between set and bloomfilters as more data is added

