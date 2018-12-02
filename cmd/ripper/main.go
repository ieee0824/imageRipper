package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	ripper "github.com/ieee0824/imageRipper"
)

func main() {
	name := flag.String("f", "", "")
	oDir := flag.String("o", "", "")
	bs := flag.Int("bs", 0, "")
	flag.Parse()

	f, err := os.Open(*name)
	if err != nil {
		panic(err)
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	r := &ripper.Ripper{
		BlockSize: &ripper.Size{
			*bs,
			*bs,
		},
	}

	results := r.Do(img)

	if _, err := ioutil.ReadDir(*oDir); err != nil {
		if err := os.Mkdir(*oDir, 0777); err != nil {
			log.Fatalln(err)
		}
	}

	for i, result := range results {
		dirNum := i / 1000
		subDirName := fmt.Sprintf("./%s/%04d", *oDir, dirNum)
		if _, err := ioutil.ReadDir(subDirName); err != nil {
			if err := os.Mkdir(subDirName, 0777); err != nil {
				log.Fatalln(err)
			}
		}
		w, err := os.Create(fmt.Sprintf("%s/%04d.jpeg", subDirName, i))
		if err != nil {
			log.Fatalln(err)
		}
		if err := jpeg.Encode(w, result, &jpeg.Options{100}); err != nil {
			log.Fatalln(err)
		}
		w.Close()
	}
}
