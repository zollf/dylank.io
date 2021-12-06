# Infrastructure
Uses ECS + Fargate, with 3 containers, nginx as reverse proxy, go backend and node frontend. MongoDB hosted on mongo atlas (TODO: move mongodb atlas config to terraform). 

You will require access to aws infrastructure. This requires `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`. 

Pipeline has some `iam` permissions to run different CRUD commands on certain services (TODO: restrict further where appropriate).

## Run
```bash
terraform init
terraform plan
terraform approve
```

## Lint
```bash
terraform fmt -check 
```