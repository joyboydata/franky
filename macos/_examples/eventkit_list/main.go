package main

import (
	"fmt"

	"github.com/progrium/darwinkit/macos/eventkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

// EventKit list demo: auth status -> list calendars -> predicate -> print 5 events.
// On a Mac without calendar permission this prints the auth status and
// (empty or permission-gated) results — that still proves the bindings load.
func main() {
	store := eventkit.NewEventStore()

	// 1. auth status (deprecated Apple API but still generated; new
	// requestFullAccessToEvents is macOS 14+ and missing from symbolsdb 1.1)
	status := eventkit.EventStore_AuthorizationStatusForEntityType(0) // EKEntityTypeEvent = 0
	fmt.Printf("AuthorizationStatusForEntityType(Event) = %d\n", status)

	// 2. list calendars
	cals := store.CalendarsForEntityType(0)
	fmt.Printf("calendars: %d\n", len(cals))
	for i, c := range cals {
		fmt.Printf("  [%d] %s (type %d)\n", i, c.Title(), c.Type())
		if i >= 9 {
			break
		}
	}

	// 3. predicate for events in next 30 days
	start := foundation.Date_Date()
	end := start.DateByAddingTimeInterval(30 * 24 * 60 * 60)
	pred := store.PredicateForEventsWithStartDateEndDateCalendars(start, end, nil)
	fmt.Printf("predicate: %s\n", pred.Description())

	// 4. matching events, print first 5
	events := store.EventsMatchingPredicate(pred)
	fmt.Printf("events matching: %d\n", len(events))
	for i, e := range events {
		if i >= 5 {
			break
		}
		fmt.Printf("  [%d] %s | %s -> %s | cal: %s\n",
			i, e.Title(),
			e.StartDate().Description(), e.EndDate().Description(),
			e.Calendar().Title())
	}
	fmt.Println("done")
}
