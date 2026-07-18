type Car struct {
	position int
	speed    int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))

	for i := range position {
		cars[i] = Car{
			position: position[i],
			speed:    speed[i],
		}
	}

	// Process cars from closest to target to farthest.
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	fleets := 0
	lastArrivalTime := 0.0

	for _, car := range cars {
		arrivalTime := float64(target-car.position) / float64(car.speed)

		if arrivalTime > lastArrivalTime {
			// This car cannot catch the fleet ahead,
			// so it creates a new fleet.
			fleets++
			lastArrivalTime = arrivalTime
		}
	}

	return fleets
}