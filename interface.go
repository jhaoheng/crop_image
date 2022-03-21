package crop_image

import (
	"image"
)

type ICropImage interface {
	CropImage(crop_rect image.Rectangle) (b64_img string, err error)
	SaveCroped(crop_name, output_folder_path string)
	//
	GetOriImgInfo() ImageInfo
	GetCropsInfo() []ImageInfo
}
