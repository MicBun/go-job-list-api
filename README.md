# go-job-list-api
go job list api that return response from http://dev3.dansmultipro.co.id/api/recruitment/positions.json

### PreInstall
1. Docker https://docs.docker.com/desktop/install/windows-install/

### How to use with Docker
1. Clone this Repository using `git clone`
2. Run Docker Desktop
3. Check config/db.go for the commented out code and adjust accordingly
4. use CMD Command on this project root
    1. `docker-compose build`
    2. `docker-compose up`
5. access http://localhost:8080/swagger/index.html#/
6. The app is ready to use.
---
### How to use with Go Build
1. Clone this Repository using `git clone`
2. Check config/db.go for the commented out code and adjust accordingly
3. use CMD Command on this project root
    1. `go build -tags netgo -ldflags '-s -w' -o app`
    2. `./app`
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.
---
### How to use with Go Run
1. Clone this Repository using `git clone`
2. Check config/db.go for the commented out code and adjust accordingly
3. use CMD Command on this project root
    1. `go run main.go`
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.

---
### Database
1. Create a database named `simple_todo`
2. It will automatically create the tables using gorm
3. Insert username and password to the `users` table
---


For more information, please contact me LinkedIn: https://www.linkedin.com/in/MicBun