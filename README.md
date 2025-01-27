Install dependencies
1) gin framework: webframework, simple apis/web apps
2) To install gin, need to setup dependency tracking for our go project
3) go mod init example/GOApi (Terminal) (Creates a go.mod file)
4) go get -u github.com/gin-gonic/gin (Terminal) (Installs gin framework)

Library API
1) Stores Books
2) Checkin, checkout books
3) Add books
4) View books
5) Get book by id

Run the program:
1) go run main.go (Listening and serving HTTP on localhost:8080)
2) curl localhost:8080/books (In a seperate terminal, must be cmd not powershell)

POST REQUEST
1) go run main.go
2) curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"

PATCH REQUEST (Checking out a book)
1) go run main.go
2) curl localhost:8080/checkout?id=2 --request "PATCH"

PATCH REQUEST (Returning a book)
1) go run main.go
2) curl localhost:8080/return?id=2 --request "PATCH"

GET: Getting Something
POST: Adding/Creating
PATCH: Updating