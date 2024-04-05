package main

import (
	"github.com/ginx-contribs/str2bytes"
	"log"
)

func main() {
	bytes := str2bytes.Str2Bytes("hello world")
	log.Println(bytes)
	str := str2bytes.Bytes2Str(bytes)
	log.Println(str)
}
