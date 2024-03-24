# todo_app

## About the App
This repository will create a web application that is an API server for managing TODO tasks with authentication.

Eventually, the following endpoints will be implemented.

HTTP Method | Path         | Overview
----------|------------|--------------------------
POST     | /register | Register a new user
POST     | /login   | Obtain an access token with registered user information
POST     | /tasks   | Register tasks using an access token
GET      | /tasks   | List tasks using an access token
GET      | /admin   | Only users with administrative privileges can access

Docker Compose is used to launch the API server, MySQL, and Redis. The commands that will mainly be executed are pre-defined in the Makefile.
```zsh
$ make
build                Build docker image to deploy
build-local          Build docker image to local development
up                   Do docker compose up with hot reload
down                 Do docker compose down
logs                 Tail docker compose logs
ps                   Check container status
test                 Execute tests
dry-migrate          Try migration
migrate              Execute migration
generate             Generate codes
help                 Show options
```

### How to check operation
This is the procedure to check if the code in this repository can be executed locally.

#### Start the server
Create a Docker image in advance.
```zsh
$ make build-local
```
Use Docker Compose to start each service.
```zsh
$ make up
```
Perform a migration to MySQL.
```zsh
$ make migrate
```
Create a user.
```zsh
$ curl -X POST localhost:18000/register -d '{"name": "budou", "password":"test", "role":"admin"}'
{"id":37}
```
Register some tasks using the user's authentication information.
```zsh
$ curl -i -XPOST -H "Authorization: Bearer $(curl -XPOST localhost:18000/login -d '{"user_name": "budou", "password":"test"}' | jq ".access_token" | sed "s/\"//g")" localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100  1023  100   982  100    41   7756    323 --:--:-- --:--:-- --:--:--  8525
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 20 Jul 2022 17:21:03 GMT
Content-Length: 9

{"id":76}%

$ curl -XPOST -H "Authorization: Bearer $(curl -XPOST localhost:18000/login -d '{"user_name": "budou", "password":"test"}' | jq ".access_token" | sed "s/\"//g")" localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100  1023  100   982  100    41   8634    360 --:--:-- --:--:-- --:--:--  9560
{"id":77}%
```
Display the tasks, and if the registered tasks are displayed, it is working as expected.
```zsh
$ curl -XGET -H "Authorization: Bearer $(curl -XPOST localhost:18000/login -d '{"user_name": "budou", "password":"test"}' | jq ".access_token" | sed "s/\"//g")" localhost:18000/tasks | jq
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100  1023  100   982  100    41   8158    340 --:--:-- --:--:-- --:--:--  9133
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100   113  100   113    0     0  13450      0 --:--:-- --:--:-- --:--:-- 28250
[
{
"id": 76,
"title": "Implement a handler",
"status": "todo"
},
{
"id": 77,
"title": "Implement a handler",
"status": "todo"
}
]
```

[//]: # (## Reference)

[//]: # (### Detailed Go Language Web Application Development)

[//]: # (##### Written by Yoichiro Shimizu)

[//]: # (https://amzn.asia/d/6VU2VFF)
