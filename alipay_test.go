package alipay_test

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"os"
)

var (
	appID = "2016073100129537"

	// RSA2(SHA256)
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2MhEVUp+rRRyAD9HZfiSg8LLxRAX18XOMJE8/MNnlSSTWCCoHnM+FIU+AfB+8FE+gGIJYXJlpTIyWn4VUMtewh/4C8uwzBWod/3ilw9Uy7lFblXDBd8En8a59AxC6c9YL1nWD7/sh1szqej31VRI2OXQSYgvhWNGjzw2/KS1GdrWmdsVP2hOiKVy6TNtH7XnCSRfBBCQ+LgqO1tE0NHDDswRwBLAFmIlfZ//qZ+a8FvMc//sUm+CV78pQba4nnzsmh10fzVVFIWiKw3VDsxXPRrAtOJCwNsBwbvMuI/ictvxxjUl4nBZDw4lXt5eWWqBrnTSzogFNOk06aNmEBTUhwIDAQAB"

	privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC/x1fGLzA0A6oE06HwO/muHQFNvko2CTraMNpLnfpnVwOqrYmakFiQGfMlccBcSAHA9eBfxsau4wyNkpbw/i6/dG+PwCb17BbBQrZTubZkHc6j0xa3JiiIsZfUgDgUV6dfOk+CbRbLNp4tfjInyP6FWtFuDBtnNv0CrQzJ6mcrk9WhI8sA9fG9asgFOiTX6WK6CHHAP/tH6V1pBAK1YLuccHA44KEuRfqN6cI6qLoutPEnjJwZNZNy2eAUwk7hi0yar/bOvD6i82qbEclABnC5EWcqVm4WgQo9bIwZi1u8r0BYh8Ly5uffQ7kX2R8swgfAAcKHRSKk9JsoPBDYGdazAgMBAAECggEAftDJ+RKxNGQL5J8xoZN3WqxxdUBVAIB/+1J8t64irH08vnt818mF1txiDau3wQ4Yosm0coEWVwVrAp8h0wCyVwYe05cD7HEO5wM/YAQZWQDg6TIn/jpsaV/Of8W9z63azrzXYO7UONLILFsbvXVK4VIZ1USKRml1S3S1VTv8tzIoARl3jkK34X0y+s8QJtKx84O9M5UMc7aTjZ+LBLMs9WYzjBVRzbEDXk7zuO0OFGIAr9G95DsnZNZxquTIJHwn5t/kYYDv0t0Te9Eqv8cCTsg1VfhPpT1C4MXfSao1GyWhBs0Z2+o7rmdp80e+5XNSKo91JftdDCD9w46aFKxRuQKBgQDwFBdoefemwdRnHzMTBAVXJ3qvPFv7uOZWVRP6MbWlKJnhYBbHGrFhVeGfZ1t9q0PFagY4yysMz9CnyDzWcwbTWO2YmvZEoWEBtCX3eKka+9oI2gnwvHsekVKrt0RNwIIATu1gyTQuVxXDF5ohRZqQPIEeATLhmnqUvB3CyHXXlQKBgQDMfz8Tl1M5xh1zGTscu6kQL0wJv4WyFlPlQg30jvMs2WNBfY7/ybZMdv6hlSjLfRFWi36RiOHkfwzkJiI/QqSbQ6+rkoLSWjK3lJ/NZuDAcSV1BRWnzCgTRh1onMYH9SeYj2Z0kE43W3epsxaOWC5c4Pxt+uING559pl32hVJDJwKBgFzP55zESjF1jZ+wOBaJ32aqJIvKUeUmviVyDp8SyJ5o69RtSWD/uMNgaDEvy/bxVuni6zTZlnMFhCzZBGwDXAgowPHWNo6htNbrxiG9y+Jaqxfw8xbWRuKyW7t2xjkoyxlDahGYt0uS7x7U6Qkj13Ubbu8il4EQL4OAliDPTn9hAoGBAJK1Y2JBuSKEoFIXstHdS9/hnrKLpXXrEMszxWFDQPBXER9F3dZiNxfKcskngnniQxMMxoPQaQcNowj21rD9tavyNlBnrtUMgsAzryWj/e/x/IZkXHHiRIn3TlfjySqiVYLctgtAD+0lAMdNRQxf9PsLKe4ZBB2VR/Iq7dRTN34XAoGAdb026vh5PErWcg5nIaAKovV25iwlBw9bIhxExNKM6hBnllSFBlPDUd3YJbQi7Z9cTERJMiPBcnRklecKAVTxpxBfm+0i0LUGUxq5zD1dPHVEEJVB1NnsNJe9OVZtbBkdUM/zl2KgsT2Z+jn4eHdKYcOLdNpVePF4mxSAeRGA5fs="
)

var client *alipay.Client

func init() {
	var err error
	client, err = alipay.New(appID, aliPublicKey, privateKey, false)

	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		os.Exit(-1)
	}

	if client.IsProduction() {
		// https://docs.open.alipay.com/291/105971#Krqvg
		fmt.Println("加载证书", client.LoadAppPublicCertFromFile("appCertPublicKey_2017011104995404.crtZX"))
		fmt.Println("加载证书", client.LoadAliPayRootCertFromFile("alipayRootCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayPublicCertFromFile("alipayCertPublicKey_RSA2.crt"))
	}
}
