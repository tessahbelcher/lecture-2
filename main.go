/**********************

Author: Tessah Belcher
Date Completed:
Description:
**********************/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type Colleague struct {
    Name  string
    Blame string
}

func main() {
    
    var colleague1 Colleague

    fmt.Println("What is your colleague's first name?")
    fmt.Scanln(&colleague1.Name)

    fmt.Println("Will you blame", colleague1.Name, "for the software bug? (0 = no, 1 = yes)")

    var response int
    fmt.Scanln(&response)

}



Note: [Colleague Name] is the name of the colleague typed in by the user at the beginning of the program.
Quick Review: Program must have
Colleague struct
Name
Blame
User input for Name on instantiated Colleague variable (i.e., variable of type "Colleague")
Random assignment of Blame for instantiated Colleague variable
User input for blaming colleague or not
Appropriate output based on potential outcome chart (above)

