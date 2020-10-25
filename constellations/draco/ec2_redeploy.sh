#!/bin/bash

chmod 777 ec2_redeploy.sh

# cd to repository
cd /home/ec2-user/mathnavigatorSite
git fetch
git checkout aws_deployment # <--- CHANGE LATER
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

# Do I still need this???
# copy index.html to /var/www/html OR create index.html @ /var/www/html
# cd to gemini-admin
# sudo cp index.html /var/www/html
# sed -i 's/dist/dist\//g' dist/bundle.js # need this to fix images in admin :(
# sudo cp -r dist /var/www/html
# sudo cp assets/*.svg /var/www/html/dist
