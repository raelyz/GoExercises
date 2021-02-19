package main

import (
	"fmt"
)

// ---------------------------------- VARIABLES ---------------------------------- //

var day [6]timeSlot

// var week [7][6]timeSlot
var week [7]map[string][6]timeSlot
var venues linkedList
var venueCapacities []int
var venuePrices []int

// // ----------------------------------- STRUCTS ----------------------------------- //

// type userInformation struct {
// 	firstName     string
// 	lastName      string
// 	contactNumber string
// 	email         string
// }

// type bookingStatus struct {
// 	booked bool
// 	user   userInformation
// }

// type timeSlot struct {
// 	timeSlotID   int
// 	startingTime string
// 	endingTime   string
// 	booking      bookingStatus
// }

// type venueInformation struct {
// 	venueID       int
// 	venueName     string
// 	venueType     string
// 	venueCapacity int
// 	venuePrice    float64
// 	venueSchedule [7]map[string][6]timeSlot
// }

// type node struct {
// 	venue venueInformation
// 	next  *node
// 	prev  *node
// }

// type linkedList struct {
// 	head *node
// 	tail *node
// 	size int
// }

// // ----------------------------------- METHODS ----------------------------------- //
// // ! METHODS

// // * userInformation methods

// func (u *userInformation) fillInUserInformation(firstName string, lastName string, contactNumber string, email string) error {
// 	if firstName == "" || lastName == "" || contactNumber == "" || email == "" {
// 		return errors.New("all fields must be filled in")
// 	}

// 	u.firstName = firstName
// 	u.lastName = lastName
// 	u.contactNumber = contactNumber
// 	u.email = email

// 	return nil
// }

// func (u *userInformation) updateUserInformation(updatedFirstName string, updatedLastName string, updatedContactNumber string, updatedEmail string) error {
// 	var edits int

// 	if updatedFirstName != "" {
// 		u.firstName = updatedFirstName
// 		edits++
// 	}

// 	if updatedLastName != "" {
// 		u.lastName = updatedLastName
// 		edits++
// 	}

// 	if updatedContactNumber != "" {
// 		u.contactNumber = updatedContactNumber
// 		edits++
// 	}

// 	if updatedEmail != "" {
// 		u.email = updatedEmail
// 		edits++
// 	}

// 	if edits == 0 {
// 		return errors.New("no edits made")
// 	}

// 	return nil
// }

// func (u *userInformation) deleteUserInformation() {
// 	u.firstName = ""
// 	u.lastName = ""
// 	u.contactNumber = ""
// 	u.email = ""
// }

// // func (u *userInformation) matchUserDetails(userEmail string) error {
// // 	if u.email == userEmail {
// // 		return nil
// // 	}
// // 	return errors.New("no bookings found under this email address")
// // }

// // * bookingStatus methods

// func (b *bookingStatus) makeBooking(firstName string, lastName string, contactNumber string, email string) error {
// 	error := b.user.fillInUserInformation(firstName, lastName, contactNumber, email)

// 	if error != nil {
// 		return error
// 	}

// 	b.booked = true
// 	fmt.Println(b.booked)
// 	return nil
// }

// func (b *bookingStatus) removeBooking() {
// 	b.booked = false
// 	b.user.deleteUserInformation()
// }

// // func (b *bookingStatus) retrieveBookings(userEmail string) error {
// // 	err := b.user.matchUserDetails(userEmail)

// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // * linkedList methods

// func (l *linkedList) addNode(venue venueInformation) {
// 	newNode := &node{venue, nil, nil}

// 	if l.head == nil && l.tail == nil {
// 		l.head = newNode
// 	} else {
// 		currentNode := l.head

// 		for currentNode.next != nil {
// 			currentNode = currentNode.next
// 		}

// 		currentNode.next = newNode
// 		newNode.prev = currentNode
// 	}
// 	l.tail = newNode
// 	l.size++
// }

// func (l *linkedList) printAllNodes() error {
// 	if l.head == nil && l.tail == nil {
// 		return errors.New("list empty")
// 	}

// 	currentNode := l.head

