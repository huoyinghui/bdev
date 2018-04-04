package controllers

import (
	"fmt"
	"testing"
)

func TestSendMail(t *testing.T) {
	recipient := "huoyinghui@apkpure.net"
	subject := "APKPure - Reset your password"
	charSet := "UTF-8"
	htmlBody := `A password reset has been requested for the APKPure username: xxx

If you do actually want to reset your password, visit this link:

https://apkpure.com/account/resetpassword?k=xxx

Thank you for using the site!

https://apkpure.com`
	cc := []string{
		"285020853@qq.com", "hyhlinux@163.com",
	}

	//ret, err := SendMail("", recipient, subject, charSet, "", htmlBody, []*string{})
	ret, err := sendMail("", recipient, subject, charSet, "", htmlBody, cc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ret-gostring", ret.GoString())
	t.Log("ret-mid", fmt.Sprintf("%v", ret.MessageId))
}
