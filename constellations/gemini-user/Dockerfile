# Docker Image which is used as foundation to create
# a custom Docker Image with this Dockerfile
FROM node:14-alpine

# Copies everything over to Docker environment
# Copy the local package files to the container's workspace.
ADD . /app
WORKDIR /app

# Reinstall website
RUN npm install

# Build production website
RUN npm run build

## Install and launch wait tool and then launch application
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

# Run the command by default when the container starts.
ENTRYPOINT /wait && node productionServer.js

# Uses port which is used by the actual application
EXPOSE 80