// 	for currentNode != nil {
// 		fmt.Println("Venue ID:", currentNode.venue.venueID)
// 		fmt.Println("Venue Name:", currentNode.venue.venueName)
// 		fmt.Println("Venue Type:", currentNode.venue.venueType)
// 		fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
// 		fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
// 		fmt.Println()

// 		currentNode = currentNode.next
// 	}

// 	return nil
// }

// // * sort by venue id number
// func (l *linkedList) sortVenuesByVenueID() error {
// 	if l.head == nil && l.tail == nil {
// 		return errors.New("list empty")
// 	}

// 	currentNode := l.head

// 	for currentNode != nil {
// 		indexNode := currentNode.next

// 		for indexNode != nil {
// 			if currentNode.venue.venueID > indexNode.venue.venueID {
// 				tempVenueID := currentNode.venue.venueID
// 				tempVenueName := currentNode.venue.venueName
// 				tempVenueType := currentNode.venue.venueType
// 				tempVenueCapacity := currentNode.venue.venueCapacity
// 				tempVenuePrice := currentNode.venue.venuePrice
// 				tempVenueSchedule := currentNode.venue.venueSchedule

// 				currentNode.venue.venueID = indexNode.venue.venueID
// 				currentNode.venue.venueName = indexNode.venue.venueName
// 				currentNode.venue.venueType = indexNode.venue.venueType
// 				currentNode.venue.venueCapacity = indexNode.venue.venueCapacity
// 				currentNode.venue.venuePrice = indexNode.venue.venuePrice
// 				currentNode.venue.venueSchedule = indexNode.venue.venueSchedule

// 				indexNode.venue.venueID = tempVenueID
// 				indexNode.venue.venueName = tempVenueName
// 				indexNode.venue.venueType = tempVenueType
// 				indexNode.venue.venueCapacity = tempVenueCapacity
// 				indexNode.venue.venuePrice = tempVenuePrice
// 				indexNode.venue.venueSchedule = tempVenueSchedule
// 			}
// 			indexNode = indexNode.next
// 		}
// 		currentNode = currentNode.next
// 	}
// 	return nil
// }

// // * sort by venue price
// func (l *linkedList) sortVenuesByAscendingPrice() error {
// 	if l.head == nil {
// 		return errors.New("list empty")
// 	}

// 	currentNode := l.head

// 	for currentNode != nil {
// 		indexNode := currentNode.next

// 		for indexNode != nil {
// 			if currentNode.venue.venuePrice > indexNode.venue.venuePrice {
// 				tempVenueID := currentNode.venue.venueID
// 				tempVenueName := currentNode.venue.venueName
// 				tempVenueType := currentNode.venue.venueType
// 				tempVenueCapacity := currentNode.venue.venueCapacity
// 				tempVenuePrice := currentNode.venue.venuePrice
// 				tempVenueSchedule := currentNode.venue.venueSchedule

// 				currentNode.venue.venueID = indexNode.venue.venueID
// 				currentNode.venue.venueName = indexNode.venue.venueName
// 				currentNode.venue.venueType = indexNode.venue.venueType
// 				currentNode.venue.venueCapacity = indexNode.venue.venueCapacity
// 				currentNode.venue.venuePrice = indexNode.venue.venuePrice
// 				currentNode.venue.venueSchedule = indexNode.venue.venueSchedule

// 				indexNode.venue.venueID = tempVenueID
// 				indexNode.venue.venueName = tempVenueName
// 				indexNode.venue.venueType = tempVenueType
// 				indexNode.venue.venueCapacity = tempVenueCapacity
// 				indexNode.venue.venuePrice = tempVenuePrice
// 				indexNode.venue.venueSchedule = tempVenueSchedule
// 			}
// 			indexNode = indexNode.next
// 		}
// 		currentNode = currentNode.next
// 	}
// 	return nil
// }

// func (l *linkedList) sortVenuesByDescendingPrice() error {
// 	if l.head == nil && l.tail == nil {
// 		return errors.New("list empty")
// 	}

// 	l.sortVenuesByAscendingPrice()

