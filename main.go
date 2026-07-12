package main
import "fmt"

func main() {	
   var conferenceName = "Go conference"
    const conferenceTickets = 50
	var availableTickets uint= 20 
   
   fmt.Printf("Welcome to %v the  booking application\n", conferenceName)
    
   fmt.Println("Get your tickets to attend")

   fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, availableTickets)
     
   var firstname string
   fmt.Println("Enter your firstname:")
   fmt.Scan(&firstname)
//    fmt.Printf("the firstname is %v\n", firstname)

   var lastname string
   fmt.Println("Enter your lastname:")
   fmt.Scan(&lastname)
//    fmt.Printf("the lastname is %v\n", lastname)

   var userTickets uint
   fmt.Println("Enter amount of tickets you want to book:")
   fmt.Scan(&userTickets)
//    fmt.Printf("the number of tickets you want to book is %v\n", userTickets)
   remainingTickets := availableTickets - userTickets
  
  fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email shortly.\n", firstname, lastname, userTickets)
  fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

  
  var bookings = [50]string{}
  bookings[0] =  firstname + " " + lastname

}
