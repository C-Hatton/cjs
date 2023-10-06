package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// returns the current date and time, as a string, in the format dd/mm/yyyy hh:mm:ss
func getCurrentTime() (currentDate string, currentTime string) {

	timeNow := time.Now() //gets the current date time

	//format the time:
	//Get date:
	day := strconv.Itoa(timeNow.Day())
	month := strconv.Itoa(int(timeNow.Month()))
	year := strconv.Itoa(timeNow.Year())
	currentDate = day + "/" + month + "/" + year

	//Get time:
	hour := strconv.Itoa(timeNow.Hour())
	minute := strconv.Itoa(timeNow.Minute())
	second := strconv.Itoa(timeNow.Second())
	currentTime = hour + ":" + minute + ":" + second

	return
}

func log(page string) (result bool) {

	result = false
	currentDate, currentTime := getCurrentTime()
	data := currentDate + "\n" + currentTime + "\n" + page + "\n"
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		result = false
	}
	if _, err := f.Write([]byte(data)); err != nil {
		result = false
	} else {
		result = true
	}
	if err := f.Close(); err != nil {
		result = false
	}

	return
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	log("/")
	io.WriteString(w, "This is my website!\n")
}

func main() {

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