// 	currentNode := l.tail

// 	for currentNode != nil {
// 		fmt.Println("Venue ID:", currentNode.venue.venueID)
// 		fmt.Println("Venue Name:", currentNode.venue.venueName)
// 		fmt.Println("Venue Type:", currentNode.venue.venueType)
// 		fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
// 		fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
// 		fmt.Println()
// 		currentNode = currentNode.prev
// 	}
// 	return nil
// }

// // * sort venues by capacity
// func (l *linkedList) sortVenuesByAscendingCapacity() error {
// 	if l.head == nil {
// 		return errors.New("list empty")
// 	}

// 	currentNode := l.head

// 	for currentNode != nil {
// 		indexNode := currentNode.next

// 		for indexNode != nil {
// 			if currentNode.venue.venueCapacity > indexNode.venue.venueCapacity {
// 				tempVenueID := currentNode.venue.venueID
// 				tempVenueName := currentNode.venue.venueName
// 				tempVenueType := currentNode.venue.venueType
// 				tempVenueCapacity := currentNode.venue.venueCapacity
// 				tempVenuePrice := currentNode.venue.venuePrice
// 				tempVenueSchedule := currentNode.venue.venueSchedule

// 				currentNode.venue.venueID = indexNode.venue.venueID
// 				currentNode.venue.venueName = indexNode.venue.venueName
// 				currentNode.venue.venueType = indexNode.venue.venueType
// 				currentNode.venue.venueCapacity = indexNode.venue.venueCapacity
// 				currentNode.venue.venuePrice = indexNode.venue.venuePrice
// 				currentNode.venue.venueSchedule = indexNode.venue.venueSchedule

// 				indexNode.venue.venueID = tempVenueID
// 				indexNode.venue.venueName = tempVenueName
// 				indexNode.venue.venueType = tempVenueType
// 				indexNode.venue.venueCapacity = tempVenueCapacity
// 				indexNode.venue.venuePrice = tempVenuePrice
// 				indexNode.venue.venueSchedule = tempVenueSchedule
// 			}
// 			indexNode = indexNode.next
// 		}
// 		currentNode = currentNode.next
// 	}
// 	return nil
// }

// func (l *linkedList) sortVenuesByDescendingCapacity() error {
// 	if l.head == nil && l.tail == nil {
// 		return errors.New("list empty")
// 	}

// 	l.sortVenuesByAscendingCapacity()

// 	currentNode := l.tail

// 	for currentNode != nil {
// 		fmt.Println("Venue ID:", currentNode.venue.venueID)
// 		fmt.Println("Venue Name:", currentNode.venue.venueName)
// 		fmt.Println("Venue Type:", currentNode.venue.venueType)
// 		fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
// 		fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
// 		fmt.Println()
// 		currentNode = currentNode.prev
// 	}
// 	return nil
// }

// func (l *linkedList) searchForVenueByVenueID(venueID string) (venueInformation, error) {
// 	if l.head == nil && l.tail == nil {
// 		return venueInformation{}, errors.New("list empty")
// 	}

// 	currentNode := l.head

// 	for currentNode != nil {
// 		if strconv.Itoa(currentNode.venue.venueID) == venueID {
// 			return currentNode.venue, nil
// 		}
// 		currentNode = currentNode.next
// 	}

// 	return venueInformation{}, errors.New("no matches found")
// }

// func (l *linkedList) searchForVenueByName(venueName string) (int, error) {
// 	var venuesFound int

// 	if l.head == nil && l.tail == nil {
// 		return venuesFound, errors.New("list empty")
// 	}

// 	currentNode := l.head

// 	for currentNode != nil {

// 		if strings.Contains(strings.ToLower(currentNode.venue.venueName), strings.ToLower(venueName)) {
// 			fmt.Println("Venue ID:", currentNode.venue.venueID)
// 			fmt.Println("Venue Name:", currentNode.venue.venueName)
// 			fmt.Println("Venue Type:", currentNode.venue.venueType)
// 			fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
// 			fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
// 			fmt.Println()

