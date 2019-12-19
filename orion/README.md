# Orion
The most famous constellation that can be seen all around the world. This will be our core API service. The user and admin websites will both funnel through this service in order for users to interact with our database.

# Installing Go
The language of choice we're using is Google's programming language Golang (https://golang.org/).
Please follow the instructions to download and install Go onto your machine.
Once finished, in your Terminal / DOS, run:
```
go version
```
This will print out the OS and library version of your Golang.

# Installing NPM
Look here to download the correct NodeJs and NPM sources onto your computer (https://nodejs.org/en/download/). Please follow instructions to correctly install.
Once finished, in your Terminal / DOS, run:
```
node -v
npm -v
```
These commands should respond with the versioning of your node and npm without errors.

# Developing Server-side
There are 3 ways to run the server for development.

 - We can run the server just by itself:
```
go run main.go configs/config_local.yaml
```
This will run the server on localhost:8080.
From here, you can use cUrl requests to test HTTP API interactions.

 - We can run tests for the server:
```
go test ./tests/*
```

 - Finally, we can create the whole server / admin website environment by running this script.
```
./run_admin_local.sh
```
In your browser, go to http://localhost:8080 to interact with the website.

# Developing Client-side
When developing on the website, go to a website folder under `./sites/*`
This script is to build the site locally and to test.
```
npm run start
```
This script is to create a production build that can then be placed onto a live host server.
```
npm run build
```

To have the website fully interact with a local web server, you may run this command in one terminal / command prompt.
```
go run main.go configs/config_local.yaml
```
Then run this command in another terminal / command prompt.
```
npm run start
```
