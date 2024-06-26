package main

import (
	"fmt"
	"time"
)

func main() {
	// Define target time
	targetTime := time.Date(2024, 6, 12, 15, 0, 0, 0, time.UTC)

	// Get current time
	currentTime := time.Now()

	// Calculate remaining duration
	duration := targetTime.Sub(currentTime)

	// Print initial message
	fmt.Println("Countdown to 3:00 PM on Wednesday, June 12th, 2024:")

	// Loop until target time is reached
	for duration > 0 {
		// **Fix for days calculation**
		days := int(duration.Hours() / 24) // Convert hours to whole days (cast to int)
		hours := int(duration.Hours()) % 24
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60

		// Print remaining time
		fmt.Printf("  - %d days, %d hours, %d minutes, %d seconds remaining\n", days, hours, minutes, seconds)

		// Sleep for 1 second
		time.Sleep(1 * time.Second)

		// Recalculate remaining duration
		duration = targetTime.Sub(time.Now())
	}

	// Print message when target time is reached
	fmt.Println("Target time reached!")
}
