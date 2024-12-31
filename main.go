package depcalldemoh

import (
	"fmt"
)

func MyFunc1() {
	fmt.Println("depcalldemoh.MyFunc1")
	MyFunc2()
	MyFunc3()
}

func MyFunc2() {
	fmt.Println("depcalldemoh.MyFunc2")
	MyFunc3()
}

func MyFunc3() {
	fmt.Println("depcalldemoh.MyFunc3")
}
