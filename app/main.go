package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/jhaoheng/crop_image"
)

func main() {
	input_file := "../test_image/unknow.png"
	reader, _ := os.Open(input_file)
	b, _ := ioutil.ReadAll(reader)
	new_crop, _ := crop_image.NewCropReadByByte(b)
	b64, err := new_crop.CropImage(image.Rect(0, 200, 800, 600))
	if err != nil {
		panic(err)
	}

	fmt.Println(b64)
}
