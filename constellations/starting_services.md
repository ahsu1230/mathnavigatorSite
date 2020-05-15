# Starting services with Docker

The MathNavigator repository follows a microservice architecture. This means that this application is made up of many "services".

## Docker commands

```unix
docker-compose up -d
docker-compose start SERVICE
docker-compose stop SERVICE
docker-compose build SERVICE
docker-compose restart SERVICE

docker-compose kill SERVICE
docker-compose rm

docker-compose exec SERVICE bash
```

## Services and their ports

You can find the following services with their corresponding host/port locations. For instance, to access the `gemini-user` website simply open `http://localhost:8000` with an Internet browser. For webservers like `orion`, nothing will show up if you go to `localhost:6001`, but you can cUrl endpoints at `http://localhost:6001/api/...`.

- orion: `localhost:6001`
- db-mysql: `localhost:3308`
- gemini-user: `localhost:8000`