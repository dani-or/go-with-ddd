# go-with-ddd
Hello world with ddd


## Pre-requisites
##- Kubectl —  https://kubernetes.io/docs/tasks/tools/install-kubectl/
##- AWS CLI -  https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
##- Aws iam authenticator — https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html
##- eksctl — https://github.com/weaveworks/eksctl


## Create VPC for EKS using stackset -- AWS CLI
aws cloudformation deploy --template-file red.yaml --stack-name my-new-stack

## Create our cluster on EKS
eksctl create cluster -f cluster.yaml --kubeconfig=C:\Users\Lenovo\.kube\config

## Crear el policy
aws iam create-policy --policy-name go-with-ddd-role-pilicy --policy-document file://policy.json

## Crear el role con el que se va ejecutar el pod
eksctl create iamserviceaccount \
  --cluster=EKS-Demo-Cluster1 \
  --role-name=go-with-ddd-role-1 \
  --namespace=default \
  --name=go-with-ddd-service-account-name-1  \
  --attach-policy-arn=arn:aws:iam::851560454673:policy/go-with-ddd-role-pilicy \
  --approve


## Create our deployment
##Cuando en el selector hay unos label eso quiere decir que ese servio aplica 
##a todos los deployments con ese label
kubectl apply -f deployment.yaml

## Create service
kubectl apply -f service.yaml
