# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.15

ARG app_env
ARG db_host
ARG db_port
ARG db_user
ARG db_password
ARG redis_host
ARG redis_port
ARG redis_password

ENV APP_ENV $app_env
ENV DB_HOST $db_host
ENV DB_PORT $db_port
ENV DB_USER $db_user
ENV DB_PASSWORD $db_password
ENV DB_DEFAULT "mathnavdb"
ENV CORS_ORIGIN "*"
ENV REDIS_HOST $redis_host
ENV REDIS_PORT $redis_port
ENV REDIS_PASSWORD $redis_password

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
