package crop_image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
)

// ex: crop_rect := image.Rect(0, 0, 800, 600)
func (mi *ImageObj) CropImage(crop_rect image.Rectangle) (b64_img string, err error) {
	if mi.Ori.FileType == "png" {
		b64_img = mi.CropImageToPNG(crop_rect)
	} else if mi.Ori.FileType == "jpeg" {
		b64_img = mi.CropImageToJPEG(crop_rect)
	}
	if len(b64_img) == 0 {
		return "", fmt.Errorf("file_type not match jpeg / png")
	}
	return
}

func (mi *ImageObj) CropImageToJPEG(crop_rect image.Rectangle) (b64_img string) {

	m := image.NewRGBA(crop_rect)
	draw.Draw(m, m.Bounds(), mi.Ori.Image, image.Point{0, 0}, draw.Src)

	buf := bytes.Buffer{}
	jpeg.Encode(&buf, m, &jpeg.Options{Quality: 100})

	mi.Crops = append(mi.Crops, ImageInfo{
		Image:     m,
		Bytes:     buf.Bytes(),
		FileType:  mi.Ori.FileType,
		Size:      ByteSize(len(buf.Bytes())),
		Rectangle: m.Bounds(),
	})

	return base64.RawStdEncoding.EncodeToString(buf.Bytes())
}

func (mi *ImageObj) CropImageToPNG(crop_rect image.Rectangle) (b64_img string) {
	m := image.NewRGBA(crop_rect)
	draw.Draw(m, m.Bounds(), mi.Ori.Image, crop_rect.Min, draw.Src)

	buf := bytes.Buffer{}
	png.Encode(&buf, m)

	mi.Crops = append(mi.Crops, ImageInfo{
		Image:     m,
		Bytes:     buf.Bytes(),
		FileType:  mi.Ori.FileType,
		Size:      ByteSize(len(buf.Bytes())),
		Rectangle: m.Bounds(),
	})

	return base64.RawStdEncoding.EncodeToString(buf.Bytes())
}
