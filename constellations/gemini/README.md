# Gemini

Gemini consists of two websites, the Admin and User site. Both of these web clients interact with the Orion webserver to provide services. The Admin site allows MathNavigator administrators to manage student/parent information, available programs, and schedule classes and announcements.On the other hand, the User site is the actual site that students and parents will interact with.

For instance, the admin site allows an administrator to schedule new programs and classes for the upcoming semester while a student/parent can register for those classes.

## Onboarding Steps:

Make sure that npm is installed before proceeding: (https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/onboarding/install_npm.md)

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
