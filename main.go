package GoGoSnakeeee

import (
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

func drawSnakes() {

}

func drawWorld() {

}

func drawEnv() {

}

func main() {

	mapEnv := [25]string{}

	mapEnv[0] = "|-----------------------------------------------------------------------------|"
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
	mapEnv[24] = "------------------------------------------------------------------------------"

	err := termbox.Init()

	if err != nil {
		panic(err)
	}

}
