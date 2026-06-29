package main
import "fmt"

func main() {	
   var conferenceName = "Go conference"
    const conferenceTickets = 50
	var availableTickets = 20 
   
   fmt.Print("Welcome to ", conferenceName, " the  booking application ")
    
   fmt.Print("Get your tickets to attend")
   
   fmt.Print("We have a total of ", conferenceTickets, " tickets and ", availableTickets, " are still available.")
   
}