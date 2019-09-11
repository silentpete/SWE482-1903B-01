# SWE482-1903B-01

For this class we are implementing a shoes sales system. We will all need to contribute to the Project Plan,

## Table of Contents

- [Production Environment](#production-environment)
  - [Pre-Reqs](#pre-reqs)
- [Development Environment](#development-environment)
- [Bundy's Program](#bundys-program)
- [MySQL Database](#mysql-database)
  - [Stand up a MySQL Database](#stand-up-a-mysql-database)
  - [Log into MySQL Container](#log-into-mysql-container)
  - [MySQL Common Commands](#mysql-common-commands)
- [References](#references)

## Production Environment

To stand up this environment, please fulfill the pre-reqs before running the first_run.sh script. This can be run locally in a development environemtn as well.

### Pre-Reqs

To start the Bundy's environment, there are some expectations.

- CentOS 7
- SSH Access
- System has access to the internet (github.com and docker.com)
- Account has ability to install packages on the system (for now, just run as root)

Once the expectations are fulfilled, the `first_run.sh` script can be run to start the stack.

## Development Environment

The minimal environment needed to interact with the Bundy's Program is having Go installed. You can find download and installation directions here: [https://golang.org/dl/](https://golang.org/dl/)

The full working development environment would need Go, Linux or Windows with Docker.

## Bundy's Program

The Bundy's API will be wrote in Go. It will present minimal API endpoints to interact with the database.

### Bundy's Help

To see the program "help", run the following command.

```none
go run bundys.go --help
```

## MySQL Database

Bundy's will interact with a MySQL Database.

Official MySQL Container: [https://hub.docker.com/_/mysql](https://hub.docker.com/_/mysql)

### Stand up a MySQL Database

```none
docker run -dit --name mysql -e MYSQL_ROOT_PASSWORD=bundys -e MYSQL_DATABASE=bundys -p 3306:3306 -v mysql:/var/lib/mysql mysql:8.0.17 --default-authentication-plugin=mysql_native_password --skip-mysqlx
```

### Log into MySQL Container

```none
docker exec -it mysql bash
```

### MySQL Common Commands

Login to MySQL database

```none
mysql --user=root --password=bundys --database=bundys
```

```none
select * from shoes;
```

```none
INSERT INTO shoes (brand,model,color,size,price,stock) VALUES ("nike", "air force one", "white", 9, 99.99, 0);
```

## References

- Go Homepage - [https://golang.org](https://golang.org)
- Go package database/sql - [https://golang.org/pkg/database/sql/](https://golang.org/pkg/database/sql/)
- Go SQL Database tutorial - [http://go-database-sql.org](http://go-database-sql.org)
SQL - [https://www.w3schools.com/sql/](https://www.w3schools.com/sql/)
- Go CRUD Example - [https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html](https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html)
- Go MySQL Tutorial - [https://tutorialedge.net/golang/golang-mysql-tutorial/](https://tutorialedge.net/golang/golang-mysql-tutorial/)
- Go API Tutorial - [https://tutorialedge.net/golang/creating-restful-api-with-golang/](https://tutorialedge.net/golang/creating-restful-api-with-golang/)
- Go ORM Tutorial - [https://tutorialedge.net/golang/golang-orm-tutorial/](https://tutorialedge.net/golang/golang-orm-tutorial/)
- Robson, E., & Freeman, E. (2012). Head first Html and Css (2nd ed.). Sebastopol, CA: OReilly.
- W3schools.com. (n.d.). How TO - Sort a Table. Retrieved from https://www.w3schools.com/howto/howto_js_sort_table.asp
- Tutorialspoint.com. (n.d.). JSON with Ajax. Retrieved from https://www.tutorialspoint.com/json/json_ajax_example.htm