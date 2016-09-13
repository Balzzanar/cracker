# Cracker
This system uses [John The Ripper][1] to perform password cracking attacks. It is designd to be able to handle sudden interupts and restarts, thus making it suitable for implmenting in a Raspberry PI.  
It also comes with a LED-system that can be installed to display the progress of the system.  
More options for John can be found [here][2]

## Flow 
---
#### Scanning

1. Scans the filesystem for wordlists given the path `DIRWORDLIST`
2. Scans the filesystem for .cap files and .bssid files given the path `DIRWPA`

#### Cracking

1. Check if there are any started session that has not finnished.
    * If so start the saved session.
    * If the start fails, this means that the run was actully done
2. Check for .cap files to crack and wordlist to use.
    * Start the new cracking by using the `--session=NAME` option when calling John
    * The sessionname shall be randomised and stored in the db, so that it can be 


#### Internet Connection
TODO


## Functions
---
#### ScanUpdate
```
ScanUpdate()
```
Updates the db with the .cap  .bssid and wordlists files that are in the given 
folders. This is done in order to prepare for crack execution.

#### LEDLight
```
LEDLight(State int, led *Led)
```
- `State` - The state of the LED, to be set
- `led` - The led to update

Controls the actual LED, by flipping the GPIO ports


[1]: http://tools.kali.org/password-attacks/john
[2]: http://www.openwall.com/john/doc/OPTIONS.shtml