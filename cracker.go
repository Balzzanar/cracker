package main

import (
    "github.com/apsdehal/go-logger"
    "os"
    "time"
    "fmt"
    "strings"
    "io/ioutil"
//    "flag"
//    "errors"
)

/* -- Constants -- */
const LEDOFF 		= 0
const LEDON 		= 1
const LEDBLINK 		= 2
const DIRWORDLIST	= "wordlist/"
const DIRWPA		= "tocrack/"

type Led struct {
	Name string
	Port string
	State int
	QueuedState int
}

var log 	*logger.Logger
var dbh 	*DBHandler
var ledList	[]*Led
var jh      *JohnHandler


/**
 * Starts up the main thread.
 * 
 * @name main
 */
func main () {
	ledList = []*Led{}
    var err error
    log, err = logger.New("cracker", 1, os.Stdout)
    if err != nil {
        panic(err) 
    }
    dbh = new(DBHandler)
    dbh.Init()
    defer dbh.Close()
    jh = new(JohnHandler)
    jh.Init()

    /* Scan the given folders */
 //   dbh.StoreWordlist(&Wordlist{id:0, name:"rockyou.txt", size:"143MB", avg_run:30012313})
 //   dbh.GetAllWpa()
 //   ScanUpdate()

    /* Adding Leds to a list */
    led := &Led{Name: "internet_access", Port: "GPIO14", State: LEDOFF, QueuedState: LEDBLINK}
    ledList = append(ledList, led)

    /* !TEST!  LED code  !TEST! */
 //    go LEDController()
	// log.Info("Sleeping...")
	// time.Sleep(10 * time.Second)
 //    log.Info("Changing Led-State to ON")
 //    led.QueuedState = LEDON
 //    time.Sleep(4 * time.Second)
}

/**
 * 
 * @name JohnController
 */
func JohnController() {
    for true {

    }
}



/**
 * Controls the LEDs, that will indicate the status of the cracker.
 *      To change a LED, set 'QueuedState' to the wanted state.
 * 
 * @name LEDController
 */
func LEDController() {
	for true {
		for _,led := range ledList {
			if led.QueuedState != led.State {
	        	led.State = led.QueuedState
			    switch led.State {
			        case LEDBLINK:
			            go LEDLight(LEDBLINK, led)
			        case LEDON: 
			        	LEDLight(LEDON, led)
			        default:
			            LEDLight(LEDOFF, led)
			    }
			}
		} 
		time.Sleep(100 * time.Millisecond)
	}
}

/**
 * Makes the LED actually blink by controlling the GPIO ports.
 * 
 * @name LEDLight
 * @param State
 * @param led
 */
func LEDLight(State int, led *Led) {
	if State == LEDON {
		// Leave the LED on.
        log.Info("Led set to ON!")
		return
	}
	if State == LEDOFF {
		// Leave the LED off.
        log.Info("Led set to OFF!")
		return 
	}
	for true {
		if led.State != LEDBLINK {
			return 
		}
		log.Info("Blinking!")
		time.Sleep(1 * time.Second)
	}
}


/**
 * Scans and updates the database file based on what 
 * files are found.
 * 
 * @name ScanUpdate
 */
func ScanUpdate() {
    var err error

    flist := scanDir(DIRWPA)
    for _,file := range flist {
        if strings.Contains(file.Name(), ".cap") {
            var bssid []byte
            bssidfile := DIRWPA + strings.Split(file.Name(), ".")[0] + ".bssid"
            if bssid, err = ioutil.ReadFile(bssidfile); err != nil {
                log.Error(fmt.Sprintf("Missing file (%s), will ignore: %s", bssidfile, file.Name()))
                continue
            }
            log.Info(fmt.Sprintf("Found bssid: %s", bssid))
            dbh.StoreWpa(&Wpa{name:file.Name(), bssid:string(bssid)})
        }
    }

	flist = scanDir(DIRWORDLIST)
    for _,file := range flist {
        dbh.StoreWordlist(&Wordlist{name:file.Name(), size:string(file.Size()), avg_run:0})
    }
}


/**
 * Scans a given directory and returns a list of 
 * all the files.
 * 
 * @name scanDir
 */
func scanDir(dir string) []os.FileInfo {
    files := []os.FileInfo{}
    d, err := os.Open(dir)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()
    fi, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi := range fi {
        if fi.Mode().IsRegular() {
            files = append(files, fi)
        }
    }
    return files
}
