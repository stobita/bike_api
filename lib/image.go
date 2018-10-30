package lib

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

func ImageUpload(fileString string) (string, error) {
	var result string
	fileType := fileString[strings.IndexByte(fileString, ':')+1 : strings.IndexByte(fileString, ';')]
	unbased, err := base64.StdEncoding.DecodeString(fileString[strings.IndexByte(fileString, ',')+1:])
	fileName := sha1.Sum([]byte(unbased))
	switch fileType {
	case "image/png":
		result = S3Upload(bytes.NewReader(unbased), fmt.Sprintf("%x.png", fileName))
		// ファイル出力
		// img, err := png.Decode(bytes.NewReader(unbased))
		// f, err := os.OpenFile(fmt.Sprintf("files/%x.png", fileName), os.O_WRONLY|os.O_CREATE, 0777)
		// if err != nil {
		// 	return "", err
		// }
		// png.Encode(f, img)
	case "image/jpeg":
		result = S3Upload(bytes.NewReader(unbased), fmt.Sprintf("%x.jpeg", fileName))
		// ファイル出力
		// img, err := jpeg.Decode(bytes.NewReader(unbased))
		// f, err := os.OpenFile(fmt.Sprintf("files/%x.jpeg", fileName), os.O_WRONLY|os.O_CREATE, 0777)
		// if err != nil {
		// 	return "", err
		// }
		// jpeg.Encode(f, img, nil)
	}
	if err != nil {
		return "", err
	}
	return result, nil
}
