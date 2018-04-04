package utils

import (
	"bdev/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender    = "noreply@apkpure.com"
	Recipient = "huoyinghui@apkpure.net"
	Subject   = "APKPure - Reset your password"
	HtmlBody  = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
	CharSet  = "UTF-8"
)

var (
	conf  = &credentials.Credentials{}
	sess  = &session.Session{}
	svc   = &ses.SES{}
	debug = false
)

type stubProvider struct {
	creds   credentials.Value
	expired bool
	err     error
}

func (s *stubProvider) Retrieve() (credentials.Value, error) {
	s.expired = false
	s.creds.ProviderName = "stubProvider"
	return s.creds, s.err
}
func (s *stubProvider) IsExpired() bool {
	return s.expired
}

func init() {
	if config.AppConf.LogLevel > 7 {
		debug = true
	}
	conf = credentials.NewCredentials(&stubProvider{
		creds: credentials.Value{
			AccessKeyID:     config.AppConf.AccessKeyID,
			SecretAccessKey: config.AppConf.SecretAccessKey,
		},
	})
	//e := credentials.NewEnvCredentials()
	awsConfig, err := conf.Get()
	if err != nil {
		fmt.Errorf("err:%v", err)
		return
	}

	if debug {
		fmt.Printf("debug: %v\nenv:%v err:%v\n\n\n", debug, awsConfig, err)
	}
	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(config.AppConf.Region),
		Credentials: conf,
	})
	svc = ses.New(sess)
}

func SendMail(sender, recipient, subject, charSet, htmlBody, textBody string, cc []string) (*ses.SendEmailOutput, error) {
	if "" == sender {
		sender = Sender
	}
	if "" == charSet {
		charSet = CharSet
	}

	body := &ses.Body{}

	if "" != htmlBody {
		body.Html = &ses.Content{
			Charset: aws.String(charSet),
			Data:    aws.String(htmlBody),
		}
	}

	if "" != textBody {
		body.Text = &ses.Content{
			Charset: aws.String(charSet),
			Data:    aws.String(textBody),
		}
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			//CcAddresses: []*string{
			//},
			CcAddresses: aws.StringSlice(cc),
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: body,
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	if debug {
		fmt.Printf("input:%v\n", input.String())
	}
	// Attempt to send the email.
	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, nil
}
