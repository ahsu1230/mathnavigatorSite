# Integration tests

## What's the difference between unit tests and integration tests?

Unit tests are for testing the functionality of **one, isolated component**. In our webserver, we'll probably have a lot of layers and different classes/packages have different roles. A unit test's job is to make sure that each ONE package is doing its own job correctly. On the other hand, an integration test's job is to make sure that multiple packages and layers are playing well together - or are working together do create the correct, final outcome.

Here's an example. Let's say in our webserver, we have a controller layer and a repo layer.
A unit test should be testing only the controller's logic i.e. if I give a bad HTTP request, does it return a BadRequest status code in the HTTP response? 
Another unit test could be testing only the repo logic i.e. if I call this repo function, is it calling the correct SQL statement?

An integration test would test that if I give a certain HTTP request, does it successfully update the database with the correct values AND does it successfully return the correct response? The integration tests would try to simulate a more "end-to-end" testing than the unit tests.

## How do integration tests work in our project?
The integration test is most useful in our CircleCI flow.
Every time a developer pushes code into a Pull Request, a CircleCI build is spun up to start testing on a separate cloud machine. In general, the CircleCI machine will download a certain version of Golang, install a specific version of MySQL, and replicate a real webserver-database interaction by running all the integration tests.

For us, we can run integration tests locally because we already have MySQL installed on our machine. Using integration tests ensure that all our components (Golang, MySql, etc.) and their respective versions are playing well together. Integration tests are also run when you run `go test ./...` and files can be found in the test_integration folder.