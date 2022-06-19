package main

import (
	"github.com/pquerna/otp/totp"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})
	f, _ := os.Create("qr.png")
	img, _ := key.Image(200, 200)
	png.Encode(f, img)
	code, _ := totp.GenerateCode(key.Secret(), time.Now())
	log.Println("code:", code)
	valid := totp.Validate(code, key.Secret())
	if valid {
		log.Println("success!")
	}
}
