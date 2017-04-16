package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("welcome to my test program.\n" +
		"please tell me your name and age\n" +
		"(if you dont tell your age or name you will not be able to use\n" +
		"if you say quit, the program will be close,this program made by Boon")
	var name string
	var number int
	fmt.Scanf("%s", &name)

	if name == "quit" {
		fmt.Println("this process finish")
		return
	}

	fmt.Scanf("%d", &number)

	if number < 10 {
		fmt.Println("you dont have permisson to use")
	} else if number >= 10 && number < 80 {
		fmt.Println("ok,", name, " can use this")
		fmt.Println("password please")
		var Password string
		fmt.Scanln(&Password)
		if Password == "Boon" && (name == "Bui" || name == "Boo" || name == "Boon") {
			fmt.Println("What do you want to open,Google or Steam")
			var Web string
			fmt.Scanln(&Web)

			if Web == "Steam" {
				fmt.Println("wait")
				_, err := exec.Command("Steam.exe").Output()
				if err != nil {
					fmt.Errorf("error; %s", err)
				}
				fmt.Println("ok now i open steam for you")
			}
			if Web == "steam" {
				fmt.Println("did you mean Steam")
			}

			if Web == "Google" {
				fmt.Println("wait")
				_, err := exec.Command("chrome.exe").Output()
				if err != nil {
					fmt.Errorf("error; %s", err)
				}
			}
			if Web == "google" {
				fmt.Println("did you mean Google")

			}

		} else {
			fmt.Println("password is wrong, or you are not authorize to use")
		}
	} else {
		fmt.Println("you are too old to use this")
	}
}
