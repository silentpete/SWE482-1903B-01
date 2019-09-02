# SWE482-1903B-01

## MySQL Database

Bundy's will use a MySQL Database.

Official MySQL Container: [https://hub.docker.com/_/mysql](https://hub.docker.com/_/mysql)

### Stand up a MySQL Database

```none
docker run -dit --name mysql -e MYSQL_ROOT_PASSWORD=bundys -e MYSQL_DATABASE=bundys -p 3306:3306 -v mysql:/var/lib/mysql mysql:8.0.17 --default-authentication-plugin=mysql_native_password --skip-mysqlx
```

## Bundy's Program

The Bundy's API will be wrote in Go. It will present minimal API endpoints to interact with the database.

To see the program "help", run the following command.

```none
go run bundys.go
```

## References

* Go Homepage - [https://golang.org](https://golang.org)
* Go package database/sql - [https://golang.org/pkg/database/sql/](https://golang.org/pkg/database/sql/)
* Go SQL Database tutorial - [http://go-database-sql.org](http://go-database-sql.org)
SQL - [https://www.w3schools.com/sql/](https://www.w3schools.com/sql/)
* Go CRUD Example - [https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html](https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html)
* Go MySQL Tutorial - [https://tutorialedge.net/golang/golang-mysql-tutorial/](https://tutorialedge.net/golang/golang-mysql-tutorial/)
* Go API Tutorial - [https://tutorialedge.net/golang/creating-restful-api-with-golang/](https://tutorialedge.net/golang/creating-restful-api-with-golang/)
* Go ORM Tutorial - [https://tutorialedge.net/golang/golang-orm-tutorial/](https://tutorialedge.net/golang/golang-orm-tutorial/)
