# Alternative Onboarding Guide

Use this guide if Docker-Desktop is not supported on your computer. The difference in this guide is the backend service setup. You will have to install MySql, get your local Mysql server running, and get the Orion webserver running.

## Install MySQL

To install MySQL, refer to this [guide](../../resources/onboarding/install_mysql.md).

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
export CORS_ORIGIN=*
```

**For Windows users**

Read this guide here <https://www.architectryan.com/2018/08/31/how-to-change-environment-variables-on-windows-10/> to add environment variables. You are NOT editing the environment variable `PATH`. You simply need to create the following 5 environment variables.

```unix
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=<YOUR_MYSQL_PASSWORD>
CORS_ORIGIN=*
```

## Running Orion

Make sure Go is installed on your computer, which you can find information for [here](../../resources/onboarding/install_go.md).

In the orion folder, run the following command:

```unix
go run main.go
```

If things are running correctly, you should see a `Listening and serving HTTP on :6001` message. The webserver will continue to run as a process as long as the CLI tab is running. To stop it, use Ctrl+C to cancel the process and the webserver will stop running. If you want it to continue running, simply create a new CLI window/tab to work on other things while the web server is running.

## Continue the Guide

Great! Now that your webserver is running, we can go ahead and get the web clients to work as well. Please go back to the guide and start the Gemini Admin and Gemini User sites. [Start the Web Clients](./guide.md#starting-gemini-admin).