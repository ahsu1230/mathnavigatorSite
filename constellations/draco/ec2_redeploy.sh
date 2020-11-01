#!/bin/bash

# Rebuild :latest from DockerHub
docker pull ahsu1230/mathnavigator-orion:latest
docker pull ahsu1230/mathnavigator-gemini-user:latest

# Stop current containers
docker stop mathnavigator-orion; docker rm mathnavigator-orion
docker stop mathnavigator-gemini-user; docker rm mathnavigator-gemini-user

# Rerun Docker commands
docker run --name mathnavigator-orion -itd -p 8001:8001 \
    -e APP_ENV=production \
    -e DB_HOST="HOST" \
    -e DB_PORT="PORT" \
    -e DB_USER="USER" \
    -e DB_PASSWORD="PASSWORD" \
    -e DB_DEFAULT=mathnavdb \
    -e REDIS_HOST=cache-redis \
    -e REDIS_PORT=6379 \
    -e REDIS_PASSWORD=redis_password \
    -e CORS_ORIGIN="*" \
    ahsu1230/mathnavigator-orion
docker run --name mathnavigator-gemini-user -p 80:80 -itd ahsu1230/mathnavigator-gemini-user
docker ps
