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
		log.Fatal(err)
	}
	log.Info(fmt.Sprintf("The date is %s\n", out))
	log.Info("JohnHandler loaded!")
}

