## GO Server GO!#

v.0.1.0

### Build

To build the project, run the following command:

```bash
$ make build
```

> The main purpose of this project is to learn GO by doing.
> Althoug having a elegant, scalable backend as a result is nice to have :)


### Part I
* `DONE` POC server listening on Port 8080 and running
* `DONE` Serve request, write json to the browser
* `DONE` DB Connection (PostGres)
* `DONE` Instantiate a DB connection
* `DONE` Model User
* `DONE` DB permissions
* `DONE` Model Groups
* `DONE` Query User from DB
* *TODO* CRUD DB Operations
* `DONE` Middlewares? CORS?
* `DONE` Generate TOKEN
* `DONE` Routes via REST

### Before moving forward
* Reassess the current project structure
* Optimze App context availability (Package Main x App its weird)
* A more elegant way to enable a new DB connection from anywhere

### Part II
* `DONE` Implement [GraphQl](https://github.com/graphql-go/graphql)
* `DONE` Link to FEend
* `DONE` Authenticate a request
* `DONE` LogIn | Out | Signup
* *TODO* Protect Routes

* *TODO* Template other DB connectors (BQ, SF)

## NodeJs features to mirror
* *TODO* Password recovery
