package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"net/http"

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

    router := gin.Default()
    router.GET("/cards", getCards)

    router.Run("localhost:8080")
}

func getCards(c *gin.Context) {
    // If an empty string is used here, you can stil use the API with stricter limits.
    // See: https://docs.pokemontcg.io/#documentationrate_limits
    poke := tcg.NewClient(os.Getenv("PKMN_TCG_API_KEY"))

	// Refer to https://docs.pokemontcg.io/#api_v2cards_list for how queries work
	cards, err := poke.GetCards(
		request.Query("name:jirachi", "types:psychic"),
		request.OrderBy("+name"),
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