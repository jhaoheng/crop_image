package crop_image

func (mi *ImageObj) GetOriImgInfo() ImageInfo {
	return mi.Ori
}

func (mi *ImageObj) GetCropsInfo() []ImageInfo {
	return mi.Crops
}
