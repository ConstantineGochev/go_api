# Go Todo REST API Example
A RESTful API example with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **mgo.v2** (MongoDb Driver)

## Installation & Run
```bash
# Download this project
go get github.com/ConstantineGochev/go_api
```


```bash
# Build and Run
cd go_api
go build
./go_api

# API Endpoint : http://127.0.0.1:3000
```

## Structure
```
├── app
│   ├── app.go
├── controllers         // Our API logic controllers
│   ├── common.go    // Common response functions
│   ├── news.go  // APIs for News model
├── models
│   ├── news.go // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /news
* `GET` : Get all news
* `POST` : Create a new project



## Todo

- [x] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Building a deployment process with Docker 