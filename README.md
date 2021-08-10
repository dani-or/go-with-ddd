# go-with-ddd
Este proyecto contiene 2 pruebas de concepto:
* healthcheck : /health, esta prueba responde un string `everything is ok!`
* servicio rest con acceso a Dynamo: /getcredit, esta prueba consulta la tabla `credit-customer-product-qa` y responde la información que se tenga guardada actualmente en la tabla con los parametros de entrada.

### Diagrama de Clases 
![Diagram de clases ](/resources/images/clases.png?raw=true "Diagrama de clases")

# Install local
Para instalar localmente el proyecto 
## Pre-requisites
##- Docker

## Construir la imagen 
```bash
docker image build -t go-with-ddd .
```
## Correr la imagen
```bash
docker run --network host -d go-with-ddd
```
## Probar los servicios apuntando al localhost
healthcheck
```bash
curl --location --request GET 'http://localhost:8080/health'
```
Servicio rest consultando dynamo
```bash
curl --location --request POST 'http://localhost:8080/getcredit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "debenture": "87100001189",
    "customerId": "CC-1026755666"
}'
```

# Install on EKS
Esta guia de instalación contiene:
- Crear artefactos de red
- Crear clúster
- Crear proveedor de identididad
- Crear de política
- Crear service account
- Crear del deployment
- Crear servicio

 Nota: La ejecución debe ser en orden


## Pre-requisites
##- Kubectl —  https://kubernetes.io/docs/tasks/tools/install-kubectl/
##- AWS CLI -  https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
##- Aws iam authenticator — https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html
##- eksctl — https://github.com/weaveworks/eksctl


## Crear artefactos de red -- AWS CLI
Se crean artefactos de red: vpc, subnets públicas y privadas, grupos de seguridad entre otros
```bash
aws cloudformation deploy --template-file red.yaml --stack-name my-new-stack
```

## Crear clúster
Para la creación del clúster se debe cambiar en el archivo cluster.yaml el id de las vpc y las subnets
```bash
eksctl create cluster -f cluster.yaml
```

## Crear proveedor de identididad
El proveedor de identidad se crea para que los pods puedan asumir un role de IAM y así acceder a los diferentes recursos de AWS, este paso es opcional para mejor documentación [aquí](https://dzone.com/articles/how-to-use-aws-iam-role-on-aws-eks-pods).
```bash
eksctl utils associate-iam-oidc-provider --cluster <CLUSTER_NAME> --approve
```

## Crear el policy
Este política se la vamos asignar al role que vamos a crear en el siguiente paso
```bash
aws iam create-policy --policy-name go-with-ddd-role-pilicy --policy-document file://policy.json
```

## Crear service account
Acá creamos el role que le vamos asigna a nuestro pod y le asociamos la política que creamos en el paso anterior

```bash
eksctl create iamserviceaccount \
  --cluster=EKS-Demo-Cluster1 \
  --role-name=go-with-ddd-role-1 \
  --namespace=default \
  --name=go-with-ddd-service-account-name-1  \
  --attach-policy-arn=arn:aws:iam::851560454673:policy/go-with-ddd-role-pilicy \
  --approve
```

## Crear del deployment
El despliegue incluye la arn del repositorio ecr donde tenemos la imagen que queremos desplegar, el puerto que queremos exponer, los labels con los que queremos clasificar nuestra app y el service account que creamos en el paso anterior que es el que nos permite que el o los pods asuman un rol de IAM
```bash
kubectl apply -f deployment.yaml
```

## Create service
El servico nos permite acceder al servicio a travé de internet, despues de ejecutado el comando podemos consultar los servicios con el comando `kubectl get services` y ahí podemos ver la ip externa de nuestro servicio para apuntarle desde internet

```bash
kubectl apply -f service.yaml
```

# Pruebas en EKS
healthcheck
```bash
curl --location --request GET 'http://acab85022b58943088e2476b80ef30f3-1519037493.us-east-1.elb.amazonaws.com:8080/health'
```
Servicio rest consultando dynamo
```bash
curl --location --request POST 'http://acab85022b58943088e2476b80ef30f3-1519037493.us-east-1.elb.amazonaws.com:8080/getcredit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "debenture": "87100001189",
    "customerId": "CC-1026755666"
}'
```