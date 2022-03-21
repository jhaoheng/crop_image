package crop_image

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_read_image(t *testing.T) {

	type Data struct {
		input_file       string
		output_file_type string
	}

	datas := []Data{
		0: {
			input_file:       "../test_image/unknow.png",
			output_file_type: "png",
		},
		1: {
			input_file:       "../test_image/unknow2.jpg",
			output_file_type: "jpeg",
		},
	}

	for _, data := range datas {
		reader, _ := os.Open(data.input_file)
		b, _ := ioutil.ReadAll(reader)
		info, err := NewMedia(b).ReadImage()
		assert.NoError(t, err)
		assert.Equal(t, data.output_file_type, info.FileType)
		fmt.Printf("file_type: %v, size: %v KB\n", info.FileType, info.Size.KB())

		// fmt.Println(base64.RawStdEncoding.EncodeToString(info.MediaBytes))
	}
}
