# Api for Backend application

| Endpoint | Method | Params | Description | 
| --- | --- | --- | --- |


# Getting JWT to login
Tokens can be supplied in multiple ways, these will be checked in respective order and if one is not nil then it will be used to check if valid or not. (It will not check all three). To get the token use endpoint `api/login` with username and password
- Cookie `dylank-io-auth`
- `token` in body
- `Authentication` in header