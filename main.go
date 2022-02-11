package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	helloMsg := emoji.Sprintf("Hello :world_map:!")

	fmt.Println(helloMsg)
}
