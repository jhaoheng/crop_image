package crop_image

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

func NewMedia(data []byte) *MediaObj {
	return &MediaObj{
		Ori: MediaInfo{
			Bytes: data,
			Size:  ByteSize(len(data)),
		},
		Crops: []MediaInfo{},
	}
}

func (mi *MediaObj) ReadImage() (read_media_info *MediaInfo, err error) {
	img, image_type, err := image.Decode(bytes.NewReader(mi.Ori.Bytes))
	if err != nil {
		return nil, err
	}
	//
	mi.Ori.FileType = image_type
	mi.Ori.Image = img
	mi.Ori.Rectangle = img.Bounds()
	return &mi.Ori, nil
}
