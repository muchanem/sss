package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	s "github.com/whitman-colm/go-1/utils/other"
	"net/smtp"
	"strings"
)

func Startup() {
	fmt.Println(c.CL)
	fmt.Println(c.B3, "Support By Ticket sss module by @whitman-colm & @donovank", c.V, "V 1.0.0")
	done := false
	/////declaring username
	fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
	s.Spacer(2)
	fmt.Println(c.G + "What is your skilstak.sh username?")
	user, _ := i.Prompt(c.B + ">" + c.M)
	/////Declaring priority
	fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
	s.Spacer(2)
	fmt.Println(c.G + "PLEASE be honest, what is the priority of the ticket?")
	fmt.Println()
	fmt.Println(c.Y + "Low")
	fmt.Println(c.O + "Medium")
	fmt.Println(c.R + "High")
	fmt.Println(c.V + "Urgent")
	priority, _ := i.Prompt(c.B + ">" + c.M)
	/////
	priority = strings.ToLower(priority)
	for done == false {
		switch priority {
		case "low", "medium", "high":
			nonUrgent()
			done = true
		case "urgent":
			urgent()
			done = true
		default:
			fmt.Println(c.B2 + "That's not valid, sorry")
		}
	}
}

func nonUrgent(user string) {
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
	fmt.Println(c.R + "{ A } " + c.B1 + "E-Mail me at an email address")
	fmt.Println(c.R + "{ B } " + c.B1 + "TALK to me on the skilstak.sh server @ " + user + ".")
	fmt.Println(c.R + "{ C } " + c.B1 + "SLACK me (must provide username)")

	newcontact := ""
	done = false

	for done == false {
		contact, _ := i.Prompt(c.B + ">" + c.M)
		switch contact {
		case "A", "a":
			studentemail, _ := i.Prompt(c.CL + c.G + "What is the email address you'd like to use?\n " + c.B + ">" + c.M)
			newcontact = "e-mail me at " + studentemail + "."
			done = true
		case "B", "b":
			newcontact = "TALK to me on the skilstak.sh server @ " + user + "."
			done = true
		case "C", "c":
			studentemail, _ := i.Prompt(c.CL + c.G + "What is your SLACK ID?\n" + c.B + "> " + c.M)
			newcontact = "slack me at " + studentemail + "."
		default:
			fmt.Println(c.B2, "Sorry M9, thats not a valid statment...")
		}
	}
	to := []string{"skilstakta@gmail.com"}
	msg := []byte("To: skilstakta@gmail.com\r\n" +
		//TODO: add ticket ID to subject line before "hey, I'm user"
		"Subject:TICKET #(id).  Hey, I'm " + user + ". I need some help with " + lang + ". For context, it's " + priority + " priority.\r\n" +
		"\r\n" +
		"So pretty much... " + details + "\r\nThe best way to contact me is to " + newcontact + "\r\n\nThis was sent using the automated TicketBot.")

	auth := smtp.PlainAuth("", "skilstakticketer@gmail.com", "obviouspassword", "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "skilstakticketer@gmail.com", to, msg)
	s.QuitAtError(err)
	s.Go(0)
	fmt.Println(c.B3 + "Sent! Help is on it's way!")
	done = true
}

func urgent(user string) {
	//SMS & slack things needed
	fmt.Println("Still working on this, give us a while")
}
