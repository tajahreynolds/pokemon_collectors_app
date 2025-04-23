package main

import (
	"log"
	"os"
	"time"

	"net/http"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tajahreynolds/pokemon_collectors_app/db"
)

type JSONSuccess struct {
	message	string
}

func main() {
	err := godotenv.Load("../../.env")
	if (err != nil) {
	  log.Fatal("Error loading .env file")
	}

	dsn := "user=postgres password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=pokemoncollectors port=5432 sslmode=disable TimeZone=America/New_York"
	if err := db.ConnectPostgres(dsn); err != nil {
		log.Fatal("Failed to connect to the databse:", err)
	}

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
	var requestBody struct {
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user := db.FindUser(requestBody.Email); user != nil {
		// send the magic login link
		c.JSON(http.StatusOK, JSONSuccess{message: "success"});
	} else {
		// user was not found
		c.JSON(http.StatusOK, JSONSuccess{});
	}
}

func register(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user in the database
	newUser := db.User{
		Name:  requestBody.Username,
		Email: requestBody.Email,
	}

	if err := db.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, JSONSuccess{message: "success"})
}