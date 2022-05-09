# Backend
Backend application, connects to mysql db and uses graphql to send responses to frontend. There is also REST api but not really needed for frontend applications.

A lot of concepts already have existing packages, however I want to do most of it on my own so I can learn.

## Start up
```bash
cp .env.example .env
go mod download
go build -o server server.go
go run server.go
```

## Testing
```bash
go test ./tests --v
```

## Lint
```bash
#TODO
```

## Environment Variables
- `ENV` - what environment you are in (test, development, production)
- `MYSQL_DSN` - url to connect to mysql db
- `JWT_SECRET` - hashing secret for json web token
- `AWS_ACCESS_KEY_ID` - Used to access aws resource(s)
- `AWS_SECRET_ACCESS_KEY` - Used to access aws resource(s)
- `AWS_REGION` - Region of aws resource(s)
- `S3_BUCKET` - S3 bucket name

## Running backend scripts
```bash
cd app
go run server.go help
go run server.go help <command> 
```

Commands
```bash
go run server.go create_user <username> <password> <email>
go run server.go migrate
```

## Endpoints
| Title | Endpoint | Method | Params |
| --- | --- | --- | --- | 
| Login | `api/login` | POST | **Required**<br>`username`<br/>`password` |
| List Projects | `api/projects` | GET | N/A |
| Create Project | `api/project/create` | POST | **Required**<br>`title`<br/>`description`<br/>**Optional**<br/>`url`<br/>`git`<br/>`tags` |
| Edit Project | `api/project/edit` | POST | **Required**<br/>`id`<br>`title`<br/>`description`<br/>**Optional**<br/>`url`<br/>`git`<br/>`tags` |
| Delete Project | `api/project/delete` | POST | **Required**<br>`id` |
| List Tags | `api/tags` | GET | N/A |
| Create Tag | `api/tag/create` | POST | **Required**<br>`title` |
| Edit Tag | `api/tag/edit` | POST | **Required**<br>`id`<br>`title` |
| Delete Tag | `api/tag/delete` | POST | **Required**<br>`id` |

## Elevated Endpoint
TODO: require permission to execute
| Title | Endpoint | Method | Params |
| --- | --- | --- | --- | 
| List User | `api/users` | GET | N/A |
| Create User | `api/user/create` | POST | **Required**<br>`username`<br>`password`<br>`email` |
| Edit User | `api/user/edit` | POST | **Required**<br>`id`<br>`username`<br>`password`<br>`email` |
| Delete User | `api/user/delete` | POST | **Required**<br>`id` |

# Getting JWT to login
Tokens can be supplied in multiple ways, these will be checked in respective order and if one is not nil then it will be used to check if valid or not. (It will not check all three). To get the token use endpoint `api/login` with username and password
- Cookie `dylank-io-auth`
- `token` in body
- `Authentication` in header

# Graphql API
This project supports graphql api extensively, and should be used for frontend development where possible. Graphql Endpoints are intended to be open, no authentication needed. These should be limited to `query` operation

Get Projects
```graphql
query {
  projects(tags: $tags, offset: $offset, limit: $limit) {
    items {
      id
      title
      slug
      assets {
        id
        slug
        title
        createdAt
        updatedAt
        url
      }
      description
      createdAt
      updatedAt
      url
      git
      tags {
        id
        slug
        title
        createdAt
        updatedAt
      }
    }
    tags {
      id
      slug
      title
      createdAt
      updatedAt
      count
    }
    total
    items_total
  }
}
```

Get Tags
```graphql
query {
  tags {
    title
    slug
  }
}
```
