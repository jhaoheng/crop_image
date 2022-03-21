package crop_image

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCropReadByByte(t *testing.T) {

	type Data struct {
		input_file       string
		output_file_type string
	}

	datas := []Data{
		0: {
			input_file:       "./test_image/unknow.png",
			output_file_type: "png",
		},
		1: {
			input_file:       "./test_image/unknow.jpg",
			output_file_type: "jpeg",
		},
	}

	for _, data := range datas {
		reader, _ := os.Open(data.input_file)
		b, _ := ioutil.ReadAll(reader)
		new_crop, err := NewCropReadByByte(b)
		assert.NoError(t, err)
		info := new_crop.GetOriImgInfo()
		assert.Equal(t, data.output_file_type, info.FileType)
		fmt.Printf("file_type: %v, size: %v KB\n", info.FileType, info.Size.KB())

		// fmt.Println(base64.RawStdEncoding.EncodeToString(info.MediaBytes))
	}
}
