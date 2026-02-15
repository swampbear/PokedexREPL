# Pokedex CLI

A command-line Pokedex built in Go as part of the [Boot.dev](https://boot.dev) guided project. Explore locations, catch Pokemon, and inspect your collection — all from the terminal.

## Build & Run

```sh
go build -o pokedexcli && ./pokedexcli
```

## Commands

| Command              | Description                          |
|----------------------|--------------------------------------|
| `help`               | Display available commands           |
| `map`                | List next page of location areas     |
| `bmap`               | List previous page of location areas |
| `explore <area>`     | List Pokemon in a location area      |
| `catch <pokemon>`    | Attempt to catch a Pokemon           |
| `inspect <pokemon>`  | View stats of a caught Pokemon       |
| `pokedex`            | List all caught Pokemon              |
| `exit`               | Exit the program                     |

## Details

- Uses the [PokeAPI](https://pokeapi.co/) with zero third-party dependencies
- In-memory cache with automatic expiration to reduce API calls
- Catch probability scales with base experience
