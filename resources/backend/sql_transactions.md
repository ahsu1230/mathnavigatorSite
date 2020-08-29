## What is a SQL transaction?

A SQL transaction is a way to use commands in the database, but not have them actually affect a main database -- sort of a testing case.

## How to use a SQL Transaction

To create the testing database, use `transaction = db.Begin()`  
To execute commands, continue to use `db.Execute('<Command>')`  
To commit the database, use `transaction.Commit()`