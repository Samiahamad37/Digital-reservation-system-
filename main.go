package main
import "fmt"

func main() {	
   var conferenceName = "Go conference"
    const conferenceTickets = 50
	var availableTickets = 20 
   
   fmt.Printf("Welcome to %v the  booking application\n", conferenceName)
    
   fmt.Println("Get your tickets to attend")

   fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, availableTickets)
   
}