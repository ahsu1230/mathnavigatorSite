# Orion

Our core API service. This webserver will provide the main API to allow web clients to interact with our MySQL database. This is where users, programs, classes, announcements are created and persisted into the database.

## How to run the Orion webserver

### Pre-requisites

- Install Go (https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/onboarding/install_go.md)
- Install MySQL (https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/onboarding/install_mysql.md)

Ensure your MySQL server is running. For MacOs, it is the `mysql.server start` command and for Windows, it is the `net start MySQL`. **Note (Windows):** In order to run this, you need to run Command Prompt as administrator. To do so, right click the Command Prompt application and select "Run as administrator."

---

### Run Tests

To run all tests for the back-end web server, run:
```
go test ./...
```
You should see `ok`s and no failures.

### Run the Webserver

In the `orion` folder,
 * Create a new folder called `configs`.
 * Inside this folder, create a new file called `config_local.yaml`.
 * Paste the following content into this file and save.
```
app:
  build: "development"
  corsOrigin: "*"
database:
  host: "localhost"
  port: 3306
  user: "root"
  pass: "<YOUR_PASSWORD_GOES_HERE>"
  dbName: "mathnavdb"
```
Remember the password you saved for MySql? Paste that password where it says `<YOUR_PASSWORD_GOES_HERE>`!

After that, go back to the `orion` directory and start the web server with this:
```
go run main.go configs/config_local.yml
```
You should see a `Listening and serving HTTP on :8080` message. It worked! Now, you are running the Math Navigator webserver locally on your machine. Any HTTP requests to the port number 8080 will be received and responded to by the local webserver!
