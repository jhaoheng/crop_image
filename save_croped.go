package crop_image

import (
	"fmt"
	"io/ioutil"
)

/*
save crops for test
*/
func (mi *MediaObj) SaveCroped(name string) {
	for index, crop := range mi.Crops {
		name_index := fmt.Sprintf("./crop_image_test_output/%v-%v.%v", name, index, crop.FileType)
		err := ioutil.WriteFile(name_index, crop.Bytes, 0755)
		if err != nil {
			panic(err)
		}
	}
}
