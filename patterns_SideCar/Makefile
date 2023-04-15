VERSION=1.2

build:
	docker build -t vnscriptkid/greeting:${VERSION} ./greet-svc
	docker build -t vnscriptkid/hello:${VERSION} ./hello-svc

push:
	docker push vnscriptkid/greeting:${VERSION}
	docker push vnscriptkid/hello:${VERSION}

deploy:
	kubectl apply -f ./infra/greeting-deployment.yaml
	kubectl apply -f ./infra/helloworld-deployment.yaml