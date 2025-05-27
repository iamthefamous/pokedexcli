package pokeapi

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type (
	LocationAreaDetails struct {
		PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
	}
	PokemonEncounters struct {
		Pokemon Pokemon `json:"pokemon"`
	}

	Pokemon struct {
		Name           string `json:"name"`
		Height         int    `json:"height"`
		Weight         int    `json:"weight"`
		BaseExperience int    `json:"base_experience"`

		Stats []struct {
			BaseStat int `json:"base_stat"`
			Stat     struct {
				Name string `json:"name"`
			} `json:"stat"`
		} `json:"stats"`

		Types []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
	}
)
