## Checking queries
```sql
SELECT rolname, rolreplication FROM pg_roles WHERE rolcanlogin;

-- from master
SELECT * FROM pg_stat_replication;

-- from replica
SELECT * FROM pg_stat_wal_receiver;
```

## Lessons learned
- Scaling APIs: Using LB
- Scaling DB reads: Using read-replicas (master-slaves arch)
- Master can take both read-writes
- Devs take care of choosing connections to master-slaves
- Redis helps scale reads but comes with invalidation headache