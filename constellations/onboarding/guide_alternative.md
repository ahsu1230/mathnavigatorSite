# Alternative Onboarding Guide

Use this guide if Docker-Desktop is not supported on your computer. The difference in this guide is the backend service setup. You will have to install MySql, get your local Mysql server running, and get the Orion webserver running.

## Install Golang

Make sure that [Golang](https://golang.org/) is installed before proceeding. Please follow the instructions to download and install Go onto your machine.

Once finished, in your Terminal / DOS, run:
```
go version
```
This will print out the OS and library version of your Golang.

## Install MySQL

To install MySQL, refer to this [guide](../../resources/install_mysql/guide.md).

## Setting up your environment for Orion Integration Testing

Open MySQL with your CLI. You can use Terminal for Mac or use MySQL Shell for Windows. Once you're signed into the MySQL (via `mysql --user=root --password`), we'll need to create another user for integration testing.

Enter the following commands into the MySQL Shell.

```
CREATE USER 'ci_tester'@'localhost' IDENTIFIED BY 'test';
GRANT ALL PRIVILEGES ON * . * TO 'ci_tester'@'localhost';
FLUSH PRIVILEGES;
CREATE DATABASE mathnavdb_test;
```

When you are finished, exit MySQL. And inside the `orion` folder, run this command to test the orion webserver is working properly.

```unix
go test -v ./...
```

## Setup Environmental Variables

Environmental variables are global constants that your computer may use in various programs. We're going to add a few to get orion to work.

**For MacOS users**

You can add environment variables using these commands:

```unix
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=<YOUR_MYSQL_PASSWORD>
export DB_DEFAULT=mathnavdb
export CORS_ORIGIN=*
```

**For Windows users**

Read this guide here <https://www.architectryan.com/2018/08/31/how-to-change-environment-variables-on-windows-10/> to add environment variables. You are NOT editing the environment variable `PATH`. You simply need to create the following 5 user environment variables.

```unix
Key: DB_HOST
Value: localhost

Key: DB_PORT
Value: 3306

Key: DB_USER
Value: root

Key: DB_PASSWORD
Value: <YOUR_MYSQL_PASSWORD>

Key: DB_DEFAULT
Value: mathnavdb

Key: CORS_ORIGIN
Value: *
```

## Running Orion

Now that environment variables are set, run the following command in the `orion` folder.

```unix
go run main.go
```

If things are running correctly, you should see a `Listening and serving HTTP on :8001` message. The webserver will continue to run as a process as long as the CLI tab is running. To stop it, use `Ctrl+C` to cancel the process and the webserver will stop running. If you want it to continue running, simply create a new CLI window/tab to work on other things while the web server is running.

## Continue the Guide

Great! Now that your webserver is running, we can go ahead and get the web clients to work as well. Please go back to the guide and start the Gemini Admin and Gemini User sites. [Start the Web Clients](./guide.md#starting-gemini-admin).