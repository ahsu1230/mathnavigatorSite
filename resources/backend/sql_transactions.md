## What is a SQL transaction?

TODO: Explain more about this: https://stackoverflow.com/questions/974596/what-is-a-database-transaction

## How to use a SQL Transaction in Golang

TODO: Explain more!

To create the testing database, use `transaction = db.Begin()`  
To execute commands, continue to use `db.Execute('<Command>')`  
To commit the database, use `transaction.Commit()`

```go
tx := db.Begin()
err := db.Exec(Command to do actionA)
if err != nil { // Command actionA failed, rollback transaction
   tx.Rollback()
   return err
}
err := db.Exec(Command to do actionB)
if err != nil { // Command actionB failed, rollback transaction
   tx.Rollback()
   return err
}

// Everything is a success, fully commit db transaction!
tx.Commit()
```