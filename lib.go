package ik

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

func RootTrav(r Rootitem) error {
	fmt.Println(r.NextTitle)
	return Menutrav(r.Lower, &r.Data, r.Skippable)
}
func Menutrav(ra []Menuitem, Data *Data, skippable bool) error {
	if exit {
		return nil
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
						return nil
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
		return nil
	}
	var resi = -1
	for {
		res := input(">")
		if skippable {
			for i := 0; i < len(exits); i++ {
				if res == exits[i] {
					exit = true
					return nil
				}
			}
		}
		var c bool
		for {
			var err error
			resi, err = toint(res)
			if err != nil {
				fmt.Println("Please input a valid integer!")
				c = true
				break
			} else {
				c = false
				break
			}
		}
		if c {
			continue
		}
		if resi > -1 {
			if resi < len(ra) {
				break
			}
		}
		fmt.Println("Please enter a proper value")
	}
	if ra[resi].Lower != nil {
		fmt.Println(ra[resi].NextTitle)
		err := Menutrav(ra[resi].Lower, Data, skippable)
		if err != nil {
			return err
		}
	}
	if ra[resi].Oncomplete != nil && !exit {
		ra[resi].Oncomplete(Data)
	}
	return nil
}
func toint(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func input(str string) string {
	fmt.Print("\n" + str)
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	return r.Text()
}
