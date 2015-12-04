package nil

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/whitman-colm/go-1/utils/input"
	s "github.com/whitman-colm/go-1/utils/other"
	"net/smtp"
	"strings"
)

Version := c.V+"V β-S-4.1.4"

func Startup() {
	done := false
	fmt.Println(c.CL)
	fmt.Println(c.B3+"Support By Ticket sss module by @whitman-colm & @donovank", Version)
	done = false
	/////declaring username
	user := ""
	for user == "" {
		fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
		s.Spacer(2)
		fmt.Println(c.G + "What is your skilstak.sh username?")
		user, _ = i.Prompt(c.B + ">" + c.M)
	}
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
		case "low":
			low()
		case "urgent", "medium", "high":
			nonUrgent(user, priority)
			done = true
		case "devurgent": /*Still doing this, give us a while*/
			devurgent(user)
			done = true
		default:
			fmt.Println(c.B2 + "That's not valid, sorry")
		}
	}
}

func nonUrgent(user string, priority string) {
	done := false
	lang := ""
	for lang == "" {

		fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
		s.Spacer(2)
		fmt.Println(c.G + "What language or software does your problem relate to? (GOlang, Python3, JS, Java, etc.)")
		lang, _ = i.Prompt(c.B + "> " + c.M)
	}
	/////
	details := ""
	for details == "" {
		fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
		s.Spacer(2)
		fmt.Println(c.G + "Give a breif description of your problem")
		details, _ = i.Prompt(c.B + "> " + c.M)
	}
	/////
	fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
	s.Spacer(2)
	fmt.Println(c.G + "Whats the best way someone can contact you for help on this problem?")
	fmt.Println(c.R + "{ A } " + c.B1 + "e-Mail me at an email address")
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
			done = true
		default:
			fmt.Println(c.B2, "Sorry M9, thats not a valid statment...")
		}
	}
	to := []string{"skilstakta@gmail.com"}
	msg := []byte("To: skilstakta@gmail.com\r\n" +
		//TODO: add ticket ID to subject line before "hey, I'm user"
		"Subject: TICKET #(id).  Hey, I'm " + user + ". I need some help with " + lang + ". For context, it's " + priority + " priority.\r\n" +
		"\r\n" +
		"So pretty much... " + details + "\r\nThe best way to contact me is to " + newcontact + "\r\n\nThis was sent using the automated TicketBot.")

	auth := smtp.PlainAuth("", "skilstakticketer@gmail.com", "obviouspassword", "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "skilstakticketer@gmail.com", to, msg)
	s.QuitAtError(err)
	s.Go(0)
	fmt.Println(c.B3 + "Sent! Help is on its way!")
}

func devurgent(user string) {
	//SMS & slack things needed
	fmt.Println("Still working on this, give us a while")
	fmt.Println(user)
}

func listForums() {
	fmt.Println(c.CL + c.B2 + "Hmm... here are all the forums:")
	fmt.Println(c.G + "Bux")
	fmt.Println(c.G + "Camp")
	fmt.Println(c.G + "Ehacking")
	fmt.Println(c.G + "FUNdimentals")
	fmt.Println(c.G + "Game")
	fmt.Println(c.G + "Java")
	fmt.Println(c.G + "Js")
	fmt.Println(c.G + "Linux π")
	fmt.Println(c.G + "Code && play")
	fmt.Println(c.G + "Code && prep")
	fmt.Println(c.G + "Python")
	fmt.Println(c.G + "Web")
	fmt.Println(c.G + "GOlang")
	fmt.Println(c.G + "Pro")
	fmt.Println(c.G + "S.W.A.T.")
	s.Go(1)
}
func low() {
	fmt.Println(c.CL + c.B2 + "Ok. To prevent spam on easy items, we ask that you SLACK for help in the SLACK forums.")
	fmt.Println(c.B1 + "Is that OK?")
	fmt.Println()
	fmt.Println(c.R + "no")
	fmt.Println(c.G + "yes")
	openBrowser := ""
	done := false
	for done == false {
		openBrowser, _ = i.Prompt(c.B + "> " + c.M)
		if openBrowser == "yes" {
			fmt.Println(c.CL + c.B1 + "Ok, before we do, what does your question pretain to?")
			fmt.Println("If you don't know, just type \"help\"")
			for done == false {
				site, _ := i.Prompt(c.B + "> " + c.M)
				site = strings.ToLower(site)
				fmt.Println(c.CL + c.B1 + "Copy and paste this into your browser")
				switch site {
				case "bux":
					fmt.Println("https://skilstak.slack.com/messages/bux/")
					done = true
				case "camp":
					fmt.Println("https://skilstak.slack.com/messages/camp/")
					done = true
				case "ehacking", "e-hacking":
					fmt.Println("https://skilstak.slack.com/messages/ehacking")
					done = true
				case "fundimentals":
					fmt.Println("https://skilstak.slack.com/messages/fundamentals/")
					done = true
				case "game":
					fmt.Println("https://skilstak.slack.com/messages/game")
					done = true
				case "java":
					fmt.Println("https://skilstak.slack.com/messages/java")
					done = true
				case "js", "javascript":
					fmt.Println("https://skilstak.slack.com/messages/javascript")
					done = true
				case "pi", "linux pi", "π":
					fmt.Println("https://skilstak.slack.com/messages/linuxpi")
					done = true
				case "play", "code & play", "code + play", "code and play", "code && play":
					fmt.Println("https://skilstak.slack.com/messages/play")
					done = true
				case "prep", "code & prep", "code + prep", "code and prep", "code && prep":
					fmt.Println("https://skilstak.slack.com/messages/prep")
					done = true
				case "py", "python":
					fmt.Println("https://skilstak.slack.com/messages/python/")
					done = true
				case "web":
					fmt.Println("https://skilstak.slack.com/messages/web")
					done = true
				case "golang", "go":
					fmt.Println("https://skilstak.slack.com/messages/golang/")
					done = true
				case "pro":
					fmt.Println("https://skilstak.slack.com/messages/pro")
					done = true
				case "swat", "s.w.a.t.", "s.w.a.t":
					fmt.Println("https://skilstak.slack.com/messages/swat")
					done = true
				default:
					listForums()
				}
				//the switch statment is now done
			}
			//the encasing for loop is done
		} else if openBrowser == "no" {
			fmt.Println(c.B2 + "Ok then. Sorry we were not of any help...")
			done = true
		} else {
			fmt.Println(c.B2 + "Yeah no... I don't know what you just said")
		}
		done = true
	}
}
