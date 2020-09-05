# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13

ARG app_env
ENV APP_ENV $app_env
ENV DB_HOST "db-mysql"
ENV DB_PORT "3306"
ENV DB_USER "user"
ENV DB_PASSWORD "password"
ENV DB_DEFAULT "mathnavdb"
ENV CORS_ORIGIN "*"
ENV REDIS_HOST "cache-redis"
ENV REDIS_PORT "6379"
ENV REDIS_PASSWORD "redis_password"

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/ahsu1230/mathnavigatorSite/constellations/orion
WORKDIR /go/src/github.com/ahsu1230/mathnavigatorSite/constellations/orion

# Install project dependencies
RUN go get github.com/ahsu1230/mathnavigatorSite/constellations/orion

# Build the command inside the container by installing the binary.
RUN go install .

## Install and launch wait tool and then launch application
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

# Run the command by default when the container starts.
ENTRYPOINT /wait && /go/bin/orion

# Document that the container service listens on port 8080.
EXPOSE 8001