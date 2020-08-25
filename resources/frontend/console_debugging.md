## Debugging in Browser
For the front-end, a lot of debugging is done by using inspect element on a browser. One of the tools this allows you to use is the console.

### How do you see the Console?
In order to see the console, you first need to inspect the webpage you are on. To do this, right click anywher on the page and click inspect.

![alt text][logo]

[logo]: ..\images\chrome-right-click-inspect.png "Logo Title Text 2"

Console will be the second tab from the left.

### How do you use the Console?
The console is a very useful tool in front-end debugging. When errors occur, they are detailed here. You can type commands into the console and execute them, allowing you to make changes to the state of the site. 

### Using console.log and debugger in code
Often, when there are errors in the code, nothing will happen. There are no error messages, and you might not know where to search for the errors. In these situations, you can use debugger. "debugger" is a keyword which stops the execution of JavaScript at the point it is called. This creates a breakpoint in which the code will stop executing past the breakpoint. This allows you to check the status of the site at different points to try to narrow down and find the error.

For more information click [here](https://www.w3schools.com/js/js_debugging.asp).