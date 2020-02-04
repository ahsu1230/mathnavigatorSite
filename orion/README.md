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
The following instructions differ between MacOS and Windows developers. Please skip to the appropriate section.

### Installing MySQL (MacOS)
Download MySQL from [here](https://dev.mysql.com/downloads/mysql/). Select Operating System macOS and download the first macOS DMG Archive.
**The installer may prompt you for a MySQL password. Please remember this password!**

After installation, we'll need to set our environment variables so your computer can use `mysql` from the Terminal.
Open up Terminal and edit this file by typing
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

And we can try this command to see if mysql works.
```
mysql --version
```
If no errors, our environment variables are setup correctly! We can start working with mysql.

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

![mysql_notifier](onboarding/mysql_notifier.png)
------
### Installing MySQL (Windows)
Download MySQL from [here](https://dev.mysql.com/downloads/windows/installer/). Download the latest and smaller MSI Installer.
*NOTE* the instructions in this article may be outdated as they use an older version of Windows and MySQL. However, the installation process should still be similar.
 - Once finished downloading the installer, double click it to start the installation process.
 - `Choosing a Setup Type` Select **Custom**
 - `Select Products and Features` Select the following: latest MySQL Server, latest MySQL Workbench (under Applications), MySQL Shell (also under Applications). Press Next.
 - `Installation` Execute! You should be downloading 3 products (Server, Workbench, Shell). Once finished Downloading and Installing, press Next.
 - You'll need to configure MySQL Server.
   - `High Availability` Select Standalone MySQL Server.
   - `Type and Networking` Make sure the Config Type is `Development Computer` and make sure you have TCP/IP checked, and Port: 3306.

<img src="onboarding/mysql_1_installer_networking.png" width="480">

   - `Authentication Method` Select the Recommended Strong Password Encryption.
   - `Account and Roles` Create a MySQL password. **Please remember this password!**

<img src="onboarding/mysql_2_installer_accounts.png" width="480">

   - `Windows Service` Under Windows Service Name, remove the numbers and just have the name be: `MYSQL`. In addtion, you can uncheck the `Start the MySQL Server at System Startup`. Tou can also run the Windows Service as the Standard system.
   - `Apply Configuration` Execute! Finish.

Now that you've installed MySQL, remember these two commands:
```
net start MySQL
net stop MySQL
```
There two commands will start or stop your MySQL local server. If the local server is not started, your MySQL will not work.

You can go ahead and open the application MySQL Shell. It will prompt you to enter your MySQL password.
<img src="onboarding/mysql_8_shell.png" width="480">

Once you're in, run this command:
```
CREATE SCHEMA mathnavdb;
exit;
```
Congratulations! MySQL is successfully installed.

**BONUS (optional)** You don't need to do this since you've installed the MySQL Shell. But if you would like to work with MySQL from the Command Prompt, you may edit your Environment Variables and include the MySQL path into the Environment Variable `PATH`.
 - Look for the folder: `C:\Program Files\MySQL\MySQL Server\bin`. Inside this directory, there should be a `mysql.exe`. If it is there, copy the Location to this folder (NOT the .exe). An example Location could be: `C:\Program Files\MySQL\MySql Server 8.0\bin`
 - From here, go to your computer's `Control Panel` > `System and Security` > `System` > `Advanced system settings` > `Environment Variables`.
   - Edit the environment variable `PATH`.
   - Add a semicolon `;` at the end of the value. 
   - Paste the copied MySQL Location folder
   - Save by pressing OK
   
<img src="onboarding/mysql_6_env_var.png" width="480">

   - Open Command Prompt and type `mysql --version` to see if mysql is recognized by the Command Prompt.

------

## Test back-end webserver
Before proceeding, double check to make sure your mysql server is running. For MacOs, it is the `mysql.server start` command and for Windows, it is the `net start MySQL`. 
In your Terminal, go to the `orion` directory. Use `cd` to travers around your file system.

To run all tests of the back-end web server, run:
```
go test ./...
```
You should see `ok`s and no failures.
**note (Windows):** The first time you run this command, you may need to run Command Prompt as administrator. Simply close Command, right click the Application and select "Run as administrator"

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

After that, go back to the `orion` directory and start the web server with this:
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
