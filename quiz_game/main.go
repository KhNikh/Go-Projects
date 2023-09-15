package main

import (
	"fmt"
	"math/rand"
	"time"
	// "time"
)

func main(){
    var score int
    l := make(chan int)
    main:
    for {
        num1 := rand.Intn(101)
        num2 := rand.Intn(101)
        fmt.Printf("%d + %d = ", num1,num2)
        go func(){
            var input int 
            fmt.Scanln(&input)
            l <- input
        }()

        select{
            case userInput := <- l:
                if userInput != num1+num2 {
                    fmt.Println("Score: ", score)
                    break main
                }else{
                    score++
                }
            case <- time.After(5*time.Second):
                fmt.Println("\nTime out")
                fmt.Println("Score: ",score)
                break main
        }
    }
    
}
