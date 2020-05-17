# Gemini Admin

One of the Gemini website clients. This web client will interact with the Orion webserver by allowing MathNavigator administrators to manage student/parent information, available programs and schedule classes and announcements.

## Onboarding Steps:

Make sure that npm is installed before proceeding: (https://github.com/ahsu1230/mathnavigatorSite/blob/master/resources/onboarding/install_npm.md)

## Running Gemini-Admin

In this directory, run these commands:

```
npm install
npm run start
```

This will run the `gemini-admin` website at <http://localhost:9001>. Open this url with an Internet browser.

Gemini consists of two websites, the Admin and User site. Both of these web clients interact with the Orion webserver to provide 
services. The Admin site allows MathNavigator administrators to manage student/parent information, available programs, and schedule classes and announcements.On the other hand, the User site is the actual site that students and parents will interact with.

For instance, the admin site allows an administrator to schedule new programs and classes for the upcoming semester while a student/parent can register for those classes.