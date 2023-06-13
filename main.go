package main

import (
	"flag"
	"fmt"
	"go-sprite-split/src/go-sprite-split"
	"time"
)

func main() {
	start := time.Now()
	var inputImgPath string
	var outputFolderPath string
	var mode string
	flag.StringVar(&inputImgPath, "input", "example/rounded-rectangles.png", "the path of source image to split")
	flag.StringVar(&outputFolderPath, "output", "output", "the path to save the sprites")
	flag.StringVar(&mode, "mode", "sprite", "sprite: output json file; split: output split images")
	flag.Parse()
	fmt.Println("input path:", inputImgPath)
	fmt.Println("output path:", outputFolderPath)
	fmt.Println("parsing image  ...")
	ok := go_sprite_split.Process(inputImgPath, outputFolderPath, mode)
	if ok == true {
		fmt.Printf("finished! (in %s)", time.Since(start))
	} else {
		fmt.Printf("failure, please chekck the parameters and try again. (in %s)", time.Since(start))
	}
}
