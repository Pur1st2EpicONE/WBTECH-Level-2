package main

import "L2.18/internal/app"

func main() {

	calendarApp := app.Boot()
	defer calendarApp.Stop()

	calendarApp.Run()

}
