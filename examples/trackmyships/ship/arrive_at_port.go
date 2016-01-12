package ship

func (ship *Ship) Arrive(port string) (string, error) {
	if len(port) == 0 {
		return "Arrival port cannot be blank.", nil
	}
	if ship.location == port {
		return "Cannot record arrival: Ship is already at this port.", nil
	}
	if ship.location != AtSea {
		return "Cannot arrive at a port: Ship is not at sea.", nil
	}
	ship.trackChange(Arrived{Port: port})
	return "Success", nil
}
