
# MathNavigator Constellation Services

The back-end infrastructure to Math Navigator is made up of many micro-services, each of which handle smaller, specific tasks for the whole application.

Every part of the entire system will have a codename. I think it's better to give codenames to different parts of the system so that each system can grow in responsibilities in a nebulous, but flexible way. It gives developers the authority into determining how the infrastructure grows and it's just more fun :).

We name our services after star constellations because sailors used to use stars to NAVIGATE through the oceans. GET IT?

------

# Orion
The most famous constellation that can be seen all around the world. This will be our core API service. Most interactions with data (programs, users, class information, etc.) will be done through Orion's webserver.

# Gemini
The twin websites (user and admin sites) will both connect to Orion. The user site will provide services for students and parents while the admin site will allow MathNavigator admins to efficiently update information for the community.

# Aquila
The constellation representing the eagle. Aquila will be our job scheduling service. Over a period of time (days / weeks), Aquila can start tasks that relate to time scheduling, like email reminders. Aquila will often work with AWS's Simple Email Service (SES) and server-less Lambdas.

# Ursa
The constellation that contains the Big Dipper and one of the largest constellations. Ursa is our logging service. Whenever failures happen, Ursa will help us track down where our issues are and provide a good way to store logs on AWS hosts and log them into Amazon ElasticSearch.

------

To learn how to start these services together, view [this](./starting_services.md).