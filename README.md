# ebiten-game-template

A basic skeleton for a 2D game using the [Ebiten](https://ebiten.org/) library.

-----✂️-----

> ⚠️ After cloning this repository:

> Write your OWN name name in the LICENSE file and run this command to replace the game name (tested on Linux and Mac):

```bash
grep -Rl ebiten-game-template | xargs sed -i '' -e "s/ebiten-game-template/${PWD##*/}/g"
```

> it assumes that the game name is the name of the current folder because that is what `go build` will call it.

> Then delete this section from the README, and start editing `main.go` to make your own game!

-----✂️-----

## For game testers

<!-- TODO: add a link to the latest downloads page -->

Game controls:
- F: toggle full-screen
- Q: quit the game
- Space: move up

## For programmers

Make sure you have [Go 1.17 or later](https://go.dev/) to contribute to the game

To build the game yourself, run: `go build .` it will produce an ebiten-game-template file and on Windows ebiten-game-template.exe.

To run the tests, run: `go test ./...` but there are no tests yet.

The project has a very simple, flat structure, the first place to start looking is the main.go file.
