#!/bin/bash

docker-compose -f docker-compose.production.yml build orion
docker-compose -f docker-compose.production.yml up --no-deps -d orion
