package email

import (
	"github.com/jasonlvhit/gocron"
	"log"
)

func task() {
	log.Println("Running SendEmail Every(5).Second()")
}

func SendEmail() {
	gocron.Every(5).Second().Do(task)
}
