# CURL

## Review the following here: [protocols](https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/README_2_protocols.md).
 - API (Application Programming Interface)
 - HTTP (Hypertext Transfer Protocol)
 - JSON (Javascript Object Notation)

A webserver has something called API endpoints. These endpoints are specific locations where the webserver accepts communication. Usually a web client will send a HTTP request to one of these endpoints. When this HTTP request is picked up by the webserver, it will proceed with its internal processing, and once finished, create a HTTP response to send back to the web client.

Here's an example with `orion`.
When you start `orion`, you should see something like this:
```
[GIN] GET    /api/programs/v1/all      --> .../programs.GetPrograms
[GIN] POST   /api/programs/v1/create   --> .../programs.CreateProgram
[GIN] GET    /api/programs/v1/program/:programId --> .../programs.GetProgram
[GIN] POST   /api/programs/v1/program/:programId --> .../programs.UpdateProgram
[GIN] DELETE /api/programs/v1/program/:programId --> .../programs.DeleteProgram
[GIN] Listening and serving HTTP on :8080
```
The format of each line looks like this: `REQUEST_VERB` `API_ENDPOINT` --> `SERVER_FUNCTION`.
So if we take the first line, it means that `orion` has an API endpoint at location: `/api/programs/v1/all` that accepts a `GET` HTTP request.
The full url would look like this: `http://localhost:8080/api/programs/v1/all`.
When this endpoint receives a valid request, the server then will:
 - call the function `GetPrograms()` somewhere in its server code.
 - create a HTTP response (i.e. `200 OK`)
 - send response back to web client who sent the request (usually in JSON)

Normally, these HTTP requests come from browsers, mobile phones, etc. which require building web applications. However, we can use a tool called [CURL](https://en.wikipedia.org/wiki/CURL) that acts as an extremely simplified web client.

## CURL
