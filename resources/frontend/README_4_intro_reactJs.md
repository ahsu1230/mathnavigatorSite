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

## Every Component has a render() function. What does it do?
The `render()` function of every Component returns HTML. Inside, a `render`, you can do some computations to return a single parent HTML element (like a `<div>`). You can also use variables that you computed to insert into the <div>. For example, if you have a custom HTML header, you could do something like this:
```
<h1>{headerText}</h1>
```
The value of variable `headerText` will then take its place inside the `<h1>` tag.
It is important to note that render() is the most called function of the Component! So, you will have to make sure your render() is as "bare-bones" as possible. In other words, if computation can be moved out of the render(), it should.

## What are the React Component Lifecycles?
 - **constructor()**
   - called when the Component is first created.
 - **componentDidMount()**
   - called before the first `render()` and when the newly created Component is attached to the DOM. This is a great place to make asynchronous HTTP requests or loads.
 - **componentDidUpdate()**
   - called before any `render()`. You can use this lifecycle function to check if your Component values will change (based on state or props).
 - **render()**
   - the main HTML rendering function. This method is in charge of producing what your user will visually see on their browser. The HTML they may see will depend on the component's states and properties (props).
 - **componentWillUnmount()**
   - called after render() and before this Component is destroyed. This happens when the component is no longer needed and detached from the DOM.

## Differences between State and Props
You've probably seen `this.state.____` or `this.props.____` in some of the React Components.
State variables are used when a React Component can be changed internally. For example, let's say I have a Component with a On/Off button inside it. When a user clicks on this button, it might change the state of the button itself from On -> Off or Off -> On. Since this button is inside the Component, this is a `state` example.
I can change the state of the button using this:
```
this.setState({ isButtonOn: true });
```
From here, in the `render()` function, I can do something different depending on the button on state.
```
render() {
    const isButtonOn = this.state.isButtonOn;
    if (isButtonOn) {
        ...
    } else {
        ...
    }
    return (
        ...
    );
}
```

Property variables are used when a React Component is changed externally by another Component. When Components have parent-children relationships, you can "pass" values from a parent Component to children Components. The children Components will inherit the passed values as "properties".
So for instance, let's say clicking on the button example from above, changes the rest of the page.
You can "pass" the button state on the child's Component creation:
```
<Child buttonOn={this.state.isButtonOn}/>
```

And inside the Child's `render()` function, you can compute on the parameter like so:
```
render() {
    const parentButtonOn = this.props.buttonOn;
    ...
    return (
        ...
    );
}
```
