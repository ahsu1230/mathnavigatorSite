# Gemini User

One of the Gemini website clients. This web client will interact with the Orion webserver by providing MathNavigator services to students and parents. Through this website, they will be able to see available programs and register for classes.

## Install Node & NPM

Make sure that npm is installed before proceeding. Install [NodeJs](https://nodejs.org/en/download/). Please follow instructions to correctly install.

Once finished, in your Terminal / DOS, run:

```
node -v
npm -v
```

These commands should respond with the versioning of your node and npm without errors.

## Running Gemini-User

Running this website requires an ongoing process, so you'll have to first create a new tab/window of your CLI. In this directory, run these commands:

```unix
npm install
npm run start
```

This will run the `gemini-user` website at <http://localhost:9000>. Open this url with an Internet browser. This website will continue to run until the CLI is closed or you quit the process using `Ctrl-C`.

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
- In general, all "page" components can be found in the `src` folder. For every component, please keep related style sheets and unit tests within the same folder so they can easily be referenced with relative pathing.
- The React `Router` can be found in `src/app/index.js`. There you will see all acceptable paths to this website and which component each page maps to.
- All "general-page" styles can be found in `src/utils/_constants.sass`. It contains many standardized values which you should use for your styling.