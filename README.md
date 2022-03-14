# dylank.io
🚧 https://dylank.io 🚧

## Startup
Create env file (only needed for backend)
```bash
cp backend/.env.example .env
```

Build + Start
```bash
# In backend
go mod download

# In frontend 
yarn

# In root
docker-compose build
docker-compose up
```

Open http://localhost

## Tech Stack
- Frontend - Typescript, React, Nextjs, SCSS 
- API Gateway - Graphql
- Backend - Golang, Iris, Gorm
- Database - MySql
- Infra - AWS (ECS + EC2), Terraform, Docker, Nginx 

