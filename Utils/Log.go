package Utils

import (
	"log"
	"os"
)

var f, _ = os.OpenFile("logfile.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)

var Logger = log.New(f, "Log at: ", log.Llongfile)
