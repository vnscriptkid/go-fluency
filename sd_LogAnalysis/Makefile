build_node:
	docker build -t vnscriptkid/log-app:node-1.0 ./node-app

build_java:
	cd ./java-app-efk && ./gradlew build
	docker build -t vnscriptkid/log-app:java-1.0 ./java-app-efk

build_check:
	docker images | grep log-app

docker_push:
	docker push vnscriptkid/log-app:java-1.0
	docker push vnscriptkid/log-app:node-1.0