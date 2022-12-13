package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	// tmt := "2022-12-12 12:09:07.39516 +0530 IST m=+0.002106201"

	// tmp := time.Now().UTC().Format(tmt)

	pl(time.Since(time.Now()))

}
