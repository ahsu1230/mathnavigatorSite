# Install MySQL
The following instructions differ between MacOS and Windows developers. Please skip to the appropriate section.

### Installing MySQL (MacOS)
Download MySQL from [here](https://dev.mysql.com/downloads/mysql/). Select Operating System macOS and download the first macOS DMG Archive.
**The installer may prompt you for a MySQL password. Please remember this password!**

After installation, we'll need to set our environment variables so your computer can use `mysql` from the Terminal.
Open up Terminal and edit this file by typing
```
vim .bash_profile
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

And we can try this command to see if `mysql` works.
```
mysql --version
```
If there are no errors, our environment variables are setup correctly! We can start working with MySQL.

Run this command:
```
mysql --user=root --password
```
This will prompt you to enter your password. The password is either nothing or the password that you entered before. After you log in, you should see a Welcome message and at the start of your command line, you should see:
```
mysql>
```

From here, type in:
```
CREATE DATABASE mathnavdb;
USE DATABASE mathnavdb;
```
If success (Query OK), you can exit MySql by typing `exit`.

Once you exit out of MySQL, remember these three commands. They will start, stop, and restart your local MySQL server on your machine, respectively.
```
mysql.server start
mysql.server stop
mysql.server restart
```
Alternatively, if that doesn't work, you can use the MySQL Notifier app that comes with installation and appears in the taskbar:

![MySQL Notifier](onboarding/mysql_notifier.png "MySQL Notifier")
------
### Installing MySQL (Windows)
Download MySQL from [here](https://dev.mysql.com/downloads/windows/installer/). Download the latest and smaller MSI Installer.
**Note:** the instructions in this article may be outdated as they use an older version of Windows and MySQL. However, the installation process should still be similar.
 - Once finished downloading the installer, double click it to start the installation process.
 - `Choosing a Setup Type` Select **Custom**
 - `Select Products and Features` Select the following: latest MySQL Server, latest MySQL Workbench (under Applications), MySQL Shell (also under Applications). Press Next.
 - `Installation` Execute! You should be downloading 3 products (Server, Workbench, Shell). Once finished downloading and installing, press Next.
 - You'll need to configure MySQL Server.
   - `High Availability` Select Standalone MySQL Server.
   - `Type and Networking` Make sure the Config Type is `Development Computer` and make sure you have TCP/IP checked, and Port: `3306`.

<img src="images/mysql_1_installer_networking.png" width="480" alt="Type and Networking">

   - `Authentication Method` Select the Recommended Strong Password Encryption.
   - `Account and Roles` Create a MySQL password. **Please remember this password!**

<img src="images/mysql_2_installer_accounts.png" width="480" alt="Accounts and Roles">

   - `Windows Service` Under Windows Service Name, remove the numbers and just have the name be: `MYSQL`. In addition, you can uncheck the `Start the MySQL Server at System Startup`. You can also run the Windows Service as the Standard system.
   - `Apply Configuration` Execute! Finish.

Now that you've installed MySQL, remember these two commands:
```
net start MySQL
net stop MySQL
```
There two commands will start or stop your MySQL local server. If the local server is not started, your MySQL will not work. You can also use the MySQL Notifier app in the taskbar.

You can go ahead and open the application MySQL Command Line Client. It will prompt you to enter your MySQL password.

<img src="images/mysql_8_shell.png" width="480" alt="MySQL Shell">

Once you're in, run this command:
```
CREATE DATABASE mathnavdb;
USE DATABASE mathnavdb;
exit;
```
If success (Query OK), you can exit MySql by typing `exit`.
Congratulations! MySQL is successfully installed.

**BONUS (optional)** You don't need to do this since you've installed the MySQL Shell. But if you would like to work with MySQL from the Command Prompt, you may edit your Environment Variables and include the MySQL path into the Environment Variable `PATH`.
 - Look for the folder: `C:\Program Files\MySQL\MySQL Server\bin`. Inside this directory, there should be a `mysql.exe`. If it is there, copy the Location to this folder (NOT the .exe). An example Location could be: `C:\Program Files\MySQL\MySql Server 8.0\bin`
 - From here, go to your computer's `Control Panel` > `System and Security` > `System` > `Advanced system settings` > `Environment Variables`.
   - Edit the environment variable `PATH`.
   - Add a semicolon `;` at the end of the value.
   - Paste the copied MySQL Location folder
   - Save by pressing OK

<img src="images/mysql_6_env_var.png" width="480" alt="Environment Variables">

   - Open Command Prompt and type `mysql --version` to see if `mysql` is recognized by the Command Prompt.


## MySQL GUI
To view your MySQL database, you can either use the Terminal or download a MySQL GUI. The most popular free GUI is [MySQL Workbench](https://dev.mysql.com/downloads/workbench/).

If you decide to work in Terminal, use this command to sign in:
```
mysql -u root -p
```

If you want to use MySQL Workbench, create a New Connection with the following properties:
 - Connection Method: Standard TCP/IP
 - Hostname:  `127.0.01`
 - Port: `3306`
 - Username: `root`
 - Password: `YOUR_PASSWORD_FROM_MYSQL_SECTION`