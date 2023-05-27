<div style="width:100%; display: flex; align-items: center;">
  <h1>Api Go REST
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" height="60" width="70" style="margin-bottom: -15px; z-index: -10; margin-left: 1.25rem"/>
  </h1> 
</div>



### 🛠  Description   

</br>

This project aims to create a simple REST API using the Go programming language (Golang). It leverages various Go libraries to handle routing (Gorilla Mux), logging (httpsnoop), CORS middleware (Gorilla handlers), and database interactions (pgx, gorm). The API connects to a PostgreSQL database and uses GORM (Go Object-Relational Mapper) for database operations..


## Preview 

</br>

<p align="center">
  <kbd>
 <img width="800" style="border-radius: 10px" height="400" src="https://github.com/JuanCampbsi/Preview_README/blob/231351c52b68b6a15709ce89888a400552c44fc3/assets/api_rest_go.gif" alt="Intro"> 
  </kbd>
  </br>
</p>

</br>

### ⌨ Database configuration
Create an .env file in the root of the project and define your database user and password in it and place the file in docker-compose.yml
For example:

```bash
# .env
 DATABASE_USER=use
 DATABASE_PASSWORD=password

# use in docker-compose.yml 
 PGADMIN_DEFAULT_EMAIL: ${DATABASE_USER}
 PGADMIN_DEFAULT_PASSWORD: ${DATABASE_PASSWORD}   

```

### ⌨ Installation
To use it, you need to clone the repository, install the dependencies and run the project.

```bash
# Open terminal/cmd and then Clone this repository
$ git clone https://github.com/JuanCampbsi/api_go_rest.git

# Access project folder in terminal/cmd
$ cd api_go_rest

# Install the dependencies
$ go mod tidy

# Run the application in development mode
$ go run main.go

# In order to create a container for this application, you'll need to run a specific Docker command. Please ensure that Docker is already installed on your machine before proceeding. If you don't have Docker installed, you can download it from the [official website](https://www.docker.com/products/docker-desktop).

$ docker compose up                                 

```

</br>	

### ⌨ Stack of technologies and libraries

-   [Golang](https://go.dev/doc/) - version 1.20
-   [Gorilla Mux](https://github.com/gorilla/mux/) - version 1.8.0
-   [Gorilla handlers](https://github.com/gorilla/handlers/) - version 1.5.1
-   [GORM](https://gorm.io/gorm ) - version 1.25.1
-   [PostgreSQL](https://www.postgresql.org/download/) - version 10.22 
 
</br>

👨‍💻 **Author** 💻

Developed by [_Juan Campos_](https://www.linkedin.com/in/juancampos-ferreira/)