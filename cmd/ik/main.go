package main

import "fmt"
import "github.com/liamnaddell/go-inputKit"

var ri = ik.Rootitem{
	Data:      make(map[string]string),
	NextTitle: "Which submenu do u want",
	Skippable: true,
	Lower: []ik.Menuitem{
		{
			Prompt: "Pick me from the menu using numbers",
			Oncomplete: func(d *ik.Data) {
				fmt.Println("Just cuz you pick numbers, does not mean you know them -Albert Winestien")
			},
			NextTitle: "How many numbers do you want?",
			Lower: []ik.Menuitem{
				{
					Prompt: "I use lots of numbers",
				},
				{
					Prompt: "I use more numbers",
				},
				{
					Prompt: "I use a heck of a lot more numbers, pick me!",
				},
			},
		},
		{
			Prompt:    "NO PICK ME PLEASE I HAVE COOL SUBMENUS",
			NextTitle: "What do you want to be cool with?",
			Lower: []ik.Menuitem{
				{
					Prompt: "Do you want to be cool with sunglasses",
				},
				{
					Prompt: "Do you want to be cool with shadez",
				},
				{
					Prompt: "Do you want to be cool with ice",
				},
			},
			Oncomplete: func(d *ik.Data) {
				fmt.Println("You will never be cool lol!")
			},
		},
		{
			Prompt:    "Click here for free youtube red subscriptions",
			NextTitle: "How long do you want it to last for?",
			Lower: []ik.Menuitem{
				{
					Prompt: "Do you want it to last for eight months",
				},
				{
					Prompt: "Do you want it to last for four months",
				},
				{
					Prompt: "Do you want it to last for one month",
				},
			},
			Oncomplete: func(d *ik.Data) {
				fmt.Println("You now have a virus on your pc lol")
			},
		},
		{
			Prompt: "Entering Data test",
			Lower: []ik.Menuitem{
				{
					Prompt:   "Please enter the number 42",
					Type:     "UserData",
					DataName: "numberone",
					Oncomplete: func(Data *ik.Data) {
						if (*Data)["numberone"] == "42" {
							fmt.Println("Thank you good sir")
						} else {
							fmt.Println("Nice try")
						}
					},
				},
				{
					Prompt:   "Please enter the number 12",
					Type:     "UserData",
					DataName: "numbertwo",
					Oncomplete: func(Data *ik.Data) {
						if (*Data)["numbertwo"] == "12" {
							fmt.Println("Good job, you can see and press 12")
						} else {
							fmt.Println("Really? you should do better man")
						}
					},
				},
			},
		},
	},
}

func main() {
	ik.RootTrav(ri)
}
