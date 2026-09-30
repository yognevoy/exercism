package speed

type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

func NewCar(speed, batteryDrain int) Car {
    return Car{
        battery: 100,
        batteryDrain: batteryDrain,
        speed: speed,
        distance: 0,
    }
}

type Track struct {
    distance int
}

func NewTrack(distance int) Track {
    return Track{distance: distance}
}

func Drive(car Car) Car {
    remains := car.battery - car.batteryDrain
    if remains >= 0 {
        car.battery = max(0, remains)
        car.distance += car.speed
    }
    return car
}

func CanFinish(car Car, track Track) bool {
    driveCount := car.battery / car.batteryDrain
    maxDistance := car.speed * driveCount
    return track.distance <= maxDistance
}
