package main

import "fmt"

type Achievement struct {
	// ...
}

func NewAchievement() *Achievement {
	a := &Achievement{}
	bus.Subscribe("UserLevelUp", a.OnUserLevelUp)
	return a
}

func (a *Achievement) OnUserLevelUp(arg interface{}) {
	event := arg.(*EventUserLevelUp)
	fmt.Printf("achievement mission old level:%d new level:%d\n", event.oldLevel, event.newLevel)
}
