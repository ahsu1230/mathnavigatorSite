# Enzyme Tests with ReactJs

You may have seen either from CircleCI or in a project's `package.json` file that you can run `npm run test` and the program will run front-end tests for a website.

For these projects, we use [EnzymeJs](https://enzymejs.github.io/enzyme/) which is a ReactJs web development testing framework created by AirBnB.

## Testing the Front-End

Testing the front-end is a bit trickier than the back-end. Of course you can test utlity functions that do certain calculations, etc. But more in-depth front-end is testing if components look okay or behave correctly after user interaction.

Enzyme is great for testing if components are always rendering certain HTML elements. Or for testing if components had `props` or `state` value of X, does it behave correctly? Or for testing if a user clicked on a button, does the component change its state variables, etc.

In certain cases, Enzyme can take "snapshots" of a "correct" look and compare that snapshot to what the website currently looks like. However, this isn't totally foolproof and at the end of the day, some manual testing for front-end devices is still required to be confident that your website works 100% great.

## Testing a React Component with Enzyme

Ideally, to test with ReactJs, you want to write a test for every main React component that you create. To write an Enzyme test, your test file should be located in the same directory as the React component you want to test.

For instance, I would like to make an `achieve.test.js` which is located in the same folder as `achieve.js`. In this `achieve.js` file, the main component we want to test here is called `AchievePage` which is a fairly simple component with a title, a few paragraphs and a list of things. This list is fetched from an API call to the backend webserver.

### Importing a test component

The first thing we want to do is figure out how to import this component into the test. And this depends on how the component is available for export. If you look at the definition of AchievePage, you'll see:

```
export class AchievePage extends ...
```

This means `AchievePage` is one of the components being exported from this file. So in `achieve.test.js`, we can import `AchievePage` using `import { AchievePage } from "./achieve.js";`. 

However, on the other hand, there are some files that export ONLY a single component. In those cases, you'll see something like this (which usually means it's the only component being exported):

```
export default class AchievePage extends ...
```

If this is the case, use `import AchievePage from "./achieve.js"` to correctly import that component.

### What is shallow rendering vs. deep rendering?

The second thing we want to do is determine how we want to render our component. Do we want just the basic skeleton of the component? Or do we want to render the component, call all lifecycle methods, and render all of its children (and their children)?

For a unit testing, often times the latter is overkill - especially with components that make API calls to webservers. Those network queries can take a long time so calling a bunch of those in unit tests may not be worth it. For the most part, in our Enzyme tests, we'll be using *shallow rendering* in order to just test the basic skeleton of a component and its fundamental features.

Here's a good summary of shallow vs. full vs. static rendering in Enzyme tests <https://medium.com/@Yohanna/difference-between-enzymes-rendering-methods-f82108f49084>.

But basically, **use `shallow` to test the basic behavior of a component.** With `shallow` we can test if the component loads (constructor), and can test how the component behaves when it receives `props` or `state` variables. 

If you want to test the FULL behavior of a component (including children and lifecycle behavior), use `mount`. However, if you just want to see children HTML (no lifecycle behavior), use `render`.

### Writing an Enzyme Test

Enzyme tests are also built on top of the `jest` framework. In `jest` we can "describe" a group of tests with a name. Inside this "test group", we will define individual unit tests that each test a different thing about the component. In every unit test, make sure to render the component (using `shallow`, `mount`, or `render`).

After that, we can observe properties about this component using Enzyme's [shallow rendering API](https://enzymejs.github.io/enzyme/docs/api/shallow.html). And make assertions using Jests [expect library](https://jestjs.io/docs/en/expect).

```
describe("Achievement Page", () => {
    test("renders", () => {
        const component = shallow(<AchievePage />);
        expect(component.exists()).toBe(true);  
    });
})
```

How does this work? We render the component using `shallow` which gives us a basic rendered Enzyme wrapper around the `AchievePage` component. We can use the Enzyme API `.exists()` to check if the component actually rendered! This API returns a boolean.

From here, we use `expect` to check if the result of `component.exists()` which returns a boolean, is indeed equal to the value `true`. It if is, the test passes. Otherwise, the test fails. And that's it! `expect` works very similarly to assertions in Java, Ruby, and even in our Golang tests.

This is a very basic and dumb test, but can already tell us a lot. If there's a typo or Javascript error, this component will not render. Before, the only way to really be sure was to manually load the webpage and check that the page worked. With this one test, we can be confident that the component loads something and everything compiled correctly.

### Finding inner components with Shallow Rendering

