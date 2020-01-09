# Math Navigator Resources
Here are some basic Q&A on working on web applications in general.

## Web Applications
Web applications are computer programs that utilize web browsers to perform tasks over the Internet. Web applications usually involve two participants. One participant (the web client) initializes a "request" while the other participant (the web server) "responds" to the request with data. As the user uses the web applications, various requests from the client are sent to the web server. The web server responds with data which the client can then use.

### What is a web client?
A web client is usually an Internet browser (Google Chrome, Firefox, Internet Explorer, etc.) or a mobile application / phone / tablet. Usually, a web client starts the interaction with a web application by sending requests to a web server. For example, going to https://www.facebook.com will contact a Facebook web server to send you the Facebook home page.

### What is a web server?
A web server is a large, highly performant computer owned by a company. A web server's job is to receive requests from various web clients and then deliver information to those clients based on the request. For example, if a client requests user A's profile, the server should deliver ONLY user A's profile content to that client.

Here's a short video about web clients vs. servers: https://www.youtube.com/watch?v=QSEDr2e1gSQ

### What does front-end and back-end mean?
As you get closer to job searching, you'll often hear about Front-End developers and Back-End developers. Front-end developers work mostly on the web client in an application, while Back-End developers work on the web server. 

Front-end developers may work on websites and mobile applications, and often work with a design or product team to create the best possible experience for users. The languages that Front-end developers use are usually Javascript web frameworks such as AngularJs, ReactJs or languages specific to mobile phones like Android Java or Swift (iOS).

Back-end developers work on the web server. Their job is usually to deliver data to web clients in a timely and efficient manner. Back-end developers use a lot of algorithmic analysis to keep their servers running effectively, no matter the situation. They work with front-end teams and data analyzers to make sure servers are optimized for their needs. Languages that Back-end developers use can be anything from Java, Python, C++, Go, NodeJs, etc.

You can think of it like working in a restaurant. The waiters and waitresses are the front-end developers who make sure their diners are getting the best experience possible while the chefs are the back-end developers, who make the delicious food!

### What is a database?
A database is a huge computer or a large collection of computers that are tasked with storing information. Their job is to not only keep large sets of information safe at all times, but to also be able to retrieve information quickly! Databases usually have their own language such as SQL, which isn't really a programming language. Database languages are built primarily to interact with databases.

## Protocols
In web applications, in order for the web client and web server to communicate with each other, a standard must be agreed upon. The most common protocol on the Internet today is the [HTTP protocol](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol). 

All requests from a web client and responses from a web server are done over HTTP. So, a web client communicates to a web server using HTTP requests. A web server responds to a web client using HTTP responses.

### What does HTTP and HTTPS mean?
HTTP stands for Hypertext Transfer Protocol. HTTPS stands for HTTP Secure. HTTPS has built-in security functionalities that make it harder for malicious hackers to intercept important transactions. You can read more about it [here](https://en.wikipedia.org/wiki/HTTPS).

The main HTTP request methods are `GET`, `POST`, `DELETE`.

### What is an API?
API stands for Application Programming Interface. It's a fancy way of explaining how different parts of a computer program are supposed to communicate each other. In this case, the API dictates how the web client will interact with the web server.

Here is an example:
Suppose a client is on www.facebook.com. When that client wants to see Amy's profile, Facebook might send her to the following url:
```
https://www.facebook.com/amy
```
When the user wants to see Bob's profile, Facebook might send her to this url:
```
https://www.facebook.com/bob
```
As you can see, the pattern to visiting someone's profile looks like this: `https://www.facebook.com/<USER_ID>` where USER_ID is some kind of unique identifier for a particular user.

Now, what if I would like to see all friends of both Amy and Bob? Facebook might send the user to this url:
```
https://www.facebook.com/amy/friends
https://www.facebook.com/bob/friends
```
The pattern to visit someone's friends is: `https://www.facebook.com/<USER_ID>/friends`.

APIs are established so both client and server knows what information is being requested and what information to respond with. In this example, the web client and server interaction looks like the following:

 - client (IP address 1.2.3.4) initializes communication by sending HTTP GET request to server at `https://www.facebook.com/amy/friends`.
 - server receives a GET HTTP request (`https://www.facebook.com/amy/friends`).
 - server realizes that it needs to find all friends for user Amy.
 - server looks through its files and databases to find all friends of Amy.
 - server finds all friends of Amy and packages it neatly for client (in JSON format).
 - server responds to client (IP address 1.2.3.4) by sending a HTTP response.
 - client receives a response from the server. The response contains a list of all of Amy's friends.
 - client changes its website layout to display all Amy's friends to the user.

Every interaction between a client and server usually follows this procedure. And all of this happens in about 200ms, but maybe longer depending on how busy the server is (handling many other clients at the same time) or how bad your Internet connection is :). 

### What is JSON?
JSON stands for Javascript Object Notation. JSON is a standardized Javascript format which can be easily read by both humans and computers. JSON has been the modern standard for HTTP API communication between clients and servers. In regards to the example above with a user sending a GET request to `https://www.facebook.com/amy/friends`, the response JSON from the server may look like this:

```
{
  user: "amy",
  friends: [
    {
      user: "chris",
      dateFirstMet: "12/4/2000"
    },
    {
      user: "dave",
      dateFirstMet: "6/21/2018"
    },
    {
      user: "ellen",
      dateFirstMet: "1/17/1998"
    }
  ]
}
```

As you can see, it resembles a Javascript Object.
