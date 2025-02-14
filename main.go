/**********************
Lecture 2 Code
Author: Tessah Belcher
Date Completed:
Description:
**********************/

package main

import "fmt"

func main() {
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	
    var randomNumber1 int = rand.Intn(10) // generates and stores random int ranging from 0 to 9
    var randomNumber2 int = rand.Intn(20) // generates and stores random int ranging from 0 to 19
    var randomNumber3 int = rand.Intn(31) // generates and stores random int ranging from 0 to 30

    fmt.Println("Random Number 1:", randomNumber1)
    fmt.Println("Random Number 2:", randomNumber2)
    fmt.Println("Random Number 3:", randomNumber3)
}
}
