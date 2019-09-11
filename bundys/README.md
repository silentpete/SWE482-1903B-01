# bundys

## Files

The `.dockerignore` file is used to list files we do not want git to track, this helps not commit things to the repository we don't want.

The `build.sh` file is used to build the Dockerfile.

The `bundys.go` file is wrote in Googles Go language. This language can be compiled into a portable binary file.

The `Dockerfile` is used as the build file for the Bundys container. This process puts the compiled binary in a static predictable environment.

The `go.mod` and `go.sum` files are used for dependency versioning.

The `invnetory.json` file is used to load the database with default values.

## Directories

`api`

The API directory will house the basic CRUD functions.

`server`

The Server directory will house the Bundys server functions.
