package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {
	bill := admin{
		user: user{
			name:  "Bill",
			email: "bill@gmail.com",
		},
		level: "super",
	}

	sendNotification(&bill)
}

func sendNotification(n notifier) {
	n.notify()
}
