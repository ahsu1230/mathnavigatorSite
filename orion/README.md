# Orion
The most famous constellation that can be seen all around the world. This will be our core API service. The user and admin websites will both funnel through this service in order for users to interact with our database.

## Onboarding Steps:
By the end of these steps, you should be able to run a local webserver to both host the back-end golang servers and a front-end web application.


## Install NodeJs and NPM
Install [NodeJs](https://nodejs.org/en/download/).
Please follow instructions to correctly install.
Once finished, in your Terminal / DOS, run:
```
node -v
npm -v
```
These commands should respond with the versioning of your node and npm without errors.
When you are finished, you should be able to run the following commands in Terminal:

## Install Go
The language of choice we're using is Google's programming language [Golang](https://golang.org/).
Please follow the instructions to download and install Go onto your machine.
Once finished, in your Terminal / DOS, run:
```
go version
```
This will print out the OS and library version of your Golang.

## Install MySQL
Download MySQL from [here](https://dev.mysql.com/downloads/mysql/). Select your operating system and download the first MySQL app. For MacOS, you should be selecting the DMG Archive. For Windows, download the ZIP Archive. From there, follow the instructions to completely install MySQL.

**The installer may ask you for a MySQL password. Please remember this password!**

Once finished installing, test it here:
```
mysql --user=root --password
```
* If you get a command not found, do the following below. Otherwise, skip this section.
------
### If MySQL command is not found...
Looks like we'll have to add `mysql` to our environment variables.

#### For Mac users...
Edit this file by typing
```
vi .bash_profile
```
You are now using the default text editor called `Vim`.
Press `I` to insert content.
At the end of the file, add this line:
```
export PATH="/usr/local/mysql/bin:$PATH"
```

Double check to see if that line is correct.
Then save this file by press the following keys: `Esc` -> `:wq` (including the colon).

If there are no errors, the file saved correctly!
Run this command:
```
source .bash_profile
```

And we can try this command one more time to see if mysql now works.
```
mysql --user=root --password
```

#### For Windows users... (coming soon!)
...

------
### If MySQL command is found...
Run this command:
```
mysql --user=root --password
```
The password is either nothing or the password that you entered before. If this worked, you should see a Welcome message and at the start of your command line, you should see:
```
mysql>
```

From here, type in:
```
Create SCHEMA mathnavdb;
```
If success (Query OK), you can exit MySql by typing `exit`.

Once you exit out of MySQL, remember these three commands. They will start or stop your local MySql server on your machine.
```
mysql.server start
mysql.server stop
mysql.server restart
```
Alternatively, if that doesn't work, you can use the MySQL Notifier app that comes with installation and appears in the taskbar:

![alt text](https://github.com/ahsu1230/mathnavigatorSite/blob/max/orion/Untitled.png)
------

## Test back-end webserver
In your Terminal, go to the `orion` directory. Use `cd` to travers around your file system.

To run all tests of the back-end web server, run:
```
go test ./...
```
You should see `ok`s and no failures.

In the `orion` folder,
 * Create a new folder called `configs`.
 * Inside this folder, create a new file called `config_local.yaml`.
 * Paste the following content into this file and save.
```
app:
  build: "development"
  corsOrigin: "http://localhost:8081"
database:
  host: "localhost"
  port: 3306
  user: "root"
  pass: "<YOUR_PASSWORD_GOES_HERE>"
```
Remember the password you saved for MySql? Paste that password in between the quotations!

After that, start the web server with this:
```
go run main.go configs/config_local.yaml
```
You should see a `Listening and serving HTTP on :8080` message. It worked! Now, you are currently running the Math Navigator webserver locally on your machine. Any HTTP requests to the port number 8080 will be received and responded to by the local webserver!

## Test front-end web application (Admin)
From here, open a new Terminal tab/window.
Use `cd` to travers to `sites/admin`.

Run these commands:
```
npm install
npm run start
```
This will run another webserver (this time at port 8081) which will host this website.
Now, if you go to http://localhost:8081 in an Internet browser like Chrome, you should be able to see the Admin website!

For future reference, `npm run start` creates a development version of the website application.
If you need to create a "production" version (in other words, a version for consumers), run this command:
```
npm run build
```
This will create the production version and will NOT start the webserver.


## Run the entire environment
Great job getting here! You've built 2 local webservers to host the web application and the backend API server.

To make it easier on yourself in the future and to develop faster, you can use this script here in the `orion` folder.
Before doing so, close the other Terminal tabs so we don't end up creating multiple web servers with conflicting port numbers.
*Mac users only*
```
./run_admin_local.sh
```
*Windows Users only*
```
coming soon...
```

Use Control+C to stop all webservers and run that command to spin it up again!

## MySql GUI
To view your MySql database, you can either use the Terminal or download a MySQL GUI. The most popular free GUI is [MySql Workbench](https://dev.mysql.com/downloads/workbench/).

If you decide to work in Terminal, use this command to sign in:
```
mysql -u root -p
```

If you want to use MySql Workbench, create a New Connection with the following properties:
 - Connection Method: Standard TCP/IP
 - Hostname:  `127.0.01`
 - Port: `3306`
 - Username: `root`
 - Password: `YOUR_PASSWORD_FROM_MYSQL_SECTION`
