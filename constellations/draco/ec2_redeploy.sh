#!/bin/bash

# Rebuild orion:latest from DockerHub
docker pull ahsu1230/mathnavigator-orion:latest
docker stop mathnavigator-orion
docker rm mathnavigator-orion
# Must change environment variables below!
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

# Rebuild gemini-user:latest from DockerHub
docker pull ahsu1230/mathnavigator-gemini-user:latest
docker stop mathnavigator-gemini-user
docker rm mathnavigator-gemini-user
docker run --name mathnavigator-gemini-user -p 80:80 -itd ahsu1230/mathnavigator-gemini-user
