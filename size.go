package crop_image

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1),  1 * 2^10
	MB                             // 1 << (10*2),	1 * 2^20
	GB                             // 1 << (10*3),	1 * 2^30
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

func (size ByteSize) KB() ByteSize {
	return size / KB
}

func (size ByteSize) MB() ByteSize {
	return size / MB
}
