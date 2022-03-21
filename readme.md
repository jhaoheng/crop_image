# Usage

```
reader, _ := os.Open("")
b, _ := ioutil.ReadAll(reader)
handle := NewMedia(b)
handle.ReadImage()
b64, _ := handle.CropImage(image.Rect(0, 300, 800, 800))
```