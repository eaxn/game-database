# game-database


## This project is not in alpha phase yet

1. ⚠️This repository doesn't contain the whole data of the most games, because it would need to much memory if you want to download it. This means you have to download the whole data, if you want to try to the run a discord bot / service at your own.⚠️
2. Much things, like price recognition or game bundle detection (steam) is not created yet. Please don't think, that this project is bad or something like this. We are just starting 🚀.
3. We create a GUI client for interacting with our API servers for creating the best experience of gamers, who are not on a discord server with a bot that uses this simple server / bot software
4. The server can't get information about any game forbidden / not available for your country (will be fixed)
## Planned features
- Detection of all free games on every popular gaming platform / gaming launcher (steam, epic games, gog, nintendo eshop, playstation store, itch.io)
- GUI client for interacting with our API servers to provide as many free games as possible
- Detection of all games in discount on every popular gaming platform / gaming launcher
- Own discord bot (for telling your buddies and you, which games are on discount, which games are currently free / or forever free)
- Price compararison (to get the game for the cheapest price)
- Game rating (for the same games on different platforms)
- Provide help, for installing a hardware emulator
- Let the user start any of his games with our own launcher

## Modules
You can easily create modules, but how? You just create a module (go file / .go file) and drag it into the modules/ folder, then it will be used for the wanted build of the executable.
### Example Module
```go
package main

var moduleLoader_example = func() error { 
    // do your stuff here
    // and DON'T create a loop in this function
    println("example plugin works lol!")
    return nil
}()
```

## Planned applications, build with the base of this project
- a CLI application (idk why, but I like CLI applications) for checking, which games are currently on discount or free
- a desktop application (same reason as above)
- a web server (just for api purposes)
- a discord bot (same reason as above the web server)
- a game launcher, which can launch games from Steam, Epic Games...

## Main game publisher platforms for this project
- Steam
- Epic Games Store
- GOG
- Humble Bundle
- Indie Planet
- Itch.io
- Jagex
