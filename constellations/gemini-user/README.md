# Gemini User

One of the Gemini website clients. This web client will interact with the Orion webserver by providing MathNavigator services to students and parents. Through this website, they will be able to see available programs and register for classes.

## Pre-requirements:

Make sure that npm is installed before proceeding: (https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/onboarding/install_npm.md)

## Running Gemini-User

Running this website requires an ongoing process, so you'll have to first create a new tab/window of your CLI. In this directory, run these commands:

```unix
npm install
npm run start
```

This will run the `gemini-user` website at <http://localhost:9000>. Open this url with an Internet browser. This website will continue to run until the CLI is closed or you quit the process using `Ctrl-C`.

## Navigating the codebase

Take a look at the "scripts" section of `package.json`. You can see all available commands when working on this project.

In general, style sheets and unit tests should be kept within the same component folder so they can easily be referenced via relative pathing.

## To Run Tests

```unix
npm run test
```

## To Format Your Code

```unix
npm run prettify
```