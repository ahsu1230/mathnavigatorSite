# SQL Advanced topics

## What is a GROUP BY?
## What is HAVING?

## What is an Index?
## What is a Constraint?
[SQL constraints](https://www.w3schools.com/sql/sql_constraints.asp) are essentially rules that limit what data can be inserted into a table. Constraints are used to ensure that the data in a table is valid. This increases the reliability of data and allows people more control over their databases. 

If a constraint is not met, SQL will throw an error and the action will be stopped.

We will use the following `example` table to demonstrate each constraint in SQL:
```mysql
CREATE TABLE example
(
    id     int NOT NULL UNIQUE,
    number int DEFAULT 0,
    CHECK (number < 10),
    PRIMARY KEY (id)
);
```

#### PRIMARY KEY

The PRIMARY KEY identifies each row in the table. Primary keys must be unique and cannot be null. In `example`, the PRIMARY KEY is `id`. A primary key can consist of one column or multiple columns.

#### NOT NULL

The NOT NULL constraint ensures that a column cannot have NULL values.

Attempting to insert null into `id` yields: `ERROR 1048 (23000): Column 'id' cannot be null`.

#### UNIQUE

The UNIQUE constraint ensures that each value in the column is unique. The PRIMARY KEY constraint automatically makes `id` UNIQUE and NOT NULL. 
Adding a duplicate value yields: `ERROR 1062 (23000): Duplicate entry '1' for key 'example.id'`

#### DEFAULT

The DEFAULT constraint specifies a value if one is not provided. Running `INSERT INTO example (id) VALUES (1);` will create a row with `id` having a value of 1 and `number` having a value of 0.

#### CHECK

The CHECK constraint only allows certain values in a column. In `example`, the check constraint ensures that `number` is less than 10.

Violating the check constraint yields: `ERROR 3819 (HY000): Check constraint 'example_chk_1' is violated.`

#### FOREIGN KEY

See below

## What is a Foreign Key?

Foreign keys links 2 tables together. A foreign key in one table refers to the primary key in another table. The FOREIGN KEY constraint prevents actions that would destroy the link.

Foreign keys are useful because it allows us to establish complex relationships between tables. In real world databases, tables often refer to or depend on each other. Foreign keys allow us to accomplish this easily.

Continuing the example from before, the `foreign_example` table below has a foreign key called `example_id` that refers to the `id` column in the `example` table.

```mysql
CREATE TABLE foreign_example
(
    id         int,
    example_id int,
    PRIMARY KEY (id),
    FOREIGN KEY (example_id) REFERENCES example (id)
);
```

A row can be inserted into `foreign_example` normally, but for the `example_id` column, the corresponding `id` should be used. Once a row is inserted, the `example` row it references cannot be updated or deleted. Attempting to do so yields 

``ERROR 1451 (23000): Cannot delete or update a parent row: a foreign key constraint fails (`mathnavdb`.`foreign_example`, CONSTRAINT `foreign_example_ibfk_1` FOREIGN KEY (`example_id`) REFERENCES `example` (\`id\`))``

and

`ERROR 3730 (HY000): Cannot drop table 'example' referenced by a foreign key constraint 'foreign_example_ibfk_1' on table 'foreign_example'.`

respectively. In order to delete a row that is being referenced in a foreign key, all the foreign key references must be deleted.

In this example, let's say we create an `example` row with an `id` of 1 and then create several `foreign_example` rows referencing the `example` row. In order to delete the `example` row, all the `foreign_example` rows referencing the `example` row must be deleted first.

#### Examples in Our Codebase

In Math Navigator, classes have a program, semester, and a location. Since multiple classes can refer to a program, semester, or a location, it makes sense to use foreign keys.

This also ensures that if any program, semester, or location is being referenced by a class, it cannot be updated or deleted. This ensures the integrity of the database. [Here](https://github.com/ahsu1230/mathnavigatorSite/blob/master/orion/pkg/repos/migrations/000006_create_table_classes.up.sql) is the `classes table`:

```mysql
CREATE TABLE classes
(
    id          int unsigned NOT NULL AUTO_INCREMENT,
    created_at  datetime     NOT NULL,
    updated_at  datetime     NOT NULL,
    deleted_at  datetime,
    program_id  varchar(64)  NOT NULL,
    semester_id varchar(64)  NOT NULL,
    class_key   varchar(64),
    class_id    varchar(192) NOT NULL UNIQUE,
    location_id      varchar(64)  NOT NULL,
    times       varchar(64)  NOT NULL,
    start_date  date         NOT NULL,
    end_date    date         NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (program_id) REFERENCES programs (program_id),
    FOREIGN KEY (semester_id) REFERENCES semesters (semester_id),
    FOREIGN KEY (location_id) REFERENCES locations (location_id)
) AUTO_INCREMENT = 1
  DEFAULT CHARSET = UTF8MB4;
```

## What is a JOIN?

SQL JOIN clauses allow us to combine rows from different tables as long as those rows share a common column. If two rows in two different tables have the same value in some common column, then we can select any other column values of that row from either table.

If our two tables are named `table1` and `table2` respectively, then the syntax is as follows:

```
SELECT table1.column1, table1.column2, ..., table2.columnA, table2.columnB, ...
FROM table1 (TYPE OF JOIN) table2
ON table1.sharedColumn=table2.sharedColumn;
```

Let's break this down line by line.

`SELECT table1.column1, table1.column2, ..., table2.columnA, table2.columnB, ...`

This is pretty standard SELECT syntax. We are indicating which columns we want to return. The only difference here is that we are mixing two different tables, and we are attempting to grab different columns from both tables.

`FROM table1 (TYPE OF JOIN) table2`

First off, `table1` and `table2` are the two tables we are joining. Next is `(TYPE OF JOIN)`: there is more than one type of JOIN clause. The `(TYPE OF JOIN)` is simply a filler; you don't literally type `(TYPE OF JOIN)`, but instead put whatever JOIN clause you are using (without the parenthesis). For example, if I use FULL OUTER JOIN, the syntax will look like this:

`FROM table1 FULL OUTER JOIN table2`

The types of joins will be explained in the next section.

`ON table1.sharedColumn=table2.sharedColumn;`

Here is where we designate our shared column. There is actually no requirement for the shared column to have the same name in both tables. As long as they share the same data type, then your SELECT statement will successfully return all rows with the same value in that shared column in both tables.

## There are actually 4 types of JOINS!

 - Inner - Returns all records with matching values in both tables.
 - Left (Outer) - Returns all records in the left table and all records with matching values in the right table.
 - Right (Outer) - Returns all records in the right table and all records with matching values in the left table.
 - Full Outer - Returns all records with a match in either the left or right table.

The specific syntax for each of these can be found on [w3schools](https://www.w3schools.com/sql/sql_join.asp) (the syntax for all four are very similar, so I won't copy paste them here).
