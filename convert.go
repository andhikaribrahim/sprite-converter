package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hallazzang/gosang"
)

func main() {
	err := filepath.Walk(`data`, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".S32" {
			return nil
		}

		fmt.Printf("name: %s\n", strings.Split(info.Name(), ".")[0])
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		defer f.Close()

		sp, err := gosang.OpenSprite(f)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		for i := 0; i < sp.FrameCount(); i++ {
			frame, err := sp.Frame(i)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			func() {
				outputDir := fmt.Sprintf(`output\%s`, strings.Split(info.Name(), ".")[0])

				if _, err := os.Stat(outputDir); os.IsNotExist(err) {
					err := os.MkdirAll(outputDir, os.ModePerm)
					if err != nil {
						log.Fatal(err)
						panic(err)
					}
				}

				// System cannot find the path specified. should create new 'output' folder if not exists and replace the path
				out, err := os.Create(fmt.Sprintf(`%s\%s_%d.png`, outputDir, strings.Split(info.Name(), ".")[0], i))
				if err != nil {
					log.Fatal(err)
					panic(err)
				}
				defer out.Close()
				if err := png.Encode(out, frame.Image()); err != nil {
					log.Fatal(err)
					panic(err)
				}
			}()
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
