# Pokedex

A Pokémon CLI game written in Go.

## Run

With Go installed, run from the project root:

```bash
go run .
```

An internet connection is required to fetch Pokémon and location data. Caught Pokémon are stored in memory for the current session.

## Commands

| Command                   | Description                                                                   |
| ------------------------- | ----------------------------------------------------------------------------- |
| `help`                    | Show all available commands.                                                  |
| `map`                     | Show the next page of up to 20 location areas.                                |
| `mapb`                    | Show the previous page of location areas.                                     |
| `explore <location_name>` | List Pokémon found in a location area. Use a location name returned by `map`. |
| `catch <pokemon_name>`    | Try to catch a Pokémon. Catching is random, so an attempt can fail.           |
| `inspect <pokemon_name>`  | Show a caught Pokémon's name, height, weight, stats, and types.               |
| `pokedex`                 | List all Pokémon you have caught during the current session.                  |
| `exit`                    | Close the game.                                                               |

For example, enter these commands at the `Pokedex >` prompt:

```text
map
explore pastoria-city-area
catch tentacool
pokedex
inspect tentacool
exit
```

The `inspect` command requires a successful catch first.
