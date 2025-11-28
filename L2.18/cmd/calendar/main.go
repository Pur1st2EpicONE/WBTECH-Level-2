package main

import "L2.18/internal/app"

func main() {

	calendarApp := app.Start()
	defer calendarApp.Stop()

	calendarApp.Run()

}
