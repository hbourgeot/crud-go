package utilities

import "fmt"

func PrintTitle(title string) {
	for i := 0; i < len(title)+4; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
	fmt.Printf("* %s *\n", title)
	for i := 0; i < len(title)+4; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")

}

func PrintSubtitle(subtitle string) {
	fmt.Println(subtitle)
	for i := 0; i < len(subtitle)+2; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n\n")
}

func PrintMenus(options ...string) {
	var count int
	for i := 0; i < len(options); i++ {
		fmt.Printf("%d. %s\n", i, options[i])
		if count < len(options[i]) {
			count = len(options[i])
		}
	}
	if count < 18 {
		count = 18
	} else {
		count += 5
	}
	for i := 0; i < count+2; i++ {
		fmt.Print("*")
	}
	fmt.Print("\nEnter your option:")
}