package main

import (
	//"fmt"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("welcome to my test program. please tell me your name and age (if you dont tell your age or name you will cannot use if you say quit it will go out)")
	var name string
	var number int
	fmt.Scanf("%s %d", &name, &number)
	if number < 10 {
		fmt.Println("you dont have permisson to use")
	} else if number >= 10 && number < 80 {
		fmt.Println("ok,", name, " can use this")
		fmt.Println("password please")
		var ABC string
		fmt.Scanln(&ABC)
		if ABC == "Boon" && (name == "Bui" || name == "Boo"  || name == "BoonBoon"){fmt.Println("wait")
			_, err := exec.Command("Steam.exe").Output()
			if err != nil {
				fmt.Errorf("error; %s", err)
			}
		} else {
			fmt.Println("password is wrong,you are not authorize to use")
		}
	} else {
		fmt.Println("you are too old to use this")
	}

}
