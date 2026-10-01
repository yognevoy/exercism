package jedlik

import "fmt"

func (c *Car) Drive() {
    remains := c.battery - c.batteryDrain
    if remains >= 0 {
        c.battery = max(0, remains)
        c.distance += c.speed
    }
}

func (c Car) DisplayDistance() string {
    remains := 100 - c.battery
    driven := c.speed * remains / c.batteryDrain
    return fmt.Sprintf("Driven %d meters", driven)
}

func (c Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
    driveCount := c.battery / c.batteryDrain
    maxDistance := c.speed * driveCount
    return trackDistance <= maxDistance
}
