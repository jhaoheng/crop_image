package crop_image

import (
	"bytes"
	"image"
)

func NewCropReadByByte(data []byte) (ICropImage, error) {
	img, image_type, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	//
	image_obj := &ImageObj{
		Ori: ImageInfo{
			Bytes:     data,
			Size:      ByteSize(len(data)),
			Image:     img,
			FileType:  image_type,
			Rectangle: img.Bounds(),
		},
		Crops: []ImageInfo{},
	}

	return image_obj, nil
}

type ImageInfo struct {
	Image     image.Image
	Bytes     []byte
	FileType  string
	Size      ByteSize
	Rectangle image.Rectangle
}

type ImageObj struct {
	Ori   ImageInfo
	Crops []ImageInfo
}
