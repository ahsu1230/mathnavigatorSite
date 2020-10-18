#!/bin/bash

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
docker-compose -f docker-compose.production.yml up -d