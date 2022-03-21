package crop_image

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_only_CropImage(t *testing.T) {
	file := "./test_image/unknow.png"
	reader, _ := os.Open(file)
	b, _ := ioutil.ReadAll(reader)
	new_crop, _ := NewCropReadByByte(b)
	info := new_crop.GetOriImgInfo()
	fmt.Printf("file_type: %v, size: %v KB\n", info.FileType, info.Size.KB())
	//
	b64, err := new_crop.CropImage(image.Rect(0, 300, 800, 800))
	assert.NoError(t, err)
	fmt.Println(b64)
}

func Test_save_CropImage(t *testing.T) {

	type Data struct {
		input_file  string
		input_crops []image.Rectangle
	}

	datas := []Data{
		0: {
			input_file: "./test_image/unknow.png",
			input_crops: []image.Rectangle{
				image.Rect(0, 0, 800, 400),
				image.Rect(0, 200, 800, 800),
			},
		},
	}

	for _, data := range datas {
		reader, _ := os.Open(data.input_file)
		b, _ := ioutil.ReadAll(reader)
		new_crop, _ := NewCropReadByByte(b)
		info := new_crop.GetOriImgInfo()
		fmt.Printf("file_type: %v, size: %v KB\n", info.FileType, info.Size.KB())
		//
		for _, crop_rect := range data.input_crops {
			_, err := new_crop.CropImage(crop_rect)
			assert.NoError(t, err)
		}
		//
		new_crop.SaveCroped("./", "output")
	}
}
