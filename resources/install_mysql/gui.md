# MySQL GUI
To view your MySQL database, you can either use the Terminal/Command Prompt or download a MySQL GUI. The most popular free GUI is [MySQL Workbench](https://dev.mysql.com/downloads/workbench/).

If you decide to work in Terminal or MySQL Shell, use this command to sign in:
```
mysql -u root -p
```

If you want to use MySQL Workbench, create a New Connection with the following properties:
 - Connection Method: Standard TCP/IP
 - Hostname:  `127.0.0.1`
 - Port: `3306`
 - Username: `root`
 - Password: `YOUR_PASSWORD_FROM_MYSQL_SECTION`