// 			venuesFound++
// 		}
// 		currentNode = currentNode.next
// 	}

// 	return venuesFound, nil
// }

// func (l *linkedList) searchForVenueByType(venueType string) (int, error) {
// 	var venuesFound int

// 	if l.head == nil && l.tail == nil {
// 		return venuesFound, errors.New("list empty")
// 	}

// 	currentNode := l.head

// 	for currentNode != nil {

// 		if strings.Contains(strings.ToLower(currentNode.venue.venueType), strings.ToLower(venueType)) {
// 			fmt.Println("Venue ID:", currentNode.venue.venueID)
// 			fmt.Println("Venue Name:", currentNode.venue.venueName)
// 			fmt.Println("Venue Type:", currentNode.venue.venueType)
// 			fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
// 			fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
// 			fmt.Println()

// 			venuesFound++
// 		}
// 		currentNode = currentNode.next
// 	}

// 	return venuesFound, nil
// }

// // ----------------------------------- FUNCTIONS --------------------------------- //
// // ! FUNCTIONS

// func seedData() {
// 	day[0] = timeSlot{1, "0000", "0400", bookingStatus{}}
// 	day[1] = timeSlot{2, "0400", "0800", bookingStatus{}}
// 	day[2] = timeSlot{3, "0800", "1200", bookingStatus{}}
// 	day[3] = timeSlot{4, "1200", "1600", bookingStatus{}}
// 	day[4] = timeSlot{5, "1600", "2000", bookingStatus{}}
// 	day[5] = timeSlot{6, "2000", "2400", bookingStatus{}}

// 	week[0] = map[string][6]timeSlot{"21": day}
// 	week[1] = map[string][6]timeSlot{"22": day}
// 	week[2] = map[string][6]timeSlot{"23": day}
// 	week[3] = map[string][6]timeSlot{"24": day}
// 	week[4] = map[string][6]timeSlot{"25": day}
// 	week[5] = map[string][6]timeSlot{"26": day}
// 	week[6] = map[string][6]timeSlot{"27": day}

// 	venueA := venueInformation{1, "Venue A", "Concert Hall", 1000, 2000, week}
// 	venueB := venueInformation{2, "Venue B", "Auditorium", 2000, 4000, week}
// 	venueC := venueInformation{3, "Venue C", "Outdoor Field", 3000, 3000, week}
// 	venueD := venueInformation{4, "Venue D", "Small Meeting Room", 50, 500, week}
// 	venueE := venueInformation{5, "Venue E", "Medium Meeting Room", 100, 100, week}

// 	venues.addNode(venueA)
// 	venues.addNode(venueB)
// 	venues.addNode(venueC)
// 	venues.addNode(venueD)
// 	venues.addNode(venueE)

// 	venueCapacities = []int{1000, 2000, 3000, 50, 100}
// 	venuePrices = []int{2000, 4000, 3000, 500, 100}
// }

// func printTimings(day [6]timeSlot) {
// 	var availableToBook string

// 	for _, timing := range day {
// 		if timing.booking.booked == false {
// 			availableToBook = "Available To Book"
// 		} else {
// 			availableToBook = "Booked"
// 		}

// 		fmt.Printf("%v. %s - %s Booking Status: %s\n", timing.timeSlotID, timing.startingTime, timing.endingTime, availableToBook)
// 	}
// }

// func getAvailableTimings(day [6]timeSlot) ([]int, error) {
// 	indexOfAvailableTimeSlots := make([]int, 0, 6)
// 	var availableToBook string
// 	var availableTimeSlots int

// 	for _, timeSlot := range day {
// 		if timeSlot.booking.booked == false {
// 			availableToBook = "Available To Book"
// 			availableTimeSlots++
// 			indexOfAvailableTimeSlots = append(indexOfAvailableTimeSlots, timeSlot.timeSlotID)
// 			fmt.Printf("%v. %s - %s Booking Status: %s\n", timeSlot.timeSlotID, timeSlot.startingTime, timeSlot.endingTime, availableToBook)
// 		}
// 	}
// 	if availableTimeSlots == 0 {
// 		return indexOfAvailableTimeSlots, errors.New("no timings available for this day")
// 	}

