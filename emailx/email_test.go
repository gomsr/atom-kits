package emailx

import (
	"fmt"
	"github.com/alice52/jasypt-go"
	"testing"
)

var secretQq string
var secretNet string
var secretGmail string

func init() {
	secretQq, _ = jasypt.New().DecryptWrapper("ENC(hlomTQKeIwivZYpT22kVC/oiRnewPAXza2LZo87/0PObwbdYVF/p5+NCb/069aZmP2D/p740TbMTl8W9uslWzg==)")
	secretNet, _ = jasypt.New().DecryptWrapper("ENC(mPbyo0f3VKO5kBGJaKOggTLbfqXR103iOfQhn548ff+EI1hMrj5q3YffhwMRKNdM)")
	secretGmail, _ = jasypt.New().DecryptWrapper("ENC(8YwWbLC7ZwTZZrE76TzTjtBq19M8NVq0AJRjzQ7jYpKQWOMQLxlka0foCqfdKbIN0Yrql8R8WUFaPhpnKg0tuw==)")
}

var register = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>WPP 邮箱验证码</title>
    <style>
        body {
            font-family: Arial, sans-serif;
        }
        .container {
            width: 100%;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            border: 1px solid #eaeaea;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        .header {
            background-color: #1d3e6e;
            color: white;
            padding: 10px;
            text-align: center;
            font-size: 24px;
        }
        .content {
            padding: 20px;
            text-align: center;
        }
        .code {
            font-size: 36px;
            font-weight: bold;
            margin: 20px 0;
        }
        .discount {
            font-size: 16px;
            color: #666;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            margin-top: 20px;
            background-color: #1d3e6e;
            color: white;
            text-decoration: none;
            border-radius: 5px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            WPP 邮箱验证码
        </div>
        <div class="content">
            <p>阁下您好，</p>
            <p>请填写以下验证码完成邮箱验证（{{VERIFICATION_MINITUE}}分钟内有效）</p>
            <div class="code">{{VERIFICATION_CODE}}</div>
            <p class="discount">新用户可以在下单时使用一次优惠码 <strong>NEW9</strong> 获得9折优惠</p>
        </div>
    </div>
</body>
</html>
`

func TestDoSendQQ(t *testing.T) {
	err := DoSend("1252068782@qq.com", "【注册验证码】WPP 网站注册", register,
		"zzhang_xz@163.com", "zack", secretNet)
	if err != nil {
		panic(err)
	}
}

var content = `
Deer [Influencer's Name],

We're Furrjoi, a specialist in crafting high-quality silicone products designed to celebrate individuality and deliver an extraordinary experience. We always focus on providing a sense of envelopment and comfortable stimulation.

Exciting news! We're currently rolling out a range of innovative items on www.Furrjoi.com—from Muscle Suits to Monster Mask, and we're just getting started! 

Been loving your vibe, especially the media art you create. We'd love for you to explore our offerings and join us in spreading the love. For every product you promote, we're thrilled to offer a sweet 5% commission plus some awesome gear!

Interested in joining forces? Let's chat!

Cheers,
Furrjoi
`

func TestDoSendNet(t *testing.T) {

	err2 := DoSend("zzhang_xz@163.com", "【注册验证码】WPP 网站注册", content,
		"1252068782@qq.com", "zack", secretQq)
	if err2 != nil {
		panic(err2)
	}
}

func TestDoSendGmail(t *testing.T) {
	err3 := DoSend("zzhang_xz@163.com", "123", "email.Bod",
		"danielzhang182@gmail.com", "dylan", secretGmail)
	if err3 != nil {
		panic(err3)
	}
}

func TestIsHTML(t *testing.T) {
	// 示例字符串
	str1 := "<html><body><h1>Hello, World!</h1></body></html>"
	str2 := "This is a plain text string."

	// 检查字符串是否是 HTML
	fmt.Printf("Is str1 HTML? %v\n", isHTML(str1)) // true
	fmt.Printf("Is str2 HTML? %v\n", isHTML(str2)) // false
}
