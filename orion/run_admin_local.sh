# This script builds the website locally AND runs the orion server locally
# So you may interact with the website and api server end to end.

# First step is to make sure your computer mysql.server is running
# mysql.server start
# mysql.server stop
# mysql.server restart
# mysql db --user=root --password 

# Second step is to build the website locally
cd sites/admin
npm run build-local
cd ../..

# Third step is to run the api server locally
go run main.go configs/config_local.yaml
