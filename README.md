# Pokédex CLI

A command-line Pokédex built in Go that lets you explore the Pokémon world using real data from the [PokeAPI](https://pokeapi.co/)!  
Catch Pokémon, explore location areas, inspect stats, and more — all from your terminal.

---

## 🚀 Features

- 🌍 `map`: View paginated Pokémon location areas (20 per page)
- 🔁 `mapb`: Go back to the previous 20 location areas
- 🕵️ `explore <location_area>`: List Pokémon found in a specific location area
- 🧢 `catch <pokemon>`: Attempt to catch a Pokémon based on its base experience
- 📘 `pokedex`: View your caught Pokémon
- 🔍 `inspect <pokemon>`: See stats and type info for a caught Pokémon
- ⚡ Fast with built-in caching using `internal/pokecache`

---

## 🧰 Requirements

- Go 1.18+
- Internet connection (for live API requests)

---

## 🧪 Installation

```bash
git clone https://github.com/iamthefamous/pokedexcli.git
cd pokedexcli
go build -o pokedex
./pokedex
