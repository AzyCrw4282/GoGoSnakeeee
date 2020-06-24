package GoGoSnakeeee

import (
	"context"
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

type snake struct {
	x           int
	y           int
	snacksEaten int
}

//Implement one snake first and then follow on....
type snake2 struct {
	x           int
	y           int
	snacksEaten int
}

func drawSnakes(snk snake) {

	//For one snake.for both, initiate with the loops
	termbox.SetCell(snk.x, snk.y, '∩', termbox.Attribute(7), termbox.Attribute(3))
	termbox.SetCell(snk.x, snk.y+1, '|', termbox.Attribute(7), termbox.Attribute(3))
	termbox.SetCell(snk.x, snk.y+2, '|', termbox.Attribute(7), termbox.Attribute(3))
	termbox.SetCell(snk.x, snk.y+3, '|', termbox.Attribute(7), termbox.Attribute(3))
	termbox.SetCell(snk.x, snk.y+4, '|', termbox.Attribute(7), termbox.Attribute(3))
}

//creates the entire map environment
func drawWorld(mapEnvDup [25]string) {
	getColour := func(x int, y int, mapEnvDup2 [25]string) int {
		switch mapEnvDup2[y][x] {
		case '0':
			return 3
		case '=':
			return 10
		case '|':
			return 14
		case '-':
			return 13
		}
		return -1
	}
	getColour(0, 0, mapEnvDup)
	wid, hei := 60, 25
	for x := 0; x < wid; x++ {
		for y := 0; y < hei; y++ {
			termbox.SetCell(x, y, rune(mapEnvDup[y][x]), termbox.Attribute(getColour(x, y, mapEnvDup)), termbox.Attribute(getColour(x, y, mapEnvDup))) //sets the colour
		}
	}
}

func drawEnv(mapEnvDup [25]string, snk snake) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawSnakes(snk)
	drawWorld(mapEnvDup)

	termbox.Flush()

}

func main() {

	mapEnv := [25]string{}

	mapEnv[0] = "|-------------------------------------------===-------------------------------|"
	mapEnv[1] = "|                                                                             |"
	mapEnv[2] = "|                                                                             |"
	mapEnv[3] = "|                                                                             |"
	mapEnv[4] = "|                                                                             |"
	mapEnv[5] = "|                                                                             |"
	mapEnv[6] = "|                                                                             |"
	mapEnv[7] = "|                                                                             |"
	mapEnv[8] = "|                                                                             |"
	mapEnv[9] = "|                                  -------------                              |"
	mapEnv[10] = "|                                                                             |"
	mapEnv[11] = "|                                                                             |"
	mapEnv[12] = "|                       |                                                     |"
	mapEnv[13] = "|                       |                                                     |"
	mapEnv[14] = "|                       |                                                     |"
	mapEnv[15] = "|                       |                                                     |"
	mapEnv[16] = "|                       |                                                     |"
	mapEnv[17] = "|                                        -------------                        |"
	mapEnv[18] = "|                                                                             |"
	mapEnv[19] = "|                                                                             |"
	mapEnv[20] = "|                                                                             |"
	mapEnv[21] = "|                                                                             |"
	mapEnv[22] = "|                                                                             |"
	mapEnv[23] = "|                                                                             |"
	mapEnv[24] = "0-------------------------------------------===-------------------------------"

	err := termbox.Init()

	if err != nil {
		panic(err)
	}

}
