# Investments portfolio API

![Go: 1.14](https://img.shields.io/badge/Go-1.14-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue)](https://opensource.org/licenses/MIT)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-no-red.svg)](https://github.com/WeNeedThePoh/investment-api/graphs/commit-activity)

***Live URL:*** https://investments-api.herokuapp.com  

***Tech stack***: GoLang, Postgres, OpenAPI, Heroku

***DISCLAIMER***: This was a small project to start learning and working with GoLang. So every code you see here was the first lines of code that I wrote in GoLang so be mindful of this.

This is an REST API to track your stock investments. With this API you can have your portfolio, with all the transactions history, current stocks, past stocks, market values, percentages, profit, loss, etc. The idea of this API was actually to be used aftewards on a webapp. This ideia got dumped for now but I might catch it later.

## Docs

For the documentation we used the OpenAPI specification. We have all the available endpoints with the schemas and examples for each use case. It's not just because is an industry standard to use it but also because is really easy to update and to read. There is also a lot of tools to convert the yaml file to a beautiful HTML.

Regarding the database we use the Database markup language -- DBML for short. Again it's really amazing and can put any new joiner into speed right away. There is also an online tool to visualize the table schemas.

## Third-party data

For the stock data, we used [alpha advantage API](https://www.alphavantage.co). It has a lot of cool feautres and the free rate is also very generous.

## Deployments (CI/CD)

For now this is not very too complex. We have heroku integrated with Github, so every time code is pushed to main branch it's going to trigger a new build to production. This means that we have an actual true Continuous integration/ Continuous deployment. This way we always make sure that every commit is good to run on production before we ship it.
## Development

For now it's the easiest setup, so we start our app using CLI.
This is going to install any dependencies, start the app and watch for new changes.
```
make start
```

In the future the best way is just to use a docker container.

For the migrations, we are using [migrate](https://github.com/golang-migrate/migrate). It's easy to use and we have a lot of choices in terms of how to use it.
To create a new migration hit the following command:
```
make migration name=MyNewMigration
```

To run new migrations hit the following command:
```
make migrate_up
```

To revert the last migration hit the following command:
```
make migrate_down
```

## Tests

At the moment we only have unit tests for the service layer, but we can and should add more unit tests for the remaing layers. Also add integration tests using a docker container to deploy a temporary database.

To run tests

```
make test
```

## License

MIT Licensed (file [LICENSE](LICENSE)).
