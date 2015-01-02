package main

import (
	"code.google.com/p/rsc/qr"
	"io/ioutil"
  "flag"
)

func main() {
  content := flag.String("content", "default", "the content to be encoded")
  fileName := flag.String("out", "qr", "the name of the outputted qr file")

  flag.Parse()

	code, err := qr.Encode(*content, qr.L)
  if err != nil {
    panic("error while encoding contents")
  }
	ioutil.WriteFile(*fileName + ".png", code.PNG(), 0644)
}
