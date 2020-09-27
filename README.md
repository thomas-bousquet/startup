# startup
<a href="https://github.com/thomas-bousquet/startup/actions?query=workflow%3AProjectPipeline">
    <img src="https://github.com/thomas-bousquet/startup/workflows/ProjectPipeline/badge.svg" />
</a>
<br>
<br>

This project is an API in charge of user and account management. The postman collection is available at `postman.json`.<br>
For now, tests only include a suite of integration tests that ensure there is no regression over time on main usecases. 

This project uses MongoDB as a database.

#### Run project on your local machine

###### Development
- First run `make start-dev`
- Then start the application either from a shell (`go run main.go`), or from your IDE (Run or Debug)
- If you want to stop, run `make stop-dev`

###### Production like
- First run `make start`
- If you want to stop, run `make stop`

#### Run tests
- First run `make start`
- Then open another shell and run `make test`. 

#### Todo
- [start]: Create account / tenant concept
- [start]: Multi tenant security
- [start]: Log system
- [start]: Configurable application
- [improve]: Wrap errors if possible
- [start]: Add integration tests to prevent regression