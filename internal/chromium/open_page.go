package chromium

import (
	"log"
	"os/exec"
)

func OpenPage(url string) {
	cmd := exec.Command("chromium", url)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}