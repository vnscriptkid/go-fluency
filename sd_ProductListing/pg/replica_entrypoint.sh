#!/bin/bash
set -e

pg_basebackup -h db -U repuser -D /var/lib/pg/data -Fp -Xs -P -R

chown -R postgres:postgres /var/lib/pg/data

# user read, write, and execute
chmod -R 0700 /var/lib/pg/data

su - postgres -c "/usr/lib/postgresql/14/bin/postgres -c config_file=/var/lib/pgcfg/data/replica.conf"

