### Tutorials

https://www.khanacademy.org/computing/computer-programming/sql

References:
https://www.w3schools.com/sql/sql_intro.asp

### Review the following [here](https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/README_5_databases.md):
 - What is SQL?
 - What is MySQL?

## Try it yourself!
We will first use MySQL in Terminal/Shell so we grasp the basics of using command line functions.
Once you understand these basic commands, then you can see what MySQL Workbench is doing under the hood.

Using MySQL in Terminal or MySQL Shell...
First sign in using the following:
```
mysql -u root -p
```
And enter your password.

To start, you should create a database called `testDb`. Using `USE`, we will set `testDb` as our current database.
A database in this context means a collection of tables.
When we use `SHOW tables`, we should get an empty list because it's new!
```
CREATE DATABASE testDb;
USE testDb;
SHOW tables;
```

Create a new table called `persons` using the following command:
```
CREATE TABLE persons (
    person_id int,
    last_name varchar(255),
    first_name varchar(255),
    address varchar(255),
    city varchar(255)
);
```
This statement creates a table called `persons` which has 5 columns, (`person_id`, `last_name`, `first_name`, `address`, `city`).
This means that every row in this table (representing a person) will have 5 properties: an id, last name, first name, address and city.
In databases, it's somewhat standard to use snake_case to represent variables names.
Once this table is created correctly, you can use `SHOW tables;` to display all tables in the database `testDb`.
You should see table `persons` in the list.

Let's start adding data! You this command to add rows.
```
INSERT INTO persons (person_id, last_name, first_name, address, city) VALUES (1, "Ketchum", "Ash", "somewhere", "some city");
```
Add 4 more rows with varying values to get more data into our database!
Once you are finished, you can use this command to see all of them:
```
SELECT * FROM persons;
```

The above line selects ALL rows in table `persons`.
If we wanted to select particular rows or columns, we would use the `WHERE` clause like so:
```
SELECT * FROM persons WHERE person_id=1;
SELECT * FROM persons WHERE first_name="Ash";
SELECT address, city FROM persons WHERE last_name="Ketchum";
```

To update certain rows, you can use this:
```
UPDATE PERSONS SET address="Route 23", city="Viridian" WHERE person_id=1;
```
To delete certain rows, you can use this:
```
DELETE FROM PERSONS WHERE person_id=1;
```
After playing around with these commands, use `SELECT * FROM persons;` to see how your table looks after your changes.

When you are finished, you can use the following commands to:
 - delete table `persons`
 - delete database `testDb`
 - quit out of MySQL Terminal/Shell
```
DROP TABLE persons;
DROP DATABASE testDb;
exit;
```

Now open MySQL Workbench, and try the same as above!
MySQL Workbench allows you to do everything here but with a GUI to manage your databases and tables.
