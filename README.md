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
- [x] POC server listening on Port 8080 and running
- [x] Serve request, write json to the browser
- [x] DB Connection (PostGres)
- [x] Instantiate a DB connection
- [x] Model User
- [x] DB permissions
- [x] Model Groups
- [x] Query User from DB
- [ ] CRUD DB Operations
- [x] Middlewares? CORS?
- [x] Generate TOKEN
- [x] Routes via REST

### Before moving forward
* Reassess the current project structure
* Optimze App context availability (Package Main x App its weird)
* A more elegant way to enable a new DB connection from anywhere

### Part II
- [x]  Implement [GraphQl](https://github.com/graphql-go/graphql)
- [x]  Link to FEend
- [x]  LogIn | Out | Signup
- [ ]  Authenticate a request
- [ ]  Session Management
- [ ] Protect Routes


* *TODO* Template other DB connectors (BQ, SF)

## NodeJs features to mirror
* *TODO* Password recovery
