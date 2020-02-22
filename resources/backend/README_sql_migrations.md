# What is a SQL migration?

When software (in our case, the webserver) uses a database (like MySQL), we can run the risk of the software and database being "out of sync".
Let's take an example:

Suppose Aaron and Joe are working together. In each of their local environments (code on their computers), they have the following:
 - Webserver version 1.0
 - Database 1.0 with table A

 Let's say Aaron writes a new feature for the webserver (now version 2.0) and pushes it onto Github. In this feature, he also introduces new tables into the database (version 2.0 now with tables A, B, C).

 When Joe pulls the latest code from Github later, he will now have webserver version 2.0. However, his local database is still on version 1.0! *Databases do not automatically update!!!*
 Now, his webserver won't work because the latest features want to interact with tables B & C which Joe has not created yet.

He could just manually create them, but that would be such a pain in the butt. So instead, migrations help with this problem. You can setup SQL migrations to automatically upgrade your database version if you don't have the latest stuff. Similarly, if there are problems in the latest version, it's a good idea to downgrade to a version that works!

So with Automatic SQL migrations, Joe can pull the latest webserver and automatically update his MySQL databases to match those of Aaron's. SO COOL!!

Sometimes, SQL Migrations also refer to transferring from one database to another. For instance, let's say one day MySQL doesn't work for us and instead, we would like to switch to another database like MongoDb, Cassandra, etc. The process of porting organized data to the next database is also called migrating.

## Using Golang Migrations
**How it works**
We use the following library to handle migrations: https://github.com/golang-migrate/migrate
All migrations are defined by files in `database/migrations`.
You should notice a couple of files in the form of `<DB_VERSION_NUMBER>.up.sql` and `<DB_VERSION_NUMBER>.down.sql`.
These files are necessary to "upgrade" or "downgrade" your database version in order to keep your database environment in-sync with other environments.

### To move in between database versions, use this:
Navigate to the `orion` root folder.
Downgrade your database by one version
```
migrate -source file://database/migrations -database mysql://USER:PASSWORD@\(localhost\)/mathnavdb down 1
```
Upgrade your database by one version
```
migrate -source file://database/migrations -database mysql://USER:PASSWORD@\(localhost\)/mathnavdb up 1
```

### Creating a new migration
To create a new migration, you'll need to create 2 new files
 - <DB_VERSION_NUMBER>.up.sql
 - <DB_VERSION_NUMBER>.down.sql

Make sure your `DB_VERSION_NUMBER` is the NEXT version number after what is available. For example, if the last version number is 000008, your next version should be 000009.

**Your up migration and down migrations must counteract each other!**
What this means is that if you apply an up migration step and then a down migration step, you should be EXACTLY at the original state.
Examples:
 - Creating a new table: Your up migration has a `CREATE TABLE` statement while your down migration has a `DROP TABLE` statement.
 - Adding a new column: Your up migration has a `ADD COLUMN` while your down migration has a `DROP COLUMN` statement.

Do remember that when you drop tables / columns, all your information is lost. So downgrading database can be dangerous if you have sensitive information you want to keep.

### Testing your migration
Use the `up` and `down` statements to upgrade to your latest version and go back to the previous version. Your job is to make these transitions stable.
Sign in to your mysql and check your database state by using:
```
USE mathnavdb;
SHOW TABLES;
DESCRIBE <TABLE_NAME>;
```
