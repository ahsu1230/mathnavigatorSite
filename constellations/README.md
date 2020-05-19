
# MathNavigator Constellation Services

The MathNavigator application follows the *microservice* architecture. A microservice architecture splits an application into smaller programs called "services" each of which have a dedicated role and together, they form the entire application.

A good example of microservices could be a music playing application, like Spotify. If Spotify was implemented with a microservice architecture, you could have one service be dedicated to playing music, another service downloading the next song to play, and a third service dedicated to handling the user interface and allowing users to browse music libraries. Spotify is one application but you can think of it as several smaller programs working together, especially if they happen simultaneously (it's common for users to play music and browse for other music at the same time).

Every microservice has a codename. It's often used to give different services codenames so each service may flexibly grow in responsibilities. It allows application infrastructure to grow without being bogged down by naming schematics... and it's just more fun :). Here, we name our services after star constellations because sailors used to use stars to NAVIGATE through the seas. GET IT?

## Onboarding

To learn how to start these services together, view [this](./onboarding.md).

## Codenames

### Orion
The most famous constellation that can be seen all around the world. This will be our core API service. Most interactions with data (programs, users, class information, etc.) will be done through Orion's webserver.

### Gemini
The twin websites (user and admin sites) will both connect to Orion. The user site will provide services for students and parents while the admin site will allow MathNavigator admins to update information for the community.

### Aquila
Aquila is our job scheduling service. Over a period of time (days / weeks), Aquila can start tasks that relate to time scheduling, like email reminders. Aquila will often work with AWS's Simple Email Service (SES) and server-less Lambdas.

### Ursa
Ursa is our logging service. Whenever failures happen, Ursa will help us track down where our issues are and provide a good way to store logs on AWS hosts and log them into Amazon ElasticSearch.
