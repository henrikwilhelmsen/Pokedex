# Pokedex

Go Pokedex CLI app, built as a part of the [boot.dev](boot.dev) backend developer course, using the [PokeAPI](https://pokeapi.co/) to get game information.

## Setup

This project uses [Nix](https://nixos.org/) to manage the dev environment.

After cloning the project, run `nix develop` in the project directory to enter a development shell.

The following requirements will be installed and configured automatically:

- [Go](https://go.dev/)
- [golangci-lint](https://golangci-lint.run/)
- [golangci-lint-lsp](https://github.com/nametake/golangci-lint-langserver)
- [boot.dev cli](https://github.com/bootdotdev/bootdev)

See the [nix flake](./flake.nix) for the full configuration.

## Usage

Run the CLI REPL with `go run`:

```shell
go run .
```

Use the help command to list available commands, or exit to exit the app:

```shell
Pokedex > help
```

```shell
Pokedex > exit
```

## Examples

### Exploring the Map

Use the `map` and `mapb` command to list locations in the Pokemon world:

```shell
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior
```

The `map` command displays the next 20 locations available and the `mapb` command displays the previous 20 locations.

Use the `explore` command to explore a location and list the available Pokemon encounters:

```shell
Pokedex > explore valley-windworks-area
tentacool
tentacruel
shellder
magikarp
gyarados
mareep
aipom
heracross
elekid
wurmple
silcoon
cascoon
wingull
pelipper
electrike
bidoof
shinx
burmy
combee
pachirisu
buizel
cherubi
shellos
gastrodon
drifloon
munchlax
finneon
lumineon
```

### Catching and Inspecting Pokemon

Use the `catch` command to catch wild Pokemon:

```shell
Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp broke free!
Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
```

Use the `inspect` command to inspect caught Pokemon:

```shell
Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  -water
```

Use the `pokedex` command to list caught Pokemon:

```shell
Your Pokedex:
- magikarp
- tentacool
- pikachu
```
