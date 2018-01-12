package main

import (
    "log"
    "net/smtp"
)

func main() {
    // Set up authentication information.
    auth := smtp.PlainAuth(
        "",
        "threatbookt@126.com",
        "threatbook12345",
        "smtp.126.com",
    )
    // Connect to the server, authenticate, set the sender and recipient,
    // and send the email all in one step.
    err := smtp.SendMail(
        "smtp.126.com:25",
        auth,
        "threatbookt@126.com",
        []string{"threatbookt@126.com"},
        []byte("go send This is the email body."),
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Print("send pass")
}
