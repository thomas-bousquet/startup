
# user-service
[![Tests](https://github.com/thomas-bousquet/user-service/actions/workflows/main.yml/badge.svg)](https://github.com/thomas-bousquet/user-service/actions/workflows/main.yml)
<br>
<br>


This project is an API in charge of user and account management. The postman collection is available at `postman.json`.<br>
For now, tests only include a suite of integration tests that ensures there is no regression over time on the main usecases. 

This project uses MongoDB as a database.

#### Run the application on your local machine

- Run `make start` to start the infrastructure required for the app to be able to start successfully (mongo)
- Go run main.go or use intellij config located at `/.config/intellij.xml` (works with both IntelliJ IDEA Ultimate and Goland)
- If you want to stop, run `make stop`

#### Run tests
- Run `make test`

#### Todo
- Create account / tenant concept
- Add integration tests to prevent regression (update, delete user, list users)