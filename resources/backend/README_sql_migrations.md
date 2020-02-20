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
