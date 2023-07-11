package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hallazzang/gosang"
)

func main() {
	f, err := os.Open(`data\AdvancedSamun_M.S32`)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sp, err := gosang.OpenSprite(f)
	if err != nil {
		panic(err)
	}
	for i := 0; i < sp.FrameCount(); i++ {
		frame, err := sp.Frame(i)
		if err != nil {
			panic(err)
		}
		func() {
			out, err := os.Create(fmt.Sprintf("output%d.png", i))
			if err != nil {
				panic(err)
			}
			defer out.Close()
			if err := png.Encode(out, frame.Image()); err != nil {
				panic(err)
			}
		}()
	}
}
