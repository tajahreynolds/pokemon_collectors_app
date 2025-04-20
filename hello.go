package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name 	  string
}

type JSONSuccess struct {
	message	string
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=postgres password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=pokemoncollectors port=5432 sslmode=disable TimeZone=America/New_York",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
	  log.Fatal("Error connecting to database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

    engine := gin.Default()
	engine.SetTrustedProxies(nil)
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	engine.Use(cors.New(cors.Config{
	  AllowOrigins:     []string{"http://localhost:5173"},
	  AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
	  AllowHeaders:     []string{"Origin"},
	  ExposeHeaders:    []string{"Content-Length"},
	//   AllowCredentials: true,
	//   AllowOriginFunc: func(origin string) bool {
	// 	return origin == "https://github.com"
	//   },
	  MaxAge: 12 * time.Hour,
	}))

	// AUTH METHODS
    engine.POST("/login", login)
    engine.POST("/register", register)
    // PROFILE METHODS

	// CARDS METHODS
	engine.GET("/cards", getCards)
    engine.POST("/search", search)
	// COLLECTIONS METHODS

	// engine.NoRoute(func(c *gin.Context) {
	// 	// todo this should point to index.html in public dir
	// 	// c.File("./vue-frontend/index.html")
	// })
    engine.Run("localhost:8080")
}

func getCards(c *gin.Context) {
    // If an empty string is used here, you can stil use the API with stricter limits.
    // See: https://docs.pokemontcg.io/#documentationrate_limits
    poke := tcg.NewClient(os.Getenv("PKMN_TCG_API_KEY"))

	// Refer to https://docs.pokemontcg.io/#api_v2cards_list for how queries work
	cards, err := poke.GetCards(
		request.Query("name:jirachi", "types:psychic"),
		request.OrderBy("+number"),
		request.PageSize(3),
		request.Page(2),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, card := range cards {
		log.Printf("%s: %s\n", card.Name, card.Set.Name)
	}
	c.JSON(http.StatusOK, cards);
}

func search(c *gin.Context) {
	var requestBody struct {
		Name string `json:"name"`
	}
    // If an empty string is used here, you can stil use the API with stricter limits.
    // See: https://docs.pokemontcg.io/#documentationrate_limits
    poke := tcg.NewClient(os.Getenv("PKMN_TCG_API_KEY"))
	
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cards, err := poke.GetCards(request.Query("name:" + requestBody.Name))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, cards)
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, "login success: true");
}

func register(c *gin.Context) {
	c.JSON(http.StatusOK, "register success: true");
}