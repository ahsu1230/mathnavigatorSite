# Gemini

The Admin site is one of the main web clients that interact with Orion. This site allows MathNavigator administrators to manage student/parent information, available programs, and schedule classes and announcements.

## Onboarding Steps:

By the end of these steps, you should be able to run a local webserver to both host the back-end golang servers and a front-end web application.

## Run the front-end web application (For both Admin and User)

Inside the respective directory (admin or user), run these commands:

```
npm install
npm run start
```

This will run another webserver (this time on port 8081) which will host this website.
Now, if you go to http://localhost:8081 in an Internet browser, you should be able to see the Admin/User website!

## Run tests on the front-end web application (For both Admin and User)

Inside the respective directory (admin or user), run:

```
npm run test
```

You should see all successes and no failures!
