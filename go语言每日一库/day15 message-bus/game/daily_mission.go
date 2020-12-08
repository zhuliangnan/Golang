package main

import "fmt"

type DailyMission struct {
	// ...
}

func NewDailyMission() *DailyMission {
	d := &DailyMission{}
	bus.Subscribe("UserLevelUp", d.OnUserLevelUp)
	return d
}

func (d *DailyMission) OnUserLevelUp(arg interface{}) {
	event := arg.(*EventUserLevelUp)
	fmt.Printf("daily mission old level:%d new level:%d\n", event.oldLevel, event.newLevel)
}
