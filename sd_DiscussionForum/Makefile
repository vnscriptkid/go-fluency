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

start:	
	echo "Enable sharding and setup sharding-key"
	docker-compose exec router01 mongosh --port 27017 --eval " \
	use mydb;
	sh.enableSharding("mydb"); \
	db.adminCommand({ shardCollection: "mydb.posts", key: { _id: "hashed" }}); \
	db.adminCommand({ shardCollection: "mydb.comments", key: { post_id: "hashed" }}); \
	"

down:
	docker compose down --volumes

cli:
	docker-compose exec router01 mongosh --port 27017

# use mydb;
# db.posts.getShardDistribution(); 

shard1:
	docker-compose exec shard01-a mongosh --port 27017

# use mydb;
# db.posts.find().pretty()
# db.posts.count()

shard2:
	docker-compose exec shard02-a mongosh --port 27017

shard3:
	docker-compose exec shard03-a mongosh --port 27017