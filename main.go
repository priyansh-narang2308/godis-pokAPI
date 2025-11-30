package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// these are the variables
var ctx = context.Background()
var redisClient *redis.Client

// this starts the redis server
func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// this defines the charactertis of the pokemon
type Pokemon struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	XP    int    `json:"xp"`
	Power string `json:"power"`
	Level int    `json:"level"`
}

func getPokemonByType(pokemonType string) ([]Pokemon, error) {
	// pick every part of the pokemon they have thats why *
	keys, err := redisClient.Keys(ctx, fmt.Sprintf("pokemon:%s:*", pokemonType)).Result()
	if err != nil {
		return nil, err
	}

	// the name is the key so for every key add it
	var pokemons []Pokemon
	for _, key := range keys {
		data, err := redisClient.Get(ctx, key).Result() //get every pokemon of the type
		if err != nil {
			return nil, err
		}

		var pok Pokemon
		// decodes json data to go struct, data converted to byte
		if err := json.Unmarshal([]byte(data), &pok); err != nil {
			return nil, err
		}
		pokemons = append(pokemons, pok)

	}

	return pokemons, nil
}

func handlePokemonType(w http.ResponseWriter, r *http.Request, pokemonType string) {
	pokemons, err := getPokemonByType(pokemonType)

	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	jsonResponse(w, pokemons)
}

// this is the json resp Encoder
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/water", func(w http.ResponseWriter, r *http.Request) {
		handlePokemonType(w, r, "water")
	})
	http.HandleFunc("/electric", func(w http.ResponseWriter, r *http.Request) {
		handlePokemonType(w, r, "electric")
	})
	http.HandleFunc("/grass", func(w http.ResponseWriter, r *http.Request) {
		handlePokemonType(w, r, "grass")
	})
	http.HandleFunc("/legendary", func(w http.ResponseWriter, r *http.Request) {
		handlePokemonType(w, r, "legendary")
	})
	http.HandleFunc("/fire", func(w http.ResponseWriter, r *http.Request) {
		handlePokemonType(w, r, "fire")
	})

	fmt.Println("Starting Pokemon API server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
