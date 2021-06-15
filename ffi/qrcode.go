package main

import "C"

import (
	"bytes"
	"fmt"
	"github.com/liyue201/goqr"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
)

func recognizeFile(path string) string {
	fmt.Printf("recognize file: %v\n", path)
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return ""
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return ""
	}
	for _, qrCode := range qrCodes {
		return string(qrCode.Payload)
	}
	return ""
}

//export Recognize
func Recognize(pathExternal *C.char) *C.char {
	path := C.GoString(pathExternal)

	text := recognizeFile(path)
    return C.CString(text)
}

func main() {
	text := recognizeFile("testdata/image.png")
	fmt.Println(text)

	text2 := Recognize(C.CString("testdata/image.png"))
	fmt.Println(C.GoString(text2))
}