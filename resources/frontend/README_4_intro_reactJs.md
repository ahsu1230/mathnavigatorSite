## What is ReactJs?
ReactJs is a web Javascript framework created by Facebook.
Web frameworks are Javascript libraries that have a lot of built-in functions to do stuff way easier.

Before frameworks, we had jQuery. But even creating common layouts (like single-page applications) was pretty hard to do. [Single-page vs. Multiple-page applications](
https://medium.com/@goldybenedict/single-page-applications-vs-multiple-page-applications-do-you-really-need-an-spa-cf60825232a3).

Web frameworks were created to make developer lives easier to easily and quickly create applications in mind.

Blogpost:
https://medium.com/javascript-in-plain-english/how-i-learned-react-js-as-a-noob-ultimate-react-js-starter-guide-36a05ab9495e

Tutorial:
https://reactjs.org/tutorial/tutorial.html
https://reactjs.org/docs/getting-started.html

## What are React Components?
A React Component is simply a part of a single page application. You can think of a web page as object-oriented. For example, usually you will see a header with links, a footer, or a profile page on the left side. These can all be thought of as different areas that together form the entire webpage.

A React Component can be any of these layouts and every Component can have multiple Components inside it to establish parent-children relationships. When these relationships are established, it's easier to create "ripple" effects so that when a user interacts with one component, another can be affected as well.

For example, maybe I'd like to click on an image (one React Component) to zoom in on it.
When that happens, you can have other parts of the page (other React Components) move away or disappear so you can get a better view of the image.

## Differences between State and Props
