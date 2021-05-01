# Matcher

## Prerequisites

Using package `github.com/joho/godotenv` to load `.env` file from directory. Need to create a `.env` file in root of the project and add the following environment variables:

```
POSTGRES_PASSWORD=password
POSTGRES_USER=postgres
POSTGRES_PORT=5432
POSTGRES_HOST=pg
POSTGRES_DB=matcher
TEST_DB_HOST=localhost
TEST_DB_NAME=matchertest
```

## To Run

`docker-compose up` starts postgres db and golang application. Golang api is listening on port `3001`.

## Overview

This is a simple api that exposes 4 endpoints:

- `GET    /v1/users/:id` Gets a user by user id
- `POST   /v1/users` Creates a user
- `GET    /v1/likes/:userId` Get all the users that like the :userId user
- `POST   /v1/likes` Create a new like relationship

Uses docker to run both a postgres DB and the golang application. On `docker-compose up` will run make docker-fresh which first cleans the DB (deletes the data) and then seeds it with new test data. 

Using a VS Code extension `https://github.com/Huachao/vscode-restclient` this allows you to run CURL commands within vscode from the http.rest file.

## Structure

Loosely modeled after Clean Architecture https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

Main entry point is in the `api/main.go`. Usecase folder contains the application services. The services is where the major business logic is each is based off of an interface which decoupels the service implementation and any code thats using it.

The api folder is all the files associated with communication of REST in this case. It contains the handlers specific for each different route. Input a collection of structs used for validating the data that is entering the system. Output used to create a contract for the data sent/returned.

Using the interface also makes it much easier to write unit tests that only are testing the service and are not dependent on other services.

Entity folder contains application wide `entities/structs`. These are mostly the db models

Repository folder contains the data access layer structs, each are responsible for one folder and the services can only access the data through the corresponding repository.

Seeders is a seperate application where there is another main.go file. This is used to seed the database with test data currently only seeding test users.

db folder is where the database connection is created and any db scripts that are nesassary for creating the database.

Config folder is used for any application configuration that is required at startup. Generally these config variables are from environment files.

## Commands

Using Makefle
- `make run` to run the application locally.
- `make seed-up` to seed the DB with test data
- `make seed-down` to remove data within the db
- `make test-unit` to run all unit tests
- `make test-integration` to run integration tests

## Endpoints

Create User

```
POST http://localhost:3001/v1/users
Content-Type: application/json

{
    "firstName": "Toby",
    "lastName": "Test",
    "email": "toby.test@test.com",
    "pass": "newPassword"
}
```
Get likes, gets who likes the user. Takes query params for pagination the default values are:
page=0
pageSize=100

```
GET http://localhost:3001/v1/likes/fb8fb77f-295d-4126-bef9-9aea2b93ae43?page=10&pageSize=3
```

Create like, fromId and toId need to be uuids and within the DB
```
POST http://localhost:3001/v1/likes
Content-Type: application/json
{
    "fromId": "05563251-45af-4bb9-b5bf-21cae961b254",
    "toId": "fb8fb77f-295d-4126-bef9-9aea2b93ae43"
}
```
## Notes

There is no authorization or authentication due to its complexity, I would use a JWT based auth solution and use gin's middleware to interagat the token for authentication and authorization (the user id within + scopes).