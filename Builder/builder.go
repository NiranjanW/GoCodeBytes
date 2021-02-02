package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

//user defined function -passed around like any other value
type action func(builder *EmailBuilder)

type EmailBuilder struct {
	email email
}

func sendEmail(email *email) {
	//logic to send email
	fmt.Println("emailsent....")
	fmt.Printf("To:%s ,\nSubject:%s\n%s\n Regards,\n%s\n", email.to, email.subject, email.body, email.from)

}

//SendEmail function is for client to send email bulder
func SendEmail(action action) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmail(&builder.email)
}

func (eb *EmailBuilder) To(value string) *EmailBuilder {
	//basic validation
	if !strings.Contains(value, "@") {
		panic("Invalid email")
	}
	eb.email.to = value

	return eb
}

//From sets the email's "From" address
func (eb *EmailBuilder) From(value string) *EmailBuilder {
	//basic validation
	if !strings.Contains(value, "@") {
		panic("Invalid email")
	}
	eb.email.from = value
	return eb
}

//Body sets the email's "Body"
func (eb *EmailBuilder) Body(value string) *EmailBuilder {
	eb.email.body = value
	return eb
}

//Subject sets the email's "Subject"
func (eb *EmailBuilder) Subject(value string) *EmailBuilder {
	eb.email.subject = value
	return eb
}

func RunBuilderParameter() {
	SendEmail(func(b *EmailBuilder) {
		b.
			To("World@gmail.com").
			From("me@gmail.com").
			Subject("Hello world").
			Body("Sample body for a dummy email")
	})

}