// 	return indexOfAvailableTimeSlots, nil
// }

// func printDays(week [7]map[string][6]timeSlot) {
// 	for _, day := range week {
// 		for date, timing := range day {
// 			fmt.Println(date)
// 			printTimings(timing)
// 			fmt.Println()
// 		}
// 	}
// }

// func dateAvailable(week [7]map[string][6]timeSlot, date string) ([6]timeSlot, error) {
// 	var timeSlots [6]timeSlot
// 	var exists bool

// 	for _, day := range week {
// 		for day, timeSlot := range day {
// 			if day == date {
// 				exists = true
// 				timeSlots = timeSlot
// 			}
// 		}
// 	}

// 	if exists == false {
// 		return timeSlots, errors.New("chosen date not available for booking")
// 	}

// 	return timeSlots, nil
// }

// func chosenTimeAvailable(chosenTimeSlot string, availableTimeSlots []int) error {
// 	for i := 0; i < len(availableTimeSlots); i++ {
// 		if strconv.Itoa(availableTimeSlots[i]) == chosenTimeSlot {
// 			return nil
// 		}
// 	}
// 	return errors.New("chosen time slot not available")
// }

// func getTimeSlot(chosenDate string, chosenTimeSlot string) (timeSlot, error) {
// 	retrievedTimeSlot := timeSlot{}

// 	for _, day := range week {
// 		for day, timeSlots := range day {
// 			if day == chosenDate {
// 				for _, timeSlot := range timeSlots {
// 					if strconv.Itoa(timeSlot.timeSlotID) == chosenTimeSlot {
// 						retrievedTimeSlot = timeSlot
// 						return retrievedTimeSlot, nil
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return retrievedTimeSlot, errors.New("time slot not retrieved")
// }

// func retrieveBookings(userEmail string) {
// 	for _, day := range week {
// 		for date, timeSlots := range day {
// 			for _, timeSlot := range timeSlots {
// 				// err := timeSlot.booking.retrieveBookings(userEmail)

// 				// if err != nil {
// 				// 	return errors.New("no bookings retrieved")
// 				// }

// 				if timeSlot.booking.user.email == userEmail {
// 					var availableToBook string

// 					if timeSlot.booking.booked == true {
// 						availableToBook = "Booked"
// 					} else {
// 						availableToBook = "Available to Book"
// 					}

// 					fmt.Println(date)
// 					fmt.Printf("%v. %s - %s Booking Status: %s\n", timeSlot.timeSlotID, timeSlot.startingTime, timeSlot.endingTime, availableToBook)
// 				}
// 			}
// 		}
// 	}
// }

// ------------------------------------- MAIN ------------------------------------ //

