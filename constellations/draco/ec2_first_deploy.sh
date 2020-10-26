#!/bin/bash

# Update EC2 instance
echo "*** Updating EC2..." 
sudo yum update -y

# Install Git
echo "*** Installing Git..." 
sudo yum install git -y
git version

# Install Docker & Docker-Compose
echo "*** Installing Docker..." 
sudo amazon-linux-extras install docker -y
sudo yum install docker -y
sudo service docker start
sudo usermod -a -G docker ec2-user
sudo docker info
docker version

sudo curl -L https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose version

# Install Golang
echo "*** Installing Golang..." 
sudo yum install golang -y
go version
# Setup: GOROOT, GOPATH and PATH (in .bash_profile) (?)

# Install NVM for npm
# https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/setting-up-node-on-ec2-instance.html
echo "*** Installing NVM..." 
cd $HOME
mkdir .nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash
# . ~/.nvm/nvm.sh
# export NVM_DIR="$HOME/.nvm"
# [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"

echo 'export NVM_DIR="$HOME/.nvm"' >> $HOME/.bashrc
echo '[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"' >> $HOME/.bashrc
 
# Dot source the files to ensure that variables are available within the current shell
. $HOME/.bashrc
. $HOME/.nvm/nvm.sh

echo "* Nvm version"
nvm --version
nvm install node

echo "* npm & node versions"
npm --version
node --version

# cd to repository
echo "*** Cloning Repository..." 
cd $HOME
git clone https://github.com/ahsu1230/mathnavigatorSite.git
wait 
echo "* Finished cloning!"
cd $HOME/mathnavigatorSite
git fetch
git checkout aws_deployment # <--- CHANGE LATER
git pull

# Rebuild user site
echo "*** Building gemini-user site..." 
cd $HOME/mathnavigatorSite/constellations/gemini-user
npm install
npm run build

# Start services with Docker-Compose
# Rebuild the orion container & gemini-user express HTTP server container
echo "*** Running Docker-Compose..." 
cd $HOME/mathnavigatorSite/constellations
docker-compose -f docker-compose.production.yml up -d --build
