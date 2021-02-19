package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ----------------------------------- STRUCTS ----------------------------------- //

type userInformation struct {
	firstName     string
	lastName      string
	contactNumber string
	email         string
}

type bookingStatus struct {
	booked bool
	user   userInformation
}

type timeSlot struct {
	timeSlotID   int
	startingTime string
	endingTime   string
	booking      bookingStatus
}

type venueInformation struct {
	venueID       int
	venueName     string
	venueType     string
	venueCapacity int
	venuePrice    float64
	venueSchedule [7]map[string][6]timeSlot
}

type node struct {
	venue venueInformation
	next  *node
	prev  *node
}

type linkedList struct {
	head *node
	tail *node
	size int
}

// ----------------------------------- METHODS ----------------------------------- //
// ! METHODS

// * userInformation methods

func (u *userInformation) fillInUserInformation(firstName string, lastName string, contactNumber string, email string) error {
	if firstName == "" || lastName == "" || contactNumber == "" || email == "" {
		return errors.New("all fields must be filled in")
	}

	u.firstName = firstName
	u.lastName = lastName
	u.contactNumber = contactNumber
	u.email = email

	return nil
}

func (u *userInformation) updateUserInformation(updatedFirstName string, updatedLastName string, updatedContactNumber string, updatedEmail string) error {
	var edits int

	if updatedFirstName != "" {
		u.firstName = updatedFirstName
		edits++
	}

	if updatedLastName != "" {
		u.lastName = updatedLastName
		edits++
	}

	if updatedContactNumber != "" {
		u.contactNumber = updatedContactNumber
		edits++
	}

	if updatedEmail != "" {
		u.email = updatedEmail
		edits++
	}

	if edits == 0 {
		return errors.New("no edits made")
	}

	return nil
}

func (u *userInformation) deleteUserInformation() {
	u.firstName = ""
	u.lastName = ""
	u.contactNumber = ""
	u.email = ""
}

// func (u *userInformation) matchUserDetails(userEmail string) error {
// 	if u.email == userEmail {
// 		return nil
// 	}
// 	return errors.New("no bookings found under this email address")
// }

// * bookingStatus methods

func (b *bookingStatus) makeBooking(firstName string, lastName string, contactNumber string, email string) error {
	error := b.user.fillInUserInformation(firstName, lastName, contactNumber, email)

	if error != nil {
		return error
	}

	b.booked = true
	fmt.Println(b.booked)
	return nil
}

func (b *bookingStatus) removeBooking() {
	b.booked = false
	b.user.deleteUserInformation()
}

// func (b *bookingStatus) retrieveBookings(userEmail string) error {
// 	err := b.user.matchUserDetails(userEmail)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// * linkedList methods

func (l *linkedList) addNode(venue venueInformation) {
	newNode := &node{venue, nil, nil}

	if l.head == nil && l.tail == nil {
		l.head = newNode
	} else {
		currentNode := l.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
		newNode.prev = currentNode
	}
	l.tail = newNode
	l.size++
}

func (l *linkedList) printAllNodes() error {
	if l.head == nil && l.tail == nil {
		return errors.New("list empty")
	}

	currentNode := l.head

	for currentNode != nil {
		fmt.Println("Venue ID:", currentNode.venue.venueID)
		fmt.Println("Venue Name:", currentNode.venue.venueName)
		fmt.Println("Venue Type:", currentNode.venue.venueType)
		fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
		fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
		fmt.Println()

		currentNode = currentNode.next
	}

	return nil
}

// * sort by venue id number
func (l *linkedList) sortVenuesByVenueID() error {
	if l.head == nil && l.tail == nil {
		return errors.New("list empty")
	}

	currentNode := l.head

	for currentNode != nil {
		indexNode := currentNode.next

		for indexNode != nil {
			if currentNode.venue.venueID > indexNode.venue.venueID {
				tempVenueID := currentNode.venue.venueID
				tempVenueName := currentNode.venue.venueName
				tempVenueType := currentNode.venue.venueType
				tempVenueCapacity := currentNode.venue.venueCapacity
				tempVenuePrice := currentNode.venue.venuePrice
				tempVenueSchedule := currentNode.venue.venueSchedule

				currentNode.venue.venueID = indexNode.venue.venueID
				currentNode.venue.venueName = indexNode.venue.venueName
				currentNode.venue.venueType = indexNode.venue.venueType
				currentNode.venue.venueCapacity = indexNode.venue.venueCapacity
				currentNode.venue.venuePrice = indexNode.venue.venuePrice
				currentNode.venue.venueSchedule = indexNode.venue.venueSchedule

				indexNode.venue.venueID = tempVenueID
				indexNode.venue.venueName = tempVenueName
				indexNode.venue.venueType = tempVenueType
				indexNode.venue.venueCapacity = tempVenueCapacity
				indexNode.venue.venuePrice = tempVenuePrice
				indexNode.venue.venueSchedule = tempVenueSchedule
			}
			indexNode = indexNode.next
		}
		currentNode = currentNode.next
	}
	return nil
}

