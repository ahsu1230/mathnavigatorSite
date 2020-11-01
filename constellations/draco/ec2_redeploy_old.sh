#!/bin/bash

# cd to repository
cd /home/ec2-user/mathnavigatorSite
git fetch
git pull

# Rebuild user site
cd /home/ec2-user/mathnavigatorSite/constellations/gemini-user
rm -rf dist
npm run build

# Start services with Docker-Compose
# Rebuild the orion container & gemini-user express HTTP server container
cd /home/ec2-user/mathnavigatorSite/constellations
docker-compose -f docker-compose.production.yml build orion
docker-compose -f docker-compose.production.yml build gemini-user-prod
docker-compose -f docker-compose.production.yml up -d
# docker-compose -f docker-compose.production.yml up --no-deps -d orion