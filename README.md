Quadria v1.0
============

## About

Here I just try to recreate one game I had seen back in 2009 on one of my school's computers. I failed to find it anywhere else and forgot it's name, so I just try to make new one.
It has only hotseat console mode now. UI version is developing here: https://github.com/rossus/codex-gen-quadria-ui

## Install

```
go get github.com/rossus/quadria
```

## Run

Go to main catalog, then enter:
```
go run quadria.go
```

Or you can build it by:
```
go build quadria.go
```
And then run it from the file.

## Gameplay

At first you should enter names of blue and red player, then enter the size of game board.
After that game will begin. You can add 1 point anywhere, except your rival's tiles. You can add one point on tile (x, y) by using:
```
go x y
```
If tile's value exceeds number of it's neighbours it will give each neighbour one point and spread it's color on them.
To win you must spread your color on each tile of the board.

## Versions
### v0.1
Initial version

### v1.0
Added sessions