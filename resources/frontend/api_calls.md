# Intro to API Calls on Gemini

# Review on Javascript Promises

A javascript `Promise` is used for asynchronous operations in JavaScript, events that occur independently from the main program.

Javascript promises have 3 possible states:
 - pending - the starting state where the operation is still pending
 - fulfilled - the operation completes successfully
 - rejected - the operation fails

Promises are resolved or settled when they are no longer pending.

# .then() and .catch()

The method `.then()` is used when the promise is fulfilled, while `.catch()` is used when the promise is rejected. `.finally()` is called when the promise is resolved.

The function inside the method is called:

```javascript
new Promise(someFunc)
    .then((res) => console.log("Success: " + res));
    .catch((err) => console.log("Error: " + err));
    .finally(() => console.log("Runs at the end"));
```

These 3 methods also return a promise object, which can be used for chaining promises.

```javascript
new Promise(someFunc))
  .then(handleFulfilled1)
  .then(handleFulfilled2)
  .then(handleFulfilled3)
  .catch(handleRejected)
```

# API

On gemini, in `api.js`, we use `axios` to create API requests to orion. In order to use this, add

```javascript
import API from "../api.js";
```

to the top of a file. This allows us to use `API.get()`, `API.post`, and `API.delete()`, which send a GET, POST, and DELETE request respectively.

Since these are promises, we can use `.then()`, `.catch()`, and `.finally()` to handle the responses.

`API.get()` and `API.delete()` take in 1 parameter, the api endpoint. For example, if we want to send a GET request to `api/programs/all`, we can use

```javascript
API.get("api/programs/all")
    .then((res) => console.log("Programs: " + res))
    .catch((err) => console.log("Error: " + err))
```

`API.post` takes in 2 parameters, the API endpoint and the body of the request. For example, we can call

```javascript
const program = { ... };
API.post("api/programs/create", program)
    .then(() => console.log("Success"))
    .catch((err) => console.log("Error: " + err))
```

# Successive API calls

In `api.js`, there is a `executeApiCalls()` function that allows us to execute a list of API calls. This function takes in 3 parameters, a list of api calls, a success callback, and a fail callback. We can import this using

```javascript
import API, { executeApiCalls } from "../api.js";
```

`executeApiCalls()` executes the API calls in order. If they all succeed, the success callback is executed. Otherwise, the fail callback is executed.

```javascript
const apiCalls = [ ... ];
let successCallback = () => console.log("Success!");
let failCallback = () => console.log("Error: " + err);

executeApiCalls(apiCalls, successCallback, failCallback);
```

This is typically used when there are a series of POST or DELETE API calls that need to be made. The function does not give output, so for GET requests, read the below section.

# Many API calls with Axios

If we want to send a series of GET requests at once, we can import `axios`:

```javascript
import axios from "axios";
```

The `.all()` method takes in a list of API calls and returns a promise. We can then use `.then()`, `.catch()`, and `.finally()` as usual.

`.spread()` gets all the responses and executes the function that is passed.

```javascript
const apiCalls = [ ... ];
axios
    .all(apiCalls)
    .then(
        axios.spread((...responses) => {
            const firstResponse = responses[0].data;
            const secondResponse = responses[1].data;
            console.log(firstResponse);
            console.log(secondResponse);
        })
    )
    .catch((err) => console.log("Error: " + err));
```
