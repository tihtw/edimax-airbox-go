package main

import (
	"fmt"
	"github.com/tihtw/edimax-airbox-go"
)

var TOKEN = ""
var MAC = ""

func main() {
	device, _ := airbox.GetDevice(TOKEN, MAC)
	fmt.Println(device)
}
