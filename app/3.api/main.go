package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
	invController "github.com/pobyzaarif/b14p2/app/3.api/controller/inventory"
	_ "github.com/pobyzaarif/b14p2/app/3.api/docs"
	"github.com/pobyzaarif/b14p2/database"
	invRepo "github.com/pobyzaarif/b14p2/repository/inventory"
	invService "github.com/pobyzaarif/b14p2/service/inventory"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:7070
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	db := database.InitDatabase()

	inventoryRepository := invRepo.NewSQLRepository(db)
	inventoryService := invService.NewService(inventoryRepository)
	inventoryController := invController.NewController(inventoryService)

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"data":{}, "message":"route not found"}`))
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(`{"data":{}, "message":"method not allowed"}`))
	})

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		log.Default().Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": http.StatusText(http.StatusInternalServerError)})
	}

	router.GET("/api/v1/inventory", inventoryController.GetAll)

	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	addr := os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Default().Println("application started on " + addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
