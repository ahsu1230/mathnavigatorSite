# CURL

## Review the following here: [protocols](https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/README_2_protocols.md).
 - API (Application Programming Interface)
 - HTTP (Hypertext Transfer Protocol)
 - JSON (Javascript Object Notation)

A webserver has something called *API endpoints*. These endpoints are specific locations where the webserver accepts communication. Usually a web client will send a HTTP request to one of these endpoints. When this HTTP request is picked up by the webserver, it will proceed with its internal processing, and once finished, create a HTTP response to send back to the web client.

Here's an example with `orion`.
Go ahead and start `orion`.
When you start `orion`, you should see something like this:
```
[GIN] GET    /api/programs/all      --> .../programs.GetPrograms
[GIN] POST   /api/programs/create   --> .../programs.CreateProgram
[GIN] GET    /api/programs/program/:programId --> .../programs.GetProgram
[GIN] POST   /api/programs/program/:programId --> .../programs.UpdateProgram
[GIN] DELETE /api/programs/program/:programId --> .../programs.DeleteProgram
[GIN] Listening and serving HTTP on :8080
```
The format of each line looks like this: `REQUEST_VERB` `API_ENDPOINT` --> `SERVER_FUNCTION`.
So if we take the first line, it means that `orion` has an API endpoint at location: `/api/programs/all` that accepts a `GET` HTTP request.
The full url would look like this: `http://localhost:8001/api/programs/all`.
When this endpoint receives a valid request, the server then will:
 - call the function `GetPrograms()` somewhere in its server code.
 - create a HTTP response (i.e. `200 OK`)
 - send response back to web client who sent the request (usually in JSON)

Normally, these HTTP requests come from browsers, mobile phones, etc. which require building web applications. However, we can use a tool called [CURL](https://en.wikipedia.org/wiki/CURL) that acts as an extremely simplified web client.

## CURL
Usually, curl is built into MacOS and Windows 10+. If you do not have it, install it [here](https://www.youtube.com/watch?v=4QNWUbrLpCk).

If you go to Terminal / Command Prompt, call:
```
curl --version
```
To see if you have it correctly installed.

### What is CURL?
CURL is a simple command line tool that acts as a VERY simplified web client. Its job is to basically turn command line parameters into a HTTP request to send to the target URL. If we want to test HTTP requests and responses in an easy, detailed way, CURL is one of the best tools to use!

### Example with Google
When we go to https://www.google.com in our browser, when we input the URL, we are actually sending a GET request to that URL.
Try to do the same in CURL. Send a CURL request to Google.
```
curl -X GET https://www.google.com
```
What you'll get next is a lot of stuff - which is the HTML + Javascript from the Google homepage!
Our browsers are actually doing the exact same thing, but then turning that HTML into what you see with the Google logo / pictures / search box.

### CURL with Orion

While Orion is still listening for HTTP requests, you can try the following command in a *separate* tab.
```
curl -X GET http://localhost:8001/api/programs/all
```
What you get back should be a list of your current programs from the database in JSON format.
Go back to your Orion tab and you should see something like this:
```
[GIN] ... | 200 |     319.055µs |             ::1 | GET      /api/programs/all
```
This tells us which endpoint was just hit, when, and how long it took (319 microseconds). The number 200 means the HTTP status OK. 4xx statuses mean an error occurred while executing the endpoint.

### Other CURL examples with Orion

Here's an example of a POST request. With a POST request, you will be sending JSON to api endpoint: `http://localhost:8001/api/programs/create`. The JSON (what we call body) contains information about the new program you want to create.
```
curl -X POST -H "Content-Type: application/json" --data '{"programId": "createdProg1", "name": "Created Program", "grade1": 6, "grade2": 8, "description": "Some Newly Created Program"}' http://localhost:8001/api/programs/create
```

Here's how to get a specific program. Observe how the `programId` we used in the Create endpoint is now here. `gin` uses url parsing to figure out that `createdProg1` is actually a variable value, which we can then use to retrieve database rows.
```
curl -X GET http://localhost:8001/api/programs/program/programId
```

Here's a DELETE example. Again, you can use `createdProg1` as the `PROGRAM_ID`.
```
curl -X DELETE http://localhost:8001/api/programs/program/<PROGRAM_ID>
```

Check the website to see your changes.
Check MySQL to see what your database now looks like.
