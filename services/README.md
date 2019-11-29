# MathNavigator Services

The back-end infrastructure to Math Navigator is made up of many micro-services, each of which handle smaller, specific tasks for the whole application.

Our language of choice is in Go: https://golang.org/

 - API
 - Database
 - Email
 - AWS
 - Jenkins

# Constellation Services
Every part of the entire system will have a codename. I think it's better to give codenames to different parts of the system so that each system can grow in responsibilities in a nebulous, but flexible way. It gives developers the authority into determining how the infrastructure grows and it's just more fun :).

We name our services after star constellations because sailors used to use stars to NAVIGATE through the oceans. GET IT????

# Orion
The most famous constellation that can be seen all around the world. This will be our core API service. The user and admin websites will both funnel through this service in order for users to interact with our database.

# Ursa
The constellation that contains the Big Dipper and one of the largest constellations. Ursa will be in charge of ... TBD!

# Aquila
The constellation representing the eagle. Aquila will be our job scheduling service. Over a period of time (days / weeks), Aquila can start tasks that relate to time scheduling, like email reminders.

# Scorpius
The zodiac representing the scorpion. Here, you will find our Jenkins configurations. Jenkins is a tool for Continuous Integration (CI) which can automatically run tests for us and finds bugs for us, hence a scorpion - a rather big bug.
