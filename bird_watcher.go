package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var ret int = 0
	for i := range birdsPerDay {
		ret += birdsPerDay[i]
	}
	return ret
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var ret int = 0
	for i := range birdsPerDay {
		if i/7 == week-1 {
			ret += birdsPerDay[i]
		}
	}
	return ret
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay)-1; i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
