
# user-service
<a href="https://github.com/thomas-bousquet/user-service/actions?query=workflow%3AProjectPipeline">
    <img src="https://github.com/thomas-bousquet/user-service/workflows/ProjectPipeline/badge.svg" />
</a>
<br>
<br>

This project is an API in charge of user and account management. The postman collection is available at `postman.json`.<br>
For now, tests only include a suite of integration tests that ensures there is no regression over time on the main usecases. 

This project uses MongoDB as a database.

#### Run on your local machine

###### Locally with the application itself in the current state
- Run `make start`
- If you want to stop, run `make stop`

###### Locally with the application in the current state through the IDE or in a separate shell
- Run `make start-local`
- Start the application either from a shell (`go run main.go`), or from your IDE (Run or Debug)
  <br>The Goland command is available in .goland, you can import it and run the application without any extra setup
- If you want to stop, run `make stop-local`

###### Locally with the application in a specific version
- Run `make start-docker`. This will pull the image of the application tagged `latest`, you can change the value in `docker-compose.docker.yml` to your will. More information about the image here: https://hub.docker.com/r/thomasbousquet/user-service
- If you want to stop, run `make stop-docker`


#### Run tests
- Run `make start`
- Open another shell and run `make test`. 

#### Todo
- Create account / tenant concept
- Multi tenant security
- Configurable application
- Add integration tests to prevent regression (update, delete user, list users)