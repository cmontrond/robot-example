package main

import (
	"gobot.io/x/gobot"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspiAdaptor)

	work := func() {
		//on := uint8(0xFF)
		gobot.Every(1000*time.Millisecond, func() {

			_ = gopigo3.SetMotorPosition(g.MOTOR_RIGHT, 180)

			// Move the wheels
			_ = gopigo3.SetMotorDps(g.MOTOR_LEFT, -150)
			_ = gopigo3.SetMotorDps(g.MOTOR_RIGHT, -150)

			// Blink LEDs
			//err2:= gopigo3.SetLED(g.LED_EYE_RIGHT, 0x00, 0x00, on)
			//if err2 != nil {
			//	fmt.Println(err)
			//}
			//
			//err3 := gopigo3.SetLED(g.LED_EYE_LEFT, ^on, 0x00, 0x00)
			//if err3 != nil {
			//	fmt.Println(err)
			//}
			//on = ^on
		})

	}

	robot := gobot.NewRobot("gopigo3",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gopigo3},
		work,
	)

	robot.Start()
}
