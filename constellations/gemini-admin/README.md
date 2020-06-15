# Gemini Admin

One of the Gemini website clients. This web client will interact with the Orion webserver by allowing MathNavigator administrators to manage student/parent information, available programs and schedule classes and announcements.

## Install Node & NPM

Make sure that npm is installed before proceeding. Install [NodeJs](https://nodejs.org/en/download/). Please follow instructions to correctly install.

Once finished, in your Terminal / DOS, run:

```
node -v
npm -v
```

These commands should respond with the versioning of your node and npm without errors.

## Running Gemini-Admin

Running this website requires an ongoing process, so you'll have to first create a new tab/window of your CLI. In this directory, run these commands:

```unix
npm install
npm run start
```

This will run the `gemini-admin` website at <http://localhost:9001>. Open this url with an Internet browser. This website will continue to run until the CLI is closed or you quit the process using `Ctrl-C`.

Gemini consists of two websites, the Admin and User site. Both of these web clients interact with the Orion webserver to provide services. The Admin site allows MathNavigator administrators to manage student/parent information, available programs, and schedule classes and announcements. On the other hand, the User site is the actual site that students and parents will interact with.

For instance, the admin site allows an administrator to schedule new programs and classes for the upcoming semester while a student/parent can register for those classes.

## To Run Tests

```unix
npm run test
```

## To Format Your Code

```unix
npm run prettify
```

## Navigating the codebase

- Take a look at the "scripts" section of `package.json`. You can see all available commands when working on this project.
- This project is managed and bundled by `webpack`. You can view webpack configurations inside `webpack.config.js`.
- The `src/index.js` file contains the `Router` which describes what browser urls will link to which components to display. Most of the folders in the `src` folder are separate pages (i.e. the `src/programs` folder contains different components about the `Programs` page).
- The `src/static.styl` and `/src/app.styl` are CSS styles that can be applied to multiple pages. `app.styl` conveys styles across all pages while the `static.styl` defines variables and constants which can be imported / inherited into other style sheets.
- In general, style sheets and unit tests should be kept within the same component folder so they can easily be referenced via relative pathing.