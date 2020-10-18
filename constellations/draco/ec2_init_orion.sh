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
