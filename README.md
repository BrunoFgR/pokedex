# Pokedex CLI

A command-line interface (CLI) Pokedex application written in Go. This application allows you to explore the Pokemon world, catch Pokemon, and build your personal Pokedex by interfacing with the [PokeAPI](https://pokeapi.co/).

## Overview

Pokedex CLI is an interactive terminal application that simulates the experience of a Pokemon trainer. You can explore different locations in the Pokemon world, discover which Pokemon can be encountered in each area, attempt to catch them, and build your personal collection.

## Installation

### Prerequisites
- Go 1.16 or later

### Steps
1. Clone the repository:
   ```
   git clone https://github.com/yourusername/pokedexcli.git
   ```

2. Navigate to the project directory:
   ```
   cd pokedexcli
   ```

3. Build the application:
   ```
   go build
   ```

4. Run the application:
   ```
   ./pokedexcli
   ```

## Usage

Once you start the application, you'll be presented with a command prompt:

```
Pokedex >
```

Type `help` to see a list of available commands.

## Commands

| Command | Description |
|---------|-------------|
| `help` | Displays a help message with all available commands |
| `map` | Get the next page of locations |
| `mapb` | Get the previous page of locations |
| `explore <location_name>` | Explore a location to find which Pokemon can be encountered there |
| `catch <pokemon_name>` | Attempt to catch a Pokemon |
| `inspect <pokemon_name>` | View details about a caught Pokemon |
| `pokedex` | View all the Pokemon you've caught |
| `exit` | Exit the application |

## Features

- **Explore Pokemon World**: Navigate through different areas in the Pokemon world
- **Discover Pokemon**: Find out which Pokemon can be encountered in each location
- **Catch Pokemon**: Try your luck at catching Pokemon (success rate based on Pokemon's base experience)
- **Build Your Pokedex**: Keep track of all the Pokemon you've caught
- **Pokemon Info**: View detailed information about Pokemon in your Pokedex
- **Caching System**: API responses are cached to improve performance and reduce API calls

## Project Structure

- **main.go**: Entry point of the application
- **repl.go**: Implementation of the Read-Eval-Print-Loop for the CLI
- **command_*.go**: Individual command implementations
- **internal/pokeapi/**: Package for interacting with the PokeAPI
- **internal/pokecache/**: Package implementing caching functionality

## Future Enhancements

Here are some ideas for extending this project:

* Update the CLI to support the "up" arrow to cycle through previous commands
* Simulate battles between pokemon
* Add more unit tests
* Refactor your code to organize it better and make it more testable
* Keep pokemon in a "party" and allow them to level up
* Allow for pokemon that are caught to evolve after a set amount of time
* Persist a user's Pokedex to disk so they can save progress between sessions
* Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
* Random encounters with wild pokemon
* Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon

## Acknowledgments

This project was developed as part of the [Boot.dev](https://www.boot.dev/courses/build-pokedex-cli-golang) curriculum, which provided the foundational knowledge and guidance for building a CLI application in Go.

## License

This project is licensed under the terms found in the LICENSE file in the repository.