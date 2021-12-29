package tool

import (
	"image/color"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	ID           string `json:"id"`
	Base64Blob   string `json:"base_64_blob"`
	VertifyValue string `json:"code"`
}

func GenerateCaptcha(context *gin.Context) {
	parameters := base64Captcha.ConfigCharacter{
		Height:             30,
		Width:              60,
		Mode:               3,
		ComplexOfNoiseText: 0,
		ComplexOfNoiseDot:  0,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4, //长度
		BgColor: &color.RGBA{ //背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
	}

	captchaID, captchaInterfaceInstance := base64Captcha.GenerateCaptcha("", parameters)

	bse64blob := base64Captcha.CaptchaWriteToBase64Encoding(captchaInterfaceInstance) //编码至base64

	captchaResult := CaptchaResult{ID: captchaID, Base64Blob: bse64blob}
	Success(context, gin.H{
		"captcha_result": captchaResult,
	})
}

//验证图形码
func VertifyCaptcha(id string, value string) bool {
	res := base64Captcha.VerifyCaptcha(id, value)
	return res
}
