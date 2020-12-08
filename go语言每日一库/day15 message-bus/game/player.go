package main

type Player struct {
	level uint32 //初始值0
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) LevelUp() {
	event := &EventUserLevelUp{
		oldLevel: p.level,
		newLevel: p.level + 1,
	}
	/*	oldLevel := p.level
		newLevel := p.level+1*/
	p.level++

	//等级升级 向topic发布信息 oldLevel  newLevel
	bus.Publish("UserLevelUp", event)
}

type EventUserLevelUp struct {
	oldLevel uint32
	newLevel uint32
}