// * sort by venue price
func (l *linkedList) sortVenuesByAscendingPrice() error {
	if l.head == nil {
		return errors.New("list empty")
	}

	currentNode := l.head

	for currentNode != nil {
		indexNode := currentNode.next

		for indexNode != nil {
			if currentNode.venue.venuePrice > indexNode.venue.venuePrice {
				tempVenueID := currentNode.venue.venueID
				tempVenueName := currentNode.venue.venueName
				tempVenueType := currentNode.venue.venueType
				tempVenueCapacity := currentNode.venue.venueCapacity
				tempVenuePrice := currentNode.venue.venuePrice
				tempVenueSchedule := currentNode.venue.venueSchedule

				currentNode.venue.venueID = indexNode.venue.venueID
				currentNode.venue.venueName = indexNode.venue.venueName
				currentNode.venue.venueType = indexNode.venue.venueType
				currentNode.venue.venueCapacity = indexNode.venue.venueCapacity
				currentNode.venue.venuePrice = indexNode.venue.venuePrice
				currentNode.venue.venueSchedule = indexNode.venue.venueSchedule

				indexNode.venue.venueID = tempVenueID
				indexNode.venue.venueName = tempVenueName
				indexNode.venue.venueType = tempVenueType
				indexNode.venue.venueCapacity = tempVenueCapacity
				indexNode.venue.venuePrice = tempVenuePrice
				indexNode.venue.venueSchedule = tempVenueSchedule
			}
			indexNode = indexNode.next
		}
		currentNode = currentNode.next
	}
	return nil
}

func (l *linkedList) sortVenuesByDescendingPrice() error {
	if l.head == nil && l.tail == nil {
		return errors.New("list empty")
	}

	l.sortVenuesByAscendingPrice()

	currentNode := l.tail

	for currentNode != nil {
		fmt.Println("Venue ID:", currentNode.venue.venueID)
		fmt.Println("Venue Name:", currentNode.venue.venueName)
		fmt.Println("Venue Type:", currentNode.venue.venueType)
		fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
		fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
		fmt.Println()
		currentNode = currentNode.prev
	}
	return nil
}

// * sort venues by capacity
func (l *linkedList) sortVenuesByAscendingCapacity() error {
	if l.head == nil {
		return errors.New("list empty")
	}

	currentNode := l.head

	for currentNode != nil {
		indexNode := currentNode.next

		for indexNode != nil {
			if currentNode.venue.venueCapacity > indexNode.venue.venueCapacity {
				tempVenueID := currentNode.venue.venueID
				tempVenueName := currentNode.venue.venueName
				tempVenueType := currentNode.venue.venueType
				tempVenueCapacity := currentNode.venue.venueCapacity
				tempVenuePrice := currentNode.venue.venuePrice
				tempVenueSchedule := currentNode.venue.venueSchedule

				currentNode.venue.venueID = indexNode.venue.venueID
				currentNode.venue.venueName = indexNode.venue.venueName
				currentNode.venue.venueType = indexNode.venue.venueType
				currentNode.venue.venueCapacity = indexNode.venue.venueCapacity
				currentNode.venue.venuePrice = indexNode.venue.venuePrice
				currentNode.venue.venueSchedule = indexNode.venue.venueSchedule

				indexNode.venue.venueID = tempVenueID
				indexNode.venue.venueName = tempVenueName
				indexNode.venue.venueType = tempVenueType
				indexNode.venue.venueCapacity = tempVenueCapacity
				indexNode.venue.venuePrice = tempVenuePrice
				indexNode.venue.venueSchedule = tempVenueSchedule
			}
			indexNode = indexNode.next
		}
		currentNode = currentNode.next
	}
	return nil
}

func (l *linkedList) sortVenuesByDescendingCapacity() error {
	if l.head == nil && l.tail == nil {
		return errors.New("list empty")
	}

	l.sortVenuesByAscendingCapacity()

	currentNode := l.tail

	for currentNode != nil {
		fmt.Println("Venue ID:", currentNode.venue.venueID)
		fmt.Println("Venue Name:", currentNode.venue.venueName)
		fmt.Println("Venue Type:", currentNode.venue.venueType)
		fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
		fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
		fmt.Println()
		currentNode = currentNode.prev
	}
	return nil
}

func (l *linkedList) searchForVenueByVenueID(venueID string) (venueInformation, error) {
	if l.head == nil && l.tail == nil {
		return venueInformation{}, errors.New("list empty")
	}

	currentNode := l.head

	for currentNode != nil {
		if strconv.Itoa(currentNode.venue.venueID) == venueID {
			return currentNode.venue, nil
		}
		currentNode = currentNode.next
	}

	return venueInformation{}, errors.New("no matches found")
}

func (l *linkedList) searchForVenueByName(venueName string) (int, error) {
	var venuesFound int

	if l.head == nil && l.tail == nil {
		return venuesFound, errors.New("list empty")
	}

	currentNode := l.head

	for currentNode != nil {

		if strings.Contains(strings.ToLower(currentNode.venue.venueName), strings.ToLower(venueName)) {
			fmt.Println("Venue ID:", currentNode.venue.venueID)
			fmt.Println("Venue Name:", currentNode.venue.venueName)
			fmt.Println("Venue Type:", currentNode.venue.venueType)
			fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
			fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
			fmt.Println()

			venuesFound++
		}
		currentNode = currentNode.next
	}

	return venuesFound, nil
}

func (l *linkedList) searchForVenueByType(venueType string) (int, error) {
	var venuesFound int

	if l.head == nil && l.tail == nil {
		return venuesFound, errors.New("list empty")
	}

	currentNode := l.head

	for currentNode != nil {

		if strings.Contains(strings.ToLower(currentNode.venue.venueType), strings.ToLower(venueType)) {
			fmt.Println("Venue ID:", currentNode.venue.venueID)
			fmt.Println("Venue Name:", currentNode.venue.venueName)
			fmt.Println("Venue Type:", currentNode.venue.venueType)
			fmt.Println("Venue Capacity:", currentNode.venue.venueCapacity, "people")
			fmt.Println("Venue Price:", currentNode.venue.venuePrice, "(per 4 hour block)")
			fmt.Println()

			venuesFound++
		}
		currentNode = currentNode.next
	}

	return venuesFound, nil
}
