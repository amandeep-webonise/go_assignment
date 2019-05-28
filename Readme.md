# webogo
Boilerplate for golang apps

## Ecosystem
On a local development environment, the Go web server will talk to the following tools (Please note that this list will evolve as we keep on building the tool out)
1. PostgreSQL 10
2. [sql-migrate](https://github.com/rubenv/sql-migrate) for database migrations.
3. [xo](https://github.com/xo/xo) for generating database access objects. It will be replaced by `ixo` which generates interface driven stubs for the same layer. Makes it easy to mock objects out.
4. [dep](github.com/golang/dep/cmd/dep) for vendoring application dependencies.
5. [packr](github.com/gobuffalo/packr) for bundling migrations, views and js/css files.
6. [envparser](github.com/caarlos0/env) for parsing environment variables.

## Setting up Pre-Requisites

### Setting up IDE
We have couple of options for open source IDE, install one of the following with their respective golang extension:
1. [visual-studio-code](https://code.visualstudio.com/Download) IDE with extension: Go (Rich Go language support for Visual Studio Code By Microsoft)
2. [atom](https://atom.io/) IDE with extension: [go-plus](https://atom.io/packages/go-plus)

### Installing Go  
1. Install Go by following the instructions mentioned [here](https://golang.org/doc/install)  
2. Please ensure to setup the `GOPATH` env var correctly.
3. Webogo is built on latest go binary release [here](https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz)

### Cloning the code.
A pre-requisite reading is highly recommended about how to write Go code - https://golang.org/doc/code.html.
For the first time setup the following commands have to be run assuming $GOPATH gives a different output than $GOROOT:  
```
cd $GOPATH
mkdir src bin pkg
cd src
mkdir github.com
cd github.com
mkdir webonise
```
The above commands will setup the basic directory structure.Now you are all set to clone the code

The location to clone the code is -`$GOPATH/src/github.com/webonise/`. cd to that location and clone this repository

### Setting environment variables
For ease of environment interoperatibility, we have chosen to store the configuration parameters in environment variables.Currently these are the following env vars with their descriptions

| Env Var       | Description                 |
| ------------- | --------------------------- |
| DB_USERNAME       | Database username       |
| DB_NAME    | Database name  |
| DB_PASSWORD   | Database password       |
| DB_HOST       | Dataabase hostname       |
| DB_PARAMS   | Database connection parameters           |
| PORT          | Port to run go web server on              |
| DBPORT | Database port on which postgres is running |
| ENV | Environment |


The `setenv.sh.example` has sensible defaults for the local development environment.
Follow the commands to set those vars on your local machine:
Also if you end up adding a new env var please add the sensible default to the `setenv.sh.example` file

For parsing environment variables we are using [env](https://github.com/caarlos0/env). The code is modified to use the `envprovider` package which has an interface to take in configuration which are assumed to be the environment variables' keys, bind them to their appropriate environment tags and hydrate the `config.ServerConfig` object.  
The only thing to follow as a rule of thumb when a new env var is added:
1. Set the sensible default in the `setenv.sh.example` file.  
2. Add the var to `config.ServerConfig` struct with env tag.

```
cp setenv.sh.example ./setenv.sh
chmod +x ./setenv.sh
source ./setenv.sh
```

At this point you are ready to run the application locally.  


## Quickstart
1. Clone the repository
2. Set the environment variables using your copy of `setenv.sh`
3. Run `go generate` to execute the build pipeline.
4. Run the app using `go run cmd/srv/main.go`
5. Your application will come up at `localhost:9999`

Till the time folks get hang of the usual go development lifecycle, it is recommended that they run `go generate` everytime they restart the server.


## A note on Development Lifecyle  

The development lifecycle philosophy is as follows:  
1. Write code which includes code and migrations.  
2. Mutate the database schema as per the current requirement using `sql-migrate`.  
3. Reflect on the database schema and generate the database access objects using `xo`. Please check in those go files and not modify them by hand.  
4. Run the application and test.  


## Deep Dive

### Architecture Overview
We have develop an architecture which is interface driven and [clean](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

#### Rules
1. Follow naming conventions strictly [here](https://talks.golang.org/2014/names.slide#1)
2. Source code dependencies can only point inwards. Nothing in an inner circle can know anything at all about something in an outer circle. This is achieve with use of interfaces.
3. A package contains dependency structure on which functions are defined on. Exported functions are accessible through an interface only.
4. Outer layers has no accessibility to inner layers's dependency structure.
5. No library code is exposed directly to application source code. Instead, library code is put in separate package under pkg/ and interface is exposed to use.
6. Use many small interfaces to model dependencies.

#### Flow
An architectural design pattern divides application into interconnected layers.
1. Starting point of go web server is ``` main.go ``` which initialises server container structure.
2. Container structure consists of dependencies on which web server rely on.
3. Server container initialises dependency structures of application layers.
4. Router can access set of defined middlewares and controllers through interfaces for maintaining routes.
5. Middleware depends upon pkg/monitoring and has access to services. We require our middleware to record and report stacktrace in case our code panics!. Services are used by middlewares to verify user session using cookie and recording of routes for audit purposes.
6. Controllers has access to list of services and call application specific use cases through ``` ServiceProvider ``` interface.
7. Services has access to list of all models methods as their interfaces are part of base service.
8. Models are auto generated considering database schema through xo and used by services.

#### Overview of packages

##### app/containers
1. A golang application can have multiple sub-applications. Their respective containers will go here.
2. A server container has following dependency structure -
```
type Server struct {
	Router    *router.Multiplexer
	Log       logger.Ilogger
	TplParser templates.ITemplateParser
	Cfg       *configs.ServerConfig
	DB        *sql.DB
}
```

##### app/routers
1. An application can have multiple types of routes. It can serve REST APIs and gRPCs as well. Those routes will go here.
2. A server router has following dependency structure - 
```
type SrvRouter struct {
	Router     router.Router
	Middleware middlewares.SrvAuthenticator
	Controller controllers.BaseController
}
```

##### app/controllers
1. Application controllers will go here. 
2. Add your controllers to ``` BaseController ``` interface otherwise they won't be accessible to router. 
3. A server base controller has following dependency structure and interface - 
```
type Srv struct {
	Log       logger.Ilogger
	TplParser templates.ITemplateParser
	Service   services.ServiceProvider
}

type BaseController interface {
	Ping(w *framework.Response, r *framework.Request)
}
```

##### app/middlewares
1. Single application can have different types of authenticators for set of routes defined.
2. Add your type of authenticators for server to ``` SrvAuthenticator ``` interface otherwise they won't be accessible to router.
3. A server middleware has following dependency structure and interface - 
```
// BaseMiddleware contains server middleware configurations
type BaseMiddleware struct {
	Service services.ServiceProvider
	Notify monitoring.PanicNotifier
}

// SrvAuthenticator provides server middleware methods
type SrvAuthenticator interface {
	Handle(handler func(*framework.Response, *framework.Request)) http.HandlerFunc
	RenderView(viewHandler func(*framework.Response, *framework.Request)) http.HandlerFunc
}
```

##### app/services
1. Controllers can access to set of services defined through their interface.
2. Register new service by defining function on type ``` Service ``` and encapsulate in interface.
3. A server base service has following dependency structure and interface - 
```
// Service contains basic dependencies on which services depends
type Service struct {
	Log logger.Ilogger
	User models.UserService
}

// ServiceProvider provides services to controllers
type ServiceProvider interface {
	FetchAllUsers() ([]*models.User, error)
}
```

##### app/models
They are explained in detail below.
To use models in your services - 
1. Make model interface part of base service structure

```
type Service struct {
	Log logger.Ilogger
	User models.UserService
}

```
2. Initialise model structure with database object in server container

```
&services.Service{
	User: &models.UserServiceImpl{DB: s.DB},
	Log:  s.Log,
}

```

### Installing Dependencies 
For dependencies we are using [dep](https://github.com/golang/dep) which is dependency management tool for Go. 
It requires Go 1.9 or newer to compile. [docs](https://golang.github.io/dep)

Initialize your project:
1. ``` dep init ```

Note: 
1. This command needs to use only once.
2. It will produce two files: Gopkg.toml Gopkg.lock.
3. Application dependencies that currently used by project will be place in vendor/


The only thing to follow as a rule of thumb when a new dependency is added:
1. ``` dep ensure -add dependency_name  ```
2. ``` dep ensure ```

#### Database migrations:
1. First you need to create dbconfig.yml.yml file for database configurations. You can refer dbconfig.yml.example for reference.

2. To add a migration 
`sql-migrate new some_migration`
This command will create a file in the `./app/migrations/` folder in which one can write vanilla PostgreSQL queries for DDL or DML.

3. To run migrations
`sql-migarate up`


#### Models/DAO generation:
1. To generate the data access objects  
`xo pgsql://local:local@localhost/webogo?sslmode=disable -o internal/models --suffix=.go --template-path templates/`

2. Please note that the models that are generated in the path denoted by the `-o` flag are to be checked in.  
***please do not modify these file by hand. Let the `xo` tool or the `go generate` command do it for you.***


### Switching between databases
We have provided easier way to switch between databases. 
For now it supports two databases mysql and postgres. 
For switching database you need to follow below step:

1. For postgres you need to add implementation for database configuration in main.go :
 ```&database.PGSQLDBConnectionInitialiser{}```
2. For mysql you need to add implementation for database configuration in main.go :
```&database.MySQLDBConnectionInitialiser{}```


### Bundling
For bundling we are using [packr](https://github.com/gobuffalo/packr) which bundles the js, css, html and migrations files into go binary.

## Syntactic Sugar
As a syntactic sugar to abstract away the need of remembering these clunky commands,a file `build_pipeline.go` is created. It is based on [go generate](https://blog.golang.org/generate).  
It assumes that `DBHOST` `DBUSERNAME` `DBPASSWORD` `DBCONNPARAMS` and `DBNAME` are set as environment variables.  

It does the following things:
1. Installs the required dependencies.
2. Runs the migrations script in the `migrations` folder against the database pointed to by the env vars.  
3. Reflects on the same and generates the DAO in the `internal/models` folder.

So the build pipeline essentially now becomes:
`go generate`


## Note about the xo models

The xo default templates have been modified so as to not let the methods be defined on the package scope.  
Currently the generator generates an interface for each of the type as follows:

```
type XXXService interface {
	DoesXXXExists(pur *XXX) (bool, error)
	InsertXXX(pur *XXX) error
	UpdateXXX(pur *XXX) error
	UpsertXXX(pur *XXX) error
	DeleteXXX(pur *XXX) error
	GetAllXXXs() ([]*XXX, error)
	GetChunkedXXXs(limit int, offset int) ([]*XXX, error)
}
```  
It also generates a struct like this which implements the above methods

```
type XXXServiceImpl struct {
	DB XODB
}
```

Currently the code base requires the XXXServiceImpl to be defined at service level as follows:
```
a := &models.XXXServiceImpl{DB: dbConn}
```  

The advantages are obvious here.  
1. We get nicely encapsulated services which can be mocked out during testing.  
2. No global scope.

Currently the xo templates generate only the CRUD methods on a Service interface.
The methods generated based on the indexes are still global. WIP to modify the templates to have the index methods to be defined on the ServiceInterface as well.


## Extending the generated service interfaces.

The models & types generated by xo should not be changed by hand. This section will define how to extend the existing set of interfaces by using the concept of embedding.