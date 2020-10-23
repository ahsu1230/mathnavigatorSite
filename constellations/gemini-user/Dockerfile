# Docker Image which is used as foundation to create
# a custom Docker Image with this Dockerfile
FROM node:14-alpine
 
# A directory within the virtualized Docker environment
# Becomes more relevant when using Docker Compose later
# WORKDIR /usr/src/app
 
# Copies everything over to Docker environment
COPY . /

RUN rm -rf dist
RUN npm install
RUN npm run build
RUN ls -la
RUN ls -la dist

## Install and launch wait tool and then launch application
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

# Run the command by default when the container starts.
ENTRYPOINT /wait && node productionServer.js

# Uses port which is used by the actual application
EXPOSE 80