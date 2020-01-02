package main

import (
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"io"
	"math"
	"os"
	"reflect"
	"sort"
)

//最初の画像を背景と仮定
//nxnの平均値からの差を保存
//背景のnxnの平均値からの差を加算
func ReadImageFromPath(path string) image.Image {
	if file, err := os.Open(path); err == nil {
		defer file.Close()
		return ReadImageFromFile(file)
	}
	return nil
}
func ReadImageFromFile(file io.Reader) image.Image{
	if img, _, e := image.Decode(file); e == nil {
		return img
	}
	return nil
}
func JpegToFile(w io.Writer, img image.Image){
	jpeg.Encode(w, img, nil)
}
func JpegToPath(path string, img image.Image) {
	if file, err := os.Create(path); err == nil {
		defer file.Close()
		JpegToFile(file,img)
	}
}

type intc = uint16
type intt = int32
type RawImage struct {
	Size image.Point
	RGB  []intc
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
func NewRawImage(img image.Image) RawImage {
	var x, y, z int
	size := img.Bounds().Size()
	s := RawImage{
		Size: size,
		RGB:  make([]intc, size.X*size.Y*3),
	}
	for y = 0; y < s.Size.Y; y++ {
		for x = 0; x < s.Size.X; x++ {
			z = (y*s.Size.X + x) * 3
			r, g, b, _ := img.At(x, y).RGBA()
			s.RGB[z+0], s.RGB[z+1], s.RGB[z+2] = intc(r), intc(g), intc(b)
		}
	}
	return s
}
func (s *RawImage) NewRawImage() RawImage {
	o := RawImage{
		Size: s.Size,
		RGB:  make([]intc, s.Size.X*s.Size.Y*3),
	}
	copy(o.RGB, s.RGB)
	return o
}
func (s *RawImage) ToImage() image.Image {
	var x, y, p, z int
	img := image.NewRGBA(image.Rect(0, 0, s.Size.X, s.Size.Y))
	p = 16 - reflect.TypeOf(s.RGB[0]).Bits()
	for y = 0; y < s.Size.Y; y++ {
		for x = 0; x < s.Size.X; x++ {
			z = (y*s.Size.X + x) * 3
			img.Set(x, y, color.RGBA64{
				R: uint16(s.RGB[z+0]) << uint(p),
				G: uint16(s.RGB[z+1]) << uint(p),
				B: uint16(s.RGB[z+2]) << uint(p),
			})
		}
	}
	return img
}
func (s *RawImage) AverageFilter(n int) {
	o := s.NewRawImage()
	var x, y, p, z int
	var r, g, b intt
	for y = 0; y < s.Size.Y; y++ {
		for x = 0; x < s.Size.X; x++ {
			r, g, b = 0, 0, 0
			for p = max(x-n/2, 0); p < min(x-n/2+n, s.Size.X); p++ {
				z = (y*s.Size.X + p) * 3
				r, g, b = r+intt(s.RGB[z+0]), g+intt(s.RGB[z+1]), b+intt(s.RGB[z+2])
			}
			z, p = (y*s.Size.X+x)*3, min(x-n/2+n, s.Size.X)-max(x-n/2, 0)
			o.RGB[z+0], o.RGB[z+1], o.RGB[z+2] = intc(r/intt(p)), intc(g/intt(p)), intc(b/intt(p))
		}
	}
	for y = 0; y < s.Size.Y; y++ {
		for x = 0; x < s.Size.X; x++ {
			r, g, b = 0, 0, 0
			for p = max(y-n/2, 0); p < min(y-n/2+n, s.Size.Y); p++ {
				z = (p*s.Size.X + x) * 3
				r, g, b = r+intt(o.RGB[z+0]), g+intt(o.RGB[z+1]), b+intt(o.RGB[z+2])
			}
			z, p = (y*s.Size.X+x)*3, min(y-n/2+n, s.Size.Y)-max(y-n/2, 0)
			s.RGB[z+0], s.RGB[z+1], s.RGB[z+2] = intc(r/intt(p)), intc(g/intt(p)), intc(b/intt(p))
		}
	}
}
func (s *RawImage) Toning(t *RawImage) {
	const limit = 100
	u := [...][]int{make([]int, limit), make([]int, limit), make([]int, limit)}
	var x, z int
	for x = 0; x < limit; x++ {
		z = (s.Size.X * s.Size.Y * x / limit) * 3
		u[0][x] = int(t.RGB[z+0]) - int(s.RGB[z+0])
		u[1][x] = int(t.RGB[z+1]) - int(s.RGB[z+1])
		u[2][x] = int(t.RGB[z+2]) - int(s.RGB[z+2])
	}
	sort.Ints(u[0])
	sort.Ints(u[1])
	sort.Ints(u[2])
	gosa := [...]intt{intt(u[0][limit/2]), intt(u[1][limit/2]), intt(u[2][limit/2])}
	for x = 0; x < s.Size.X*s.Size.Y*3; x++ {
		t.RGB[x] -= intc(gosa[x%3])
	}
}
func (b *RawImage) Synth(fore []RawImage, colorrate float64,filterrate float64) RawImage {
	var y, x, z, n int
	n = int(math.Sqrt(float64(b.Size.X*b.Size.Y))*filterrate)
	o := b.NewRawImage()
	ba:=b.NewRawImage()
	ba.AverageFilter(n)
	for y = 0; y < len(fore); y++ {
		f:=&fore[y]
		b.Toning(f)
		fa:=f.NewRawImage()
		fa.AverageFilter(n)
		for x = 0; x < b.Size.X*b.Size.Y; x++ {
			z=0
			z+=abs(int(ba.RGB[x*3+0])-int(fa.RGB[x*3+0]))
			z+=abs(int(ba.RGB[x*3+1])-int(fa.RGB[x*3+1]))
			z+=abs(int(ba.RGB[x*3+2])-int(fa.RGB[x*3+2]))
			if z > int(float64(^intc(0))*colorrate) {
				o.RGB[x*3+0] = f.RGB[x*3+0]
				o.RGB[x*3+1] = f.RGB[x*3+1]
				o.RGB[x*3+2] = f.RGB[x*3+2]
			}
		}
	}
	return o
}