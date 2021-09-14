package subtraction

import (
	"fmt"

	addition "github.com/JogeAW/Joge/src/maths"
)

var a uint = 3     //Переменная типа uint
var c int = int(a) // Конвертация в int
var b = addition.Summ
var Result = b - c

func subtraction() {
	fmt.Println(Result)
}
