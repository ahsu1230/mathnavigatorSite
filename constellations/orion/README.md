# Orion

Our core API service. This webserver will provide the main API to allow web clients to interact with our MySQL database. This is where users, programs, classes, announcements are created and persisted into the database.

## Pre-requirements

Make sure that Go is installed before proceeding: (https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/onboarding/install_go.md)

## How to run the Orion webserver

```unix
docker-compose build orion
docker-compose start orion
```

View your DockerDesktop to check if the `orion` service is running and healthy.

## Running Tests

To run all tests for the back-end web server, run:
```
go test ./...
```
You should see `ok`s and no failures.

## Navigating the codebase

There are 3 architecture "layers" to Orion.

- Domains - simple objects that represent an entity (i.e. a Program, a Class, a User, etc.)
  - Domains are very simple and don't have much logic to them.
  - All domains do are describe attributes of an entity.

- Controllers - objects that handle network stuff (JSON, HTTP Request & Responses, Serializing & Deserializing, etc.)
  - The controller layer is built on top of `gin` - a golang http framework.

- Repos - objects that handle database stuff (MySQL connections, Database queries, reads & writes, etc.)
  - For testing, database queries are checked via Datadog's SQL query validation.

Both Controllers and Repos are often declared as interfaces. The reasoning is to promote mock unit testing. Essentially, there are usually two implementations of every Controller and Repo. One implementation is for the business logic (works as you would expect), the other implementation is done by test classes which help simply unit testing.

Finally you can find all integration tests inside the folder `tests_integration`.

## Formatting your code

```
go fmt ./...
```