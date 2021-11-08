# dylank.io
https://dylank.io

Front end is served through s3, and backend AWS Fargate.

## Startup
Create env file (only needed for backend app)
```bash
cp app/.env.example .env
```

Build + Start
```bash
yarn
docker-compose build
docker-compose up
```

Open http://localhost

## Running backend scripts
```bash
cd app
go run server.go help
go run server.go help <command> 
```

Commands
```bash
go run server.go create_user <username> <password> <email>
```

## Test
```bash
yarn test
```

## Lint
```bash
yarn lint
terraform -chdir=terraform fmt -check 
```

## Tech Stack
- Frontend - Typescript, React, Webpack, SCSS 
- API Gateway - Graphql
- Backend - Golang, Iris
- Database - MongoDB
- Infra - AWS, Terraform, Docker, Nginx 

## Credits
- Penguin Model - https://sketchfab.com/ZhangWanqing
