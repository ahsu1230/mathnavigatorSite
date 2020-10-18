#!/bin/bash

#chmod 777 ec2_orion_init.sh
#chmod 777 ec2_orion_update.sh

# Update EC2 instance
sudo yum update -y

# Install Git
sudo yum install git -y
git version
git clone https://github.com/ahsu1230/mathnavigatorSite.git

# Install Golang
sudo yum install golang -y
go version
# Setup: GOROOT, GOPATH and PATH (in .bash_profile) (?)

# Install Docker & Docker-Compose
sudo amazon-linux-extras install docker -y
sudo yum install docker -y
sudo service docker start
sudo usermod -a -G docker ec2-user
sudo docker info
docker version

sudo curl -L https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose version

# Start services with Docker
# cd to orion folder
# DO I HAVE TO USE REAL ADMIN AND PASSWORD????
#-        db_user: user
#-        db_password: password
#+        db_user: admin
#+        db_password: "___$$____"
docker-compose -f docker-compose.production.yml up -d

# Install NVM for npm
# cd to gemini-admin
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.32.0/install.sh | bash
. ~/.nvm/nvm.sh
nvm install stable
npm install --verbose
npm run build # create production build
sudo cp index.html /var/www/html
sudo cp -r dist /var/www/html

# Install apache webserver?
sudo yum install httpd -y
sudo service httpd start
sudo chkconfig httpd on

# copy index.html to /var/www/html OR create index.html @ /var/www/html