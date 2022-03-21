package crop_image

import "image"

type MediaInfo struct {
	Image     image.Image
	Bytes     []byte
	FileType  string
	Size      ByteSize
	Rectangle image.Rectangle
}

type MediaObj struct {
	Ori   MediaInfo
	Crops []MediaInfo
}
