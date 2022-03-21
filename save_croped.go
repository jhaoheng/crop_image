package crop_image

import (
	"fmt"
	"io/ioutil"
)

/*
save crops for test
*/
func (mi *ImageObj) SaveCroped(output_folder_path, crop_name string) {
	for index, crop := range mi.Crops {
		name_index := fmt.Sprintf("%v/%v-%v.%v", output_folder_path, crop_name, index, crop.FileType)
		err := ioutil.WriteFile(name_index, crop.Bytes, 0755)
		if err != nil {
			panic(err)
		}
	}
}
