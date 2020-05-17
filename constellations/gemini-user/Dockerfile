# Docker Image which is used as foundation to create
# a custom Docker Image with this Dockerfile
FROM node:14-alpine
 
# A directory within the virtualized Docker environment
# Becomes more relevant when using Docker Compose later
WORKDIR /usr/src/app
ARG app_env
ENV APP_ENV $app_env
ENV ORION_HOST="orion:8080"
 
# Copies package.json and package-lock.json to Docker environment
COPY package*.json ./
 
# Installs all node packages
RUN npm install
 
# Copies everything over to Docker environment
COPY . .
 
# Finally runs the application
CMD if [ ${NODE_ENV} = production ]; \
	then \
	npm install -g http-server && \
	npm run build && \
	cd build && \
	hs -p 9000; \
	else \
	npm run start; \
	fi

# Uses port which is used by the actual application
EXPOSE 9000