package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	s "github.com/whitman-colm/go-1/utils/other"
	//w "github.com/skratchdot/open-golang/open"
	"net/smtp"
	"os"
	"strings"
)

func main() {
	fmt.Println(c.CL)
	fmt.Println(c.B3, "SkilStak Script Support by @whitman-colm & @donovank", c.V, "V S-1.0.0")
	fmt.Println(c.B2 + "Hey there, need some help?")
	s.Spacer(3)
	fmt.Println(c.G + "Select one of the keywords in red")
	fmt.Println(c.R+"{A}", c.G, "No, I'm fine.")
	fmt.Println(c.R+"{B}", c.G, "Yeah, I need help")

	done := false

	for done == false {
		function, _ := i.Prompt(c.B + "> " + c.M)
		mode := strings.ToLower(function)
		switch mode {
		case "a":
			fmt.Println(c.B2, "Ok! sorry for botherin' ya.")
			os.Exit(-1)
			done = true
		case "b":
			s.Go(1)
			done = true
		default:
			fmt.Println(c.B2, "Sorry M9, thats not a valid statment...")
		}
	}
	fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
	s.Spacer(2)
	fmt.Println(c.G, "What is your skilstak.sh username?")
	user, _ := i.Prompt(c.B + ">" + c.M)
	/////
	fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
	s.Spacer(2)
	fmt.Println(c.G + "What priority is it? (select ONE)")
	s.Spacer(2)
	fmt.Println(c.Y + "Low\n" + c.O + "Medium\n" + c.R + "High\n" + c.B3 + "Urgent")
	oldpriority, _ := i.Prompt(c.B + "> " + c.M)
	priority := strings.ToLower(oldpriority)
	/////
	done = false
	//done has been defined before, reset here.
	for done == false {
		//Sorry about the unnesisary, switch statment, It's for V 2 will extend this later
		switch priority {
		case "low", "medium", "high", "urgent":
			fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
			s.Spacer(2)
			fmt.Println(c.G + "What language or software does your problem relate to?")
			lang, _ := i.Prompt(c.B + "> " + c.M)
			/////
			fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
			s.Spacer(2)
			fmt.Println(c.G + "Give a breif description of your problem")
			details, _ := i.Prompt(c.B + "> " + c.M)
			/////
			fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
			s.Spacer(2)
			fmt.Println(c.G + "Whats the best way someone can contact you for help on this problem?")
			fmt.Println(c.R + " { A } " + c.B1 + "E-Mail me at an email address")
			fmt.Println(c.R + " { B } " + c.B1 + "TALK to me on the skilstak.sh server @ " + user + ".")

			newcontact := ""
			done = false

			for done == false {
				contact, _ := i.Prompt(c.B + ">" + c.M)
				switch contact {
				case "A", "a":
					studentemail, _ := i.Prompt(c.G + "What is the email address you'd like to use?\n " + c.B + ">" + c.M)
					newcontact = "e-mail me at " + studentemail + "."
					done = true
				case "B", "b":
					newcontact = "TALK to me on the skilstak.sh server @ " + user + "."
					done = true
				default:
					fmt.Println(c.B2, "Sorry M9, thats not a valid statment...")
				}
			}
			to := []string{"skilstakta@gmail.com"}
			msg := []byte("To: skilstakta@gmail.com\r\n" +
				//TODO: add ticket ID to subject line before "hey, I'm user"
				"Subject: Hey, I'm " + user + ". I need some help with " + lang + ". For context, it's " + priority + " priority.\r\n" +
				"\r\n" +
				"So pretty much... " + details + "\r\nThe best way to contact me is to " + newcontact + "\r\n\nThis was sent using the automated TicketBot.")

			auth := smtp.PlainAuth("", "skilstakticketer@gmail.com", "obviouspassword", "smtp.gmail.com")

			err := smtp.SendMail("smtp.gmail.com:587", auth, "skilstakticketer@gmail.com", to, msg)
			s.QuitAtError(err)
			s.Go(0)
			fmt.Println(c.B3 + "Sent! Help is on it's way!")
			done = true

			//That was a mouthful
		default:
			fmt.Println(c.B3 + "Thats not a valid argument...")
			//function, _ := i.Prompt(c.B + ">" + c.M)
			//mode := strings.ToLower(oldpriority)
		}
	}
}