Cool, so we've just figured out how to check if the component actually renders. We can extend the Enzyme features by searching for inner elements, or changing state/prop values of the component. Observe the following code:

```
test("renders", () => {
    expect(component.exists()).toBe(true);
    expect(component.find("h1").text()).toContain("All Achievements");
    expect(component.find("Link").text()).toContain("Add Achievement");
    expect(component.find("AchieveRow").length).toBe(0);
});
```

Here, we use Enzyme's `.find(...)` API function to look for specific elements in the `AchievePage` component. We look for `<h1>`, a `<Link>`, and `<AchieveRow>` and can even check the text inside those components. We then use `expect` to check if the text we found has a certain property (`.toContain` is a jest function that fails the test if the string does not contain the given substring).

When the component looks for `<AchieveRow>` we can get a list of child components. The last line in the test checks if by default, there are no `AchieveRow `components inside `AchievePage`. This is because it is a state variable which starts as an empty list.

**IMPORTANT NOTE** Shallow only renders "one level" deep when it comes to children. So in this case, `AchieveRow` could contain a bunch of inner HTML inside it. Shallow on the top level will not render what's inside. If we want to test what is inside `AchieveRow`, we will either need to use `render`/`mount` OR create unit tests that shallow render `AchieveRow`.

### Playing with component's state and props values

Before, we checked if `AchieveRow` had 0 occurences by default in `AchievePage` because the state variable in `AchievePage` starts as an empty list. Let's test the behavior of this state variable. Using Enzyme's API, we can set the state variable like this:

```
const testAchievements = [
    {
        Id: 1,
        year: 2020,
        message: "Awesome",
    },
    {
        Id: 2,
        year: 2019,
        message: "Possum",
    },
];
component.setState({ list: testAchievements });
```

Now the component's `state.list` variable now contains a list of two objects. We can use `.find()` to obtain a list of `AchieveRow`s and use `.at()` to get elements by index.

```
let row0 = component.find("AchieveRow").at(0);
expect(row0.prop("achieve")).toHaveProperty("Id", 1);
expect(row0.prop("achieve")).toHaveProperty("year", 2020);
```

What the above does is finds the first `AchieveRow` element in `AchievePage` and checks its passed props. Somewhere in AchievePage, we create rows like this:

```
const achievements = this.state.list.map((achieve, index) => {
    return <AchieveRow key={index} achieve={achieve} />;
});
```

Here, we grab the first `AchieveRow`, look at its property `achieve` (which is equal to the passed in achievementObject) and check to see if it has a property `Id` equal to 1. Then we check if its property `year` is equal to 2019.

We can do more using `setProps`. View the Enzyme docs to learn more about how to modify components for testing purposes
<https://enzymejs.github.io/enzyme/docs/api/ShallowWrapper/setProps.html>

### Triggering and testing an action

You can also trigger an event (onclick, onchange, etc.) and test the result of that event.

<https://enzymejs.github.io/enzyme/docs/api/ShallowWrapper/simulate.html>

---

## What are snapshots?

Snapshots are a Jest tool that double checks your UI is not changing unexpectedly. A snapshot of a component / web page is taken during a test and is compared to a previously marked "correct" snapshot. Snapshot files (content mostly looks like HTML short-hand) are stored locally (next to tests) and everytime a test is run, the component is rendered, a snapshot is taken, and then compared to the original "correct" snapshot.

Under the hood, the snapshot is just a tree of what HTML elements are in your component. It will double check constants are correct, certain elements are in the correct place, etc.

## Creating and testing a snapshot

To use a snapshot in a test, use the `.toMatchSnapshot()` Jest function.

```
import renderer from "react-test-renderer";

...

test("snapshot", () => {
    const tree = renderer
        .create(
            <Router>
                <Footer />
            </Router>
        )
        .toJSON();
    expect(tree).toMatchSnapshot();
});
```

The first time you run this, a snapshot image is created and saved into a `__snapshots__` folder. This folder should also be commited in your Git commit!

Subsequent runs of this test will take a snapshot of the current webpage, and compare the HTML positions of the snapshot that you find in `__snapshots__`. If they are different, the test will fail. If they are the same, the test succeeds!

If you make a change to the HTML, you can update the "correct" snapshot (the saved one) by using `jest --updateSnapshot`. This will update the save snapshot to the current snapshot your component is producing.

I recommend using snapshots when our front-end pages become more stable, so for now, only use snapshots for components that should not be changing often. For example, I have snapshots on the gemini-user `Footer` component because I believe it won't change in a while so we need tests to make sure that this component always renders what we have for the correct snapshot.