package main

import (
	"os/exec"
	"fmt"
)

/* -- Constants -- */



type JohnHandler struct {}


/**
 * 
 * @name Init
 */
func (this *JohnHandler) Init() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(fmt.Sprint(err))
	}
	log.Info(fmt.Sprintf("The date is %s\n", out))
	log.Info("JohnHandler loaded!")
}

/**
 * 
 * @name Status
 */
func (this *JohnHandler) Start(wpa *Wpa) *Run {
	return &Run{id_wpa:0, id_wordlist:1, result:"Running", time:0, started:1234567, session:"adfdfsd-sdfsdf88"}
}

