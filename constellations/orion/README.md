# Orion

Our core API service. This webserver will provide the main API to allow web clients to interact with our MySQL database. This is where users, programs, classes, announcements are created and persisted into the database.

## How to run the Orion webserver

```unix
docker-compose build orion
docker-compose start orion
```

View your DockerDesktop to check if the `orion` service is running and healthy.

---

## Running Tests

To run all tests for the back-end web server, run:
```
go test ./...
```
You should see `ok`s and no failures.

---

## Navigating the codebase

Coming Soon...