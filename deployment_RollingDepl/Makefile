check:
	chmod +x ./check.sh \
	&& bash ./check.sh

port:
	kubectl port-forward svc/my-hello-app-service-nodeport 8080:8080

port-lb:
	kubectl port-forward svc/my-hello-app-service-loadbalancer 8080:80
