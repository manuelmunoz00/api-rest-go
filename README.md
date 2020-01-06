cp .env.example .env

go get -v -u github.com/go-chi/chi
go get -v -u github.com/jinzhu/gorm
go get -v -u github.com/joho/godotenv
go get -u github.com/go-sql-driver/mysql

go install github.com/manuelmunoz00/api-rest-go/models
go install github.com/manuelmunoz00/api-rest-go/config