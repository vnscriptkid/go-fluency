up:
	docker compose up -d
	sleep 1
	docker-compose exec configsvr01 sh -c "mongosh < /scripts/init-configserver.js"
	sleep 1
	docker-compose exec shard01-a sh -c "mongosh < /scripts/init-shard01.js"
	sleep 1
	docker-compose exec shard02-a sh -c "mongosh < /scripts/init-shard02.js"
	sleep 1
	docker-compose exec shard03-a sh -c "mongosh < /scripts/init-shard03.js"
	echo "Wait a bit for the config server and shards to elect their primaries before initializing the router"
	sleep 5
	echo "Initializing the router"
	docker-compose exec router01 sh -c "mongosh < /scripts/init-router.js"
	echo "Enable sharding and setup sharding-key"
	docker-compose exec router01 mongosh --port 27017 --eval " \
	use feed_db;
	sh.enableSharding("feed_db"); \
	db.adminCommand({ shardCollection: "feed_db.feeds", key: { user_id: "hashed" }}); \
	"


down:
	docker compose down

rediscli:
	docker exec -it tinder_location_tracker redis-cli

status:
	docker-compose exec router01 mongosh --port 27017 --eval " \
	sh.status(); \
	"

dbcli:
	docker-compose exec router01 mongosh --port 27017

stats:
	docker-compose exec router01 mongosh --port 27017 --eval " \
	use feed_db; \
	db.stats(); \
	"
