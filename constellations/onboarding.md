# General Onboarding

The MathNavigator repository follows a microservice architecture. This means that the application is made up of many "services", each service haivng a particular role. And together, the services make up the entire application.

We use a service called **Docker** to manage our microservices. Please refer to the Docker resources [here](https://github.com/ahsu1230/mathnavigatorSite/tree/master/resources/docker) to learn how to use Docker. You should install Docker onto your computer (DockerDesktop also recommended) and learn the basics of Containerization. Once Docker is installed, follow the below instructions.

## Docker-Compose

In the `constellations` folder, use `docker-compose` to manage your services. Run this command to start all services.

```unix
docker-compose up -d
```

Once it is done running, you should be able to see all services as healthy in the Docker Desktop app.

![SCREENSHOT_DOCKER_ALL_SERVICES](https://github.com/ahsu1230/mathnavigatorSite/tree/master/resources/images/docker/desktop_all_services.png)

***Note*** Even though our stack uses MySQL, you won't have to install it because Docker will already create a MySQL image for you (as done in `docker-compose.yml`).

When you are finished working with these services, you can use this command to deactive all of them:

```unix
docker-compose rm
```

## Starting Gemini services

In `constellations/gemini-user` and `constellations/gemini-admin`, follow the instructions to install Node and NPM. Then start both web-clients by running `npm run start` in both projects.

At this point, you should have backend-services active on DockerDesktop and two websites running.

## Starting Development

If you're a back-end developer, you will probably be working on `orion`. Go to the `orion` directory and read the README file there. You will also need to install Golang to start developing.

If you're a front-end developer, you won't need any more installation steps. Go to the `gemini-user` or `gemini-admin` directories to read more about developing in those projects.

---

For more information about using `docker-compose`, go [here](./onboarding_docker-compose.md).