func main() {
	seedData()

	// startProgram()

	for {
		var userInput string

		fmt.Println("Welcome to venue booking system!")
		fmt.Println("\n1. Browse Venues")
		fmt.Println("2. Search Venues")
		fmt.Println("3. Make A Booking")
		fmt.Println("4. View My Bookings")
		fmt.Print("\nPlease Choose: ")

		fmt.Scanln(&userInput)
		fmt.Println()

		switch userInput {
		case "1":
			var arrangeBy string

			fmt.Println("Sort By:")
			fmt.Println("\n1. Price: Low to High")
			fmt.Println("2. Price: High to Low")
			fmt.Println("3. Capacity: Smallest to Largest")
			fmt.Println("4. Capacity: Largest to Smallest")
			fmt.Print("\nPlease Choose: ")

			fmt.Scanln(&arrangeBy)
			fmt.Println()

			switch arrangeBy {
			case "1":
				venues.sortVenuesByAscendingPrice()
				venues.printAllNodes()
			case "2":
				venues.sortVenuesByDescendingPrice()
				// venues.printAllNodes()
			case "3":
				venues.sortVenuesByAscendingCapacity()
				venues.printAllNodes()
			case "4":
				venues.sortVenuesByDescendingCapacity()
				// venues.printAllNodes()
			}

		case "2":
			var searchBy string

			fmt.Println("Search By:")
			fmt.Println("\n1. Venue Name")
			fmt.Println("2. Venue Type")
			fmt.Print("\nPlease Choose: ")

			fmt.Scanln(&searchBy)
			fmt.Println()

			switch searchBy {
			case "1":
				var searchWord string

				fmt.Print("Search For: ")
				fmt.Scanln(&searchWord)
				fmt.Println()

				venuesFound, error := venues.searchForVenueByName(searchWord)

				if error != nil {
					fmt.Println(error)
				} else {
					if venuesFound == 0 {
						fmt.Println("No venues found!")
					}
				}

			case "2":
				var searchWord string

				fmt.Print("Search For: ")
				fmt.Scanln(&searchWord)
				fmt.Println()

				venuesFound, error := venues.searchForVenueByType(searchWord)

				if error != nil {
					fmt.Println(error)
				} else {
					if venuesFound == 0 {
						fmt.Println("No venues found!")
					}
				}
			}

		case "3":
			var (
				venueID        string
				chosenDate     string
				chosenTimeSlot string
			)

			var (
				userFirstName     string
				userLastName      string
				userContactNumber string
				userEmail         string
			)

			venues.sortVenuesByVenueID()
			venues.printAllNodes()

			fmt.Print("Enter venue ID of venue you would like to book: ")

			fmt.Scanln(&venueID)
			fmt.Println()

			venue, err := venues.searchForVenueByVenueID(venueID)

			// * break out of switch case if venueID not recognised
			if err != nil {
				fmt.Println("Venue not found in system!")
				fmt.Println()
				break
			}

			fmt.Printf("Booking schedule for %s:\n\n", venue.venueName)
			printDays(venue.venueSchedule)

			fmt.Print("Choose a Date: ")
			fmt.Scanln(&chosenDate)
			fmt.Println()

			timeSlotsForChosenDate, err := dateAvailable(week, chosenDate)

			// * break out of switch case if date stated not available
			if err != nil {
				fmt.Println("Chosen date not available for booking!")
				fmt.Println()
				break
			}

			// * break out of switch case if all slots are already booked
			availableTimeSlots, err := getAvailableTimings(timeSlotsForChosenDate)

			if err != nil {
				fmt.Println("No available time slots for this date!")
				fmt.Println()
				break
			}

			fmt.Print("\nChoose time slot by keying in timeslot id: ")
			fmt.Scanln(&chosenTimeSlot)

			err = chosenTimeAvailable(chosenTimeSlot, availableTimeSlots)

			// * break out of switch case if user key in out of bounds input for time slot
			if err != nil {
				fmt.Println("Chosen time slot was not available for booking!")
				fmt.Println()
				break
			}

			timeSlotInformation, err := getTimeSlot(chosenDate, chosenTimeSlot)

			if err != nil {
				fmt.Println("Unable to retrieve time slot information!")
				fmt.Println()
				break
			}

			// * ask for user information
			fmt.Println()
			fmt.Println("Please fill in your details to make the booking: ")
			fmt.Print("First Name: ")
			fmt.Scanln(&userFirstName)
			fmt.Print("Last Name: ")
			fmt.Scanln(&userLastName)
			fmt.Print("Contact Number: ")
			fmt.Scanln(&userContactNumber)
			fmt.Print("Email: ")
			fmt.Scanln(&userEmail)

			err = timeSlotInformation.booking.makeBooking(userFirstName, userLastName, userContactNumber, userEmail)

			if err != nil {
				fmt.Println()
				fmt.Println("All fields must be filled in!")
				fmt.Println()
				break
			}

			fmt.Println("\nBooking made Succesfully! Check out 'View My Bookings' from the main page to see all your bookings!")
			fmt.Println()

		case "4":
			var (
				userEmail string
			)

			fmt.Print("Please key in the email you used to make the booking: ")
			fmt.Scanln(&userEmail)

			fmt.Println("Here are your bookings: ")
			fmt.Println()
			retrieveBookings(userEmail)
		}
	}
}
