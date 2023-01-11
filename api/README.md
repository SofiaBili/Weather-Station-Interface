# CRM-Customer-Details-API

In this API, it is concluded the services and the data access objects for the CRM project.

It is implemented in Go with Gin as the main framework and gorp as the ORM library.

## Requirements

- Go v1.18.1
- MySQL v15.1/MariaDB 10.9.4

## How to run

### Unix

You will need to refer to your OS's package manager to how to install them.

For the sake of the guide, there will be a demonstration to how to install on ubuntu-based distros.

```bash
sudo apt install mariadb-server golang-go
sudo mysql_secure_installation
go get .
```

Once everything is ready, then run `go run .` to launch the API.

Hit the endpoints with: `localhost:3000`.

### Windows

To install on Windows follow the guides to how to install [MySQL](https://dev.mysql.com/downloads/installer/) and [Golang](https://go.dev/doc/install).

Install the packages with `go get .` and once everything is ready, run `go run .` to launch the API.

Hit the endpoints with `localhost:3000`.
