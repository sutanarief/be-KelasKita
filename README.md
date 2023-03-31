# KelasKita
<!--- These are examples. See https://shields.io for others or to customize this set of shields. You might want to include dependencies, project status and licence info here --->
![GitHub repo size](https://img.shields.io/github/repo-size/sutanarief/be-KelasKita?style=plastic)
![Github languange](https://img.shields.io/github/languages/top/sutanarief/be-KelasKita?logo=go&style=plastic)
![GitHub contributors](https://img.shields.io/github/contributors/sutanarief/be-KelasKita?style=plastic)
![GitHub stars](https://img.shields.io/github/stars/sutanarief/be-KelasKita?style=social)
![GitHub forks](https://img.shields.io/github/forks/sutanarief/be-KelasKita?style=social)

KelasKita is an application for teachers and students to discuss, which provides a feature to post answers to any questions that have been made by other users.


# Prerequisites

Before you begin, ensure you have met the following requirements:
* Go version 1.20.2
* PostgreSQL

# Using KelasKita
Add .env file with these variables
```env
DATABASE_URL=postgresql://${{ PGUSER }}:${{ PGPASSWORD }}@${{ PGHOST }}:${{ PGPORT }}/${{ PGDATABASE }}

PGDATABASE=<databasename>
PGHOST=<host>
PGPASSWORD=<pg user password>
PGPORT=<database port>
PGUSER=<pg username>
JWTKEY=<jwt secret>
PORT=<localhost port>
```

Installing dependencies :
```
go mod tidy
```
Run application :
```
go run main.go
```
Run application with nodemon (*if you have installed nodemon*) :
```
nodemon --exec go run main.go --signal SIGTERM
```

# API Documentation
See all API Documentation on [KelasKita Postman Documenter](https://documenter.getpostman.com/view/14405021/2s93RTSsve)