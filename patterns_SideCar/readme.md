## Setup
```bash
docker build -t vnscriptkid/helloworld:1.0 .

docker login

docker push vnscriptkid/helloworld:1.0

istioctl install --set profile=demo -y

kubectl label namespace default istio-injection=enabled

kubectl apply -f infra/helloworld-deployment.yaml

kubectl apply -f ./infra/samples/addons 

istioctl dashboard jaeger

istioctl dashboard kiali
```

- Cleanup: 
    - `kubectl delete mutatingwebhookconfiguration istio-sidecar-injector`
    - `istioctl manifest generate | kubectl delete -f -`
    - Check pods again: `kubectl get pods -n istio-system`
    - Check svc again: `kubectl get services -n istio-system`
    - `kubectl get mutatingwebhookconfigurations`
    - `kubectl delete mutatingwebhookconfigurations <name-of-the-istio-webhook>`

