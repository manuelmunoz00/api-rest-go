package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/manuelmunoz00/api-rest-go/models"
)

func homeLink(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", dbURL)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("Connection Established")
		// fmt.Println(db)

	}
	fmt.Fprintf(w, "Welcome")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home page"))
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is admin page"))
}

func adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	r.Get("/", adminPageHandler)
	return r
}

// AdminOnly parses a regular expression and returns, if successful,
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If user is admin, allows access.
		if IsLoggedInAdmin(r) {
			next.ServeHTTP(w, r)
		} else {
			// Otherwise, 403.
			http.Error(w, http.StatusText(403), 403)
			return
		}

		return
	})
}

// IsLoggedInAdmin parses.
func IsLoggedInAdmin(r *http.Request) bool {
	return rand.Float32() < 0.5
}

func migraciones(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(models.Cliente{}, &models.Oferta{}, &models.Servicio{})

	//NuevoCliente
	cliente := models.Cliente{Rut: "16006363-7", Nombres: "Manuel", Apellidos: "Muñoz", Correo: "manuel.munoz00@gmail.com"}
	c := &cliente
	// cliente.NewCliente()
	// fmt.Println(&cliente)
	resultado := c.NewCliente()
	if resultado == true {
		fmt.Println("nuevocliente ", resultado)
	}
	w.Write([]byte("migraciones"))
}

func main() {
	var c models.Cliente
	c.Correo = "correo@gmail.com"
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homeLink)
	// log.Fatal(http.ListenAndServe(":8080", router))
	r := chi.NewRouter()
	r.Get("/", migraciones)
	// r.Get("/admin", adminPageHandler)
	r.Mount("/admin", adminRouter())
	http.ListenAndServe(":3000", r)
}
