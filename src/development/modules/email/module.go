package nil

import (
	timeid "development/modules/TimeID"
	//id "development/modules/UUID"
	i "development/modules/other"
	"fmt"
	c "github.com/skilstak/go/colors"
	s "github.com/whitman-colm/go-1/utils/other"
	"net/smtp"
	usr "os/user"
	"strings"
)

const Version string = c.V + "V β-S-4.4.0"

func Startup() {
	done := false
	fmt.Println(c.CL)
	fmt.Println(c.B3+"Support By Ticket S3 module by @whitman-colm", Version)
	done = false
	/////declaring username
	userStruct, _ := usr.Current()
	user := userStruct.Username
	/////Declaring priority
	fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
	s.Spacer(2)
	fmt.Println(c.B1 + "PLEASE be honest, what is the priority of the ticket?")
	fmt.Println()
	fmt.Println(c.B01 + "Low")
	fmt.Println(c.B00 + "Medium")
	fmt.Println(c.B01 + "High")
	fmt.Println(c.B00 + "Urgent")
	/////
	for done == false {
		priority, _ := i.Input("email", "sss")
		priority = strings.ToLower(priority)
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
			fmt.Println(c.X + "That's not valid, sorry")
		}
	}
}

func nonUrgent(user string, priority string) {
	done := false
	lang := ""
	for lang == "" {

		fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
		s.Spacer(2)
		fmt.Println(c.B1 + "What language or software does your problem relate to? (GOlang, Python3, JS, Java, etc.)")
		lang, _ = i.Input("email", "sss")
	}
	/////
	details := ""
	for details == "" {
		fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
		s.Spacer(2)
		fmt.Println(c.B1 + "Give a breif description of your problem")
		details, _ = i.Input("email", "sss")
	}
	/////
	fmt.Println(c.CL + c.B2 + "No shame in needing help. \nBut please fill this out and people will be able to help ASAP!")
	s.Spacer(2)
	fmt.Println(c.G + "Whats the best way someone can contact you for help on this problem?")
	fmt.Println(c.R + "{A} " + c.B01 + "e-Mail me at an email address")
	fmt.Println(c.R + "{B} " + c.B00 + "TALK to me on the skilstak.sh server @ " + user + ".")
	fmt.Println(c.R + "{C} " + c.B01 + "SLACK me (must provide username)")

	newcontact := ""
	done = false

	for done == false {
		contact, _ := i.Input("email", "sss")
		switch contact {
		case "A", "a":
			fmt.Println(c.CL + c.B1 + "What is the email address you'd like to use?")
			studentemail, _ := i.Input("email", "sss")
			newcontact = "e-mail me at " + studentemail + "."
			done = true
		case "B", "b":
			newcontact = "TALK to me on the skilstak.sh server @ " + user + "."
			done = true
		case "C", "c":
			fmt.Println(c.CL + c.B1 + "What is the SLACK ID you'd like to use?")
			studentemail, _ := i.Input("email", "sss")
			newcontact = "slack me at " + studentemail + "."
			done = true
		default:
			fmt.Println(c.X, "Sorry M8, thats not a valid statment...")
		}
	}
	//	uuid := id.LastLine("listOfUUID.txt")
	uuid := timeid.GenerateID()
	to := []string{"skilstakta@gmail.com"}
	msg := []byte("To: skilstakta@gmail.com\r\n" +
		"Subject: TICKET " + uuid + ":  Hey, I'm " + user + ". I need some help with " + lang + ". For context, it's " + priority + " priority.\r\n" +
		"\r\n" +
		"So pretty much... " + details + " The best way to contact me is to " + newcontact + "\r\n\nThis was sent using the automated TicketBot.")

	auth := smtp.PlainAuth("", "skilstakticketer@gmail.com", "obviouspassword", "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "skilstakticketer@gmail.com", to, msg)
	s.QuitAtError(err)
	s.Go(0)
	fmt.Println(c.B3 + "Sent! Help is on its way!")
	/*<<<<<<< HEAD
	  	id.WriteUUID("listOfUUID.txt", uuid)
	  =======
	  	newUuid := byte(int(uuid) + 49)
	  	id.WriteUUID("listOfUUID.txt", []byte{newUuid})
	  >>>>>>> f4b01a6e7669947e6b961cb215ced408a28ee7b2
	*/
}

func devurgent(user string) {
	//SMS & slack things needed
	fmt.Println("Still working on this, give us a while")
	fmt.Println(user)
}

func listForums() {
	fmt.Println(c.CL + c.B2 + "Hmm... here are all the forums:")
	s.Spacer(1)
	fmt.Println(c.B01 + "Bux")
	fmt.Println(c.B00 + "Camp")
	fmt.Println(c.B01 + "Ehacking")
	fmt.Println(c.B00 + "FUNdamentals")
	fmt.Println(c.B01 + "Game")
	fmt.Println(c.B00 + "Java")
	fmt.Println(c.B01 + "Js")
	fmt.Println(c.B00 + "Linux π")
	fmt.Println(c.B01 + "Code && play")
	fmt.Println(c.B00 + "Code && prep")
	fmt.Println(c.B01 + "Python")
	fmt.Println(c.B00 + "Web")
	fmt.Println(c.B01 + "GOlang")
	fmt.Println(c.B00 + "Pro")
	fmt.Println(c.B01 + "S.W.A.T.")
	s.Go(1)
}
func low() {
	fmt.Println(c.CL + c.B2 + "Ok. To prevent spam on easy items, we ask that you SLACK for help in the SLACK forums.")
	fmt.Println(c.B1 + "Is that OK?")
	fmt.Println()
	fmt.Println(c.R + "no")
	fmt.Println(c.R + "yes")
	openBrowser := ""
	done := false
	otherDone := false
	for done == false {
		openBrowser, _ = i.Input("email", "sss")
		if openBrowser == "yes" {
			fmt.Println(c.CL + c.B2 + "Ok, before we do, what does your question pretain to?")
			fmt.Println(c.B1 + "If you don't know, just type \"help\"")
			for otherDone == false {
				//For loop encases switch for definate answer
				site, _ := i.Input("email", "sss")
				site = strings.ToLower(site)
				fmt.Println(c.CL + c.B1 + "Copy and paste this into your browser")
				s.Spacer(1)
				switch site {
				case "bux":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/bux/")
					otherDone = true
					s.Go(1)
				case "camp":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/camp/")
					otherDone = true
					s.Go(1)
				case "ehacking", "e-hacking":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/ehacking")
					otherDone = true
					s.Go(1)
				case "fundamentals":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/fundamentals/")
					otherDone = true
					s.Go(1)
				case "game":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/game")
					otherDone = true
					s.Go(1)
				case "java":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/java")
					otherDone = true
					s.Go(1)
				case "js", "javascript":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/javascript")
					otherDone = true
					s.Go(1)
				case "pi", "linux pi", "π":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/linuxpi")
					otherDone = true
					s.Go(1)
				case "play", "code & play", "code + play", "code and play", "code && play":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/play")
					otherDone = true
					s.Go(1)
				case "prep", "code & prep", "code + prep", "code and prep", "code && prep":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/prep")
					otherDone = true
					s.Go(1)
				case "py", "python":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/python/")
					otherDone = true
					s.Go(1)
				case "web":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/web")
					otherDone = true
					s.Go(1)
				case "golang", "go":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/golang/")
					otherDone = true
					s.Go(1)
				case "pro":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/pro")
					otherDone = true
					s.Go(1)
				case "swat", "s.w.a.t.", "s.w.a.t":
					fmt.Println(c.O)
					fmt.Println("https://skilstak.slack.com/messages/swat")
					otherDone = true
					s.Go(1)
				default:
					listForums()
				}
				//the switch statment is now done
			}
			s.Bye()
			//the encasing for loop is done
			done = true
		} else /*Stupid long first if is done*/ if openBrowser == "no" {
			fmt.Println(c.B2 + "Ok then. Sorry we were not of any help...")
			s.Bye()
		} else {
			fmt.Println(c.B2 + "Yeah no... I don't know what you just said")
		}
	}
	//for loop done
} //function done
