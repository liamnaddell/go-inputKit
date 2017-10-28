package pk

import "bufio"
import "fmt"
import "os"
import "strconv"

var exits = []string{
	"exit",
	"Exit",
	"Quit",
	"Leave",
	"bye",
	"Bye",
	"quit",
	"leave",
}

var exit bool

type Data map[string]string
type Rootitem struct {
	Lower      []Menuitem
	NextTitle  string
	Data       Data
	Skippable  bool
	Oncomplete func()
}
type Menuitem struct {
	Lower      []Menuitem
	Prompt     string
	Oncomplete func(*Data)
	Type       string
	Name       string
	NextTitle  string
	DataName   string
}

/*
func testusage() {
	var ri = Rootitem{
		Data:      make(map[string]string),
		NextTitle: "Which submenu do u want",
		Lower: []Menuitem{
			{
				Prompt: "Pick me from the menu using numbers",
				Oncomplete: func(d *Data) {
					fmt.Println("Just cuz you pick numbers, does not mean you know them -Albert Winestien")
				},
				NextTitle: "How many numbers do you want?",
				Lower: []Menuitem{
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
				Lower: []Menuitem{
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
				Oncomplete: func(d *Data) {
					fmt.Println("You will never be cool lol!")
				},
			},
			{
				Prompt:    "Click here for free youtube red subscriptions",
				NextTitle: "How long do you want it to last for?",
				Lower: []Menuitem{
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
				Oncomplete: func(d *Data) {
					fmt.Println("You now have a virus on your pc lol")
				},
			},
			{
				Prompt: "Entering Data test",
				Lower: []Menuitem{
					{
						Prompt:   "Please enter the number 42",
						Type:     "UserData",
						DataName: "numberone",
						Oncomplete: func(Data *Data) {
							if (*Data)["numberone"] == "42" {
								fmt.Println("Thank you good sir")
							} else {
								fmt.Println("Boo you suck")
							}
						},
					},
					{
						Prompt:   "Please enter the number 12",
						Type:     "UserData",
						DataName: "numbertwo",
						Oncomplete: func(Data *Data) {
							if (*Data)["numbertwo"] == "12" {
								fmt.Println("Good job, you can see and press 12")
							} else {
								fmt.Println("Really, do better man")
							}
						},
					},
				},
			},
		},
	}
	litter.Dump(ri)
	fmt.Println("")
	RootTrav(ri)
}
*/
func RootTrav(r Rootitem) {
	fmt.Println(r.NextTitle)
	Menutrav(r.Lower, &r.Data, r.Skippable)
}
func Menutrav(ra []Menuitem, Data *Data, skippable bool) {
	if exit {
		return
	}
	var userDataprompt bool
	for i := 0; i < len(ra); i++ {
		if ra[i].Type == "" {
			userDataprompt = false
			fmt.Println(fmt.Sprintf("[%d] %s", i, ra[i].Prompt))
		} else if ra[i].Type == "UserData" {
			userDataprompt = true
			fmt.Println(ra[i].Prompt)
			userdata := input(">")
			if skippable {
				for i := 0; i < len(exits); i++ {
					if userdata == exits[i] {
						exit = true
						return
					}
				}
			}
			fmt.Println(userdata)
			(*Data)[ra[i].DataName] = userdata
			if ra[i].Oncomplete != nil && !exit {
				ra[i].Oncomplete(Data)
			}
		}
	}
	if userDataprompt == true {
		return
	}
	var resi = -1
	for {
		res := input(">")
		if skippable {
			for i := 0; i < len(exits); i++ {
				if res == exits[i] {
					exit = true
					return
				}
			}
		}
		resi = toint(res)
		if resi > -1 {
			if resi < len(ra) {
				break
			}
		}
		fmt.Println("Please enter a proper value")
	}
	if ra[resi].Lower != nil {
		fmt.Println(ra[resi].NextTitle)
		Menutrav(ra[resi].Lower, Data, skippable)
	}
	if ra[resi].Oncomplete != nil && !exit {
		ra[resi].Oncomplete(Data)
	}

}
func toint(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
func input(str string) string {
	fmt.Print("\n" + str)
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	return r.Text()

}

/*func menutrav(ra []menuitem) {
	for i := 0; i < len(ra); i++ {
		fmt.Println(ra[i].Name)
		menutrav(ra[i].Lower)
	}
}
*/
