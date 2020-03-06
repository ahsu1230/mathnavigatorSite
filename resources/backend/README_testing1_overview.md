# How does testing work in Golang?

In any project, tests are imperative to making sure our system works. It's the "proof" that allows us to draw confidence that what we code indeed does what we expect. For the backend side, there are a few ways we can run "tests" for ourselves.

### 1. Check our database
We could, for example, log into MySQL and check that after every operation, our MySQL tables have the expected rows and values.

### 2. Send HTTP requests with curl
We could run curl commands against our webserver. This helps us check that the webserver is listening to requests and correctly processing them. If these HTTP requests are changing any values, we can double check our database to see what we spit out is exactly what we expect.

### 3. Run our unit & integration test frameworks
Both ways above are good ways to "sanity-check" our webserver functionality, but they can be time-consuming. The most robust and time-efficient way is to write code to test our code! And this is where unit and integration tests come in to help us. They are easier to run and can be easily run over and over again to make sure any little changes we make are not causing unexpected results!

## How to write good tests?
Just like regular coding, make sure your unit tests are testing for particular situations, especially the weird ones! Your job is to make tests cover as many weird cases as possible! However, at the same time, you don't want them to be repetitive.

For instance, let's say I wanted to test an add function, which takes 2 parameters. Maybe I could test the following situations:
```
add(2,3) == 5
add(3,2) == 5
add(3,0) == 3
add(0,3) == 3
add(3,-1) == 2
add(-3,1) == -2
add(3,-3) == 0
add(0,0) == 0
```
Here, we are testing many properties of addition. We have tests for regular positive numbers, negative numbers, the reflexive and inverse properties of addition, etc.

However, the following additional tests might be repetitive.
```
add(5,4) == 9
add(4,5) == 9
add(1,1) == 2
```
These unit tests would NOT be necessary because they already test what we've already tested (adding positive numbers, adding repeated numbers).