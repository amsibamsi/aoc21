package d20p2

import (
	"strconv"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	lines := aoc21.ToLines(input)
	algText := lines[0]
	imgText := lines[2:]
	alg := [512]int{}
	for i := 0; i < len(alg); i++ {
		if algText[i] == '#' {
			alg[i] = 1
		}
	}
	img := newImage(len(imgText), len(imgText[0]), 0)
	for y, line := range imgText {
		for x, c := range line {
			if c == '#' {
				img.set(x, y, 1)
			}
		}
	}
	for passes := 0; passes < 50; passes++ {
		newBg := alg[0]
		if img.bg == 1 {
			newBg = alg[511]
		}
		newImg := newImage(img.width+2, img.height+2, newBg)
		for x := 0; x < newImg.width; x++ {
			for y := 0; y < newImg.height; y++ {
				xo := x - 1
				yo := y - 1
				enc := img.get(xo-1, yo-1)<<8 |
					img.get(xo, yo-1)<<7 |
					img.get(xo+1, yo-1)<<6 |
					img.get(xo-1, yo)<<5 |
					img.get(xo, yo)<<4 |
					img.get(xo+1, yo)<<3 |
					img.get(xo-1, yo+1)<<2 |
					img.get(xo, yo+1)<<1 |
					img.get(xo+1, yo+1)
				newImg.set(x, y, alg[enc])
			}
		}
		img = newImg
	}
	count := 0
	for _, v := range img.data {
		count += v
	}
	return strconv.Itoa(count), nil
}

type image struct {
	data          map[[2]int]int
	width, height int
	bg            int
}

func newImage(width, height, bg int) *image {
	return &image{map[[2]int]int{}, width, height, bg}
}

func (img *image) get(x, y int) int {
	if x < 0 || y < 0 || x >= img.width || y >= img.height {
		return img.bg
	}
	return img.data[[2]int{x, y}]
}

func (img *image) set(x, y, v int) {
	img.data[[2]int{x, y}] = v
}
