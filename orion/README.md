# Orion
The most famous constellation that can be seen all around the world. This will be our core API service. The user and admin websites will both funnel through this service in order for users to interact with our database.

# Installing Go
The language of choice we're using is Google's programming language [Golang](https://golang.org/).
Please follow the instructions to download and install Go onto your machine.
Once finished, in your Terminal / DOS, run:
```
go version
```
This will print out the OS and library version of your Golang.

# Installing NPM
Look [here]((https://nodejs.org/en/download/)) to download the correct NodeJs and NPM sources onto your computer. Please follow instructions to correctly install.

Once finished, in your Terminal / DOS, run:
```
node -v
npm -v
```
These commands should respond with the versioning of your node and npm without errors.

# Installing and starting MySQL
Download MySQL from [here](https://dev.mysql.com/downloads/mysql/). Select your operating system and download the first MySQL app. For MacOS, you should be selecting the DMG Archive. For Windows, download the ZIP Archive. From there, follow the instructions to completely install MySQL.

Once finished installing, in Terminal / DOS, run: 
```
mysql --user=root --password
```
The password is nothing (just press Enter). If this worked, you should see a Welcome message and at the start of your command line, you should see:
```
mysql>
```

From here, type in:
```
Create SCHEMA mathnavdb;
```
If success, you can exit MySql by typing `exit`.

Once you exit out of MySQL, remember these three commands. They will start or stop your local MySql server on your machine.
```
mysql.server start
mysql.server stop
mysql.server restart
```

# Run the environment
We can create the whole admin server/website environment by running this script.
```
./run_admin_local.sh
```
In your browser, go to http://localhost:8080 to interact with the website.

# Developing Server-side only
When developing only on the server component of the web application, you can run the following commands from this directory.

 - We can run the server just by itself:
```
go run main.go configs/config_local.yaml
```
This will run the server on http://localhost:8080.
From here, you can use cUrl requests to test HTTP API interactions.

 - We can run tests for the server:
```
go test ./tests/*
```

# Developing Client-side only
When developing on the website, go to a website folder under `./sites/*`. From there, you can follow more specific directions to run the website. But, in general, it will usually follow this pattern.

This following script is to build the site locally and to test.
```
npm run start
```
This following script is to create a production build that can then be placed onto a live host server.
```
npm run build
```
