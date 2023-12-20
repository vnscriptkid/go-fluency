## Ref
- https://coursehunter.net/course/vhod-v-kubernetes-s-pomoshchyu-efk-stack-polnoe-rukovodstvo?lesson=4

## java setups
- java -version
- javac -version
- https://www.oracle.com/sg/java/technologies/downloads/#jdk20-mac
    - M1 chip -> ARM
    - Intel chip -> x64
- Set env: JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk-1.8.jdk/Contents/Home

## EKS setups
- ap-southeast-1
- aws configure --profile eks_personal
```sh
aws eks --region ap-southeast-1 --profile eks_personal list-clusters
{
    "clusters": [
        "log-app-cluster"
    ]
}

aws eks --region ap-southeast-1 update-kubeconfig --name log-app-cluster --profile eks_personal
```
```

#### eksctl
brew tap weaveworks/tap
brew install weaveworks/tap/eksctl