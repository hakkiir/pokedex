# Pokedex

A command line Pokedex game that fetches data from PokeAPI

## Installation

You need the latest [Go toolchain](https://golang.org/dl/) installed
You can then install `pokedex` with:

```bash
go install https://github.com/hakkiir/pokedex@latest
```

## Usage

Run the program:
```bash
pokedex
```
## Commands

Exit the program:

```bash
Pokedex > exit
```

Display commands:

```bash
Pokedex > help
```

Show 20 next location areas:

```bash
Pokedex > map
```

Show 20 previous location areas:

```bash
Pokedex > mapb
```

Explore a location:

```bash
Pokedex > explore <location name>
```

Try to catch a Pokemon:

```bash
Pokedex > catch <pokemon>
```

Inspect a catch Pokemon:

```bash
Pokedex > inspect <pokemon>
```

List all catch Pokemons:

```bash
Pokedex > pokedex
```
