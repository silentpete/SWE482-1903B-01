# SWE482-1903B-01

## MySQL Database

Bundy's will use a MySQL Database.

Stand up a MySQL Database

MySQL Container: [https://hub.docker.com/_/mysql](https://hub.docker.com/_/mysql)

    ```none
    docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=bundys -e MYSQL_DATABASE=bundys -p 3306:3306 -v mysql:/var/lib/mysql mysql:8.0.17
    ```

## Bundy's Program

To see the program "help", run the following command.

```none
go run bundys.go
```
