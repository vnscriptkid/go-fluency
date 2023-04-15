## Step-by-step
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