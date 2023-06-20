package utils

import (
	"adr/backend/src/utils/s3Config"
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/arsmn/fastgql/graphql"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/sunshineplan/imgconv"
)

func ImageProcessor(media graphql.Upload, directory string, id string, index int) (string, error) {

	content, err := io.ReadAll(media.File)
	if err != nil {
		return "", err
	}

	fileBytes := bytes.NewReader(content)

	imageData, err := imgconv.Decode(fileBytes)

	if err != nil {
		return "", err
	}

	bounds := imageData.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	width := w
	height := h

	if w > 1000 {
		width = 1000
		height = width * h / w
	}

	var da []byte
	fourth := bytes.NewBuffer(da)
	finalData := imaging.Resize(imageData, width, height, imaging.Lanczos)
	webp.Encode(fourth, finalData, &webp.Options{Lossless: false, Quality: 90})

	finalFileByte := bytes.NewReader(fourth.Bytes())

	key := directory + "/" + id + "/" + fmt.Sprint(index) + media.Filename + ".webp"
	_, err = s3Config.Upload(key, finalFileByte, "public-read")

	if err != nil {
		log.Println("s3Config.Upload", err)
		return "", err
	}

	return key, nil
}
