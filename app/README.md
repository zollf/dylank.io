# Api for Backend application

| Title | Endpoint | Method | Params |
| --- | --- | --- | --- | 
| Login | `api/login` | POST | **Required**<br>`username`<br/>`password` |
| List Projects | `api/projects` | GET | N/A |
| Create Project | `api/project/create` | POST | **Required**<br>`title`<br/>`description`<br/>**Optional**<br/>`url`<br/>`git`<br/>`tags` |
| Delete Project | `api/project/delete` | POST | **Required**<br>`id` |
| List Tags | `api/tags` | GET | N/A |
| Create Tag | `api/tag/create` | POST | **Required**<br>`title` |
| Delete Tag | `api/tag/delete` | POST | **Required**<br>`id` |


# Getting JWT to login
Tokens can be supplied in multiple ways, these will be checked in respective order and if one is not nil then it will be used to check if valid or not. (It will not check all three). To get the token use endpoint `api/login` with username and password
- Cookie `dylank-io-auth`
- `token` in body
- `Authentication` in header

# Graphql API
This project supports graphql api extensively, and should be used for frontend development where possible.