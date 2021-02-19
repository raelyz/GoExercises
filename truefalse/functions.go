package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ----------------------------------- FUNCTIONS --------------------------------- //
// ! FUNCTIONS

func seedData() {
	day[0] = timeSlot{1, "0000", "0400", bookingStatus{}}
	day[1] = timeSlot{2, "0400", "0800", bookingStatus{}}
	day[2] = timeSlot{3, "0800", "1200", bookingStatus{}}
	day[3] = timeSlot{4, "1200", "1600", bookingStatus{}}
	day[4] = timeSlot{5, "1600", "2000", bookingStatus{}}
	day[5] = timeSlot{6, "2000", "2400", bookingStatus{}}

	week[0] = map[string][6]timeSlot{"21": day}
	week[1] = map[string][6]timeSlot{"22": day}
	week[2] = map[string][6]timeSlot{"23": day}
	week[3] = map[string][6]timeSlot{"24": day}
	week[4] = map[string][6]timeSlot{"25": day}
	week[5] = map[string][6]timeSlot{"26": day}
	week[6] = map[string][6]timeSlot{"27": day}

	venueA := venueInformation{1, "Venue A", "Concert Hall", 1000, 2000, week}
	venueB := venueInformation{2, "Venue B", "Auditorium", 2000, 4000, week}
	venueC := venueInformation{3, "Venue C", "Outdoor Field", 3000, 3000, week}
	venueD := venueInformation{4, "Venue D", "Small Meeting Room", 50, 500, week}
	venueE := venueInformation{5, "Venue E", "Medium Meeting Room", 100, 100, week}

	venues.addNode(venueA)
	venues.addNode(venueB)
	venues.addNode(venueC)
	venues.addNode(venueD)
	venues.addNode(venueE)

	venueCapacities = []int{1000, 2000, 3000, 50, 100}
	venuePrices = []int{2000, 4000, 3000, 500, 100}
}

func printTimings(day [6]timeSlot) {
	var availableToBook string

	for _, timing := range day {
		if timing.booking.booked == false {
			availableToBook = "Available To Book"
		} else {
			availableToBook = "Booked"
		}

		fmt.Printf("%v. %s - %s Booking Status: %s\n", timing.timeSlotID, timing.startingTime, timing.endingTime, availableToBook)
	}
}

func getAvailableTimings(day [6]timeSlot) ([]int, error) {
	indexOfAvailableTimeSlots := make([]int, 0, 6)
	var availableToBook string
	var availableTimeSlots int

	for _, timeSlot := range day {
		if timeSlot.booking.booked == false {
			availableToBook = "Available To Book"
			availableTimeSlots++
			indexOfAvailableTimeSlots = append(indexOfAvailableTimeSlots, timeSlot.timeSlotID)
			fmt.Printf("%v. %s - %s Booking Status: %s\n", timeSlot.timeSlotID, timeSlot.startingTime, timeSlot.endingTime, availableToBook)
		}
	}
	if availableTimeSlots == 0 {
		return indexOfAvailableTimeSlots, errors.New("no timings available for this day")
	}

	return indexOfAvailableTimeSlots, nil
}

func printDays(week [7]map[string][6]timeSlot) {
	for _, day := range week {
		for date, timing := range day {
			fmt.Println(date)
			printTimings(timing)
			fmt.Println()
		}
	}
}

func dateAvailable(week [7]map[string][6]timeSlot, date string) ([6]timeSlot, error) {
	var timeSlots [6]timeSlot
	var exists bool

	for _, day := range week {
		for day, timeSlot := range day {
			if day == date {
				exists = true
				timeSlots = timeSlot
			}
		}
	}

	if exists == false {
		return timeSlots, errors.New("chosen date not available for booking")
	}

	return timeSlots, nil
}

func chosenTimeAvailable(chosenTimeSlot string, availableTimeSlots []int) error {
	for i := 0; i < len(availableTimeSlots); i++ {
		if strconv.Itoa(availableTimeSlots[i]) == chosenTimeSlot {
			return nil
		}
	}
	return errors.New("chosen time slot not available")
}

func getTimeSlot(chosenDate string, chosenTimeSlot string) (timeSlot, error) {
	retrievedTimeSlot := timeSlot{}

	for _, day := range week {
		for day, timeSlots := range day {
			if day == chosenDate {
				for _, timeSlot := range timeSlots {
					if strconv.Itoa(timeSlot.timeSlotID) == chosenTimeSlot {
						retrievedTimeSlot = timeSlot
						return retrievedTimeSlot, nil
					}
				}
			}
		}
	}
	return retrievedTimeSlot, errors.New("time slot not retrieved")
}

func retrieveBookings(userEmail string) {
	for _, day := range week {
		for date, timeSlots := range day {
			for _, timeSlot := range timeSlots {
				// err := timeSlot.booking.retrieveBookings(userEmail)

				// if err != nil {
				// 	return errors.New("no bookings retrieved")
				// }

				if timeSlot.booking.user.email == userEmail {
					var availableToBook string

					if timeSlot.booking.booked == true {
						availableToBook = "Booked"
					} else {
						availableToBook = "Available to Book"
					}

					fmt.Println(date)
					fmt.Printf("%v. %s - %s Booking Status: %s\n", timeSlot.timeSlotID, timeSlot.startingTime, timeSlot.endingTime, availableToBook)
				}
			}
		}
	}
}
