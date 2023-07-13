package main

import (
	"fmt"

	"github.com/wenlng/go-captcha/captcha"
)

func main() {
	// Captcha Single Instances
	capt := captcha.GetCaptcha()

	// Generate Captcha
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		panic(err)
		return
	}

	// Main image base64 code
	fmt.Println(len(b64))

	// Thumb image base64 code
	fmt.Println(len(tb64))

	// Only key
	fmt.Println(key)

	// Dot data For verification
	fmt.Println(dots)
}
