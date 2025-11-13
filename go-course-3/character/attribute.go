package character

import (
	"fmt"
	"go-course-3/commend"
)

type Attribute struct {
	Hp         int
	Attraction int
	Health     int
	Skill      int
}

func (a *Attribute) Attract(target commend.Beattracter) {
	fmt.Printf("you attract,-%d\n", a.GetAttract())
	target.SetHp(target.GetHp() - a.GetAttract())
	fmt.Printf("slime now have %d Hp\n", target.GetHp())
	fmt.Printf("you now have %d Hp\n", a.GetHp())
}
func (a *Attribute) Restore(Health int) {
	a.Hp += Health
	fmt.Printf("yuo restore,+%d\n", Health)
	fmt.Printf("you now have %d Hp\n", a.GetHp())
}
func (a *Attribute) GetAttract() int {
	return a.Attraction
}
func (a *Attribute) GetHp() int {
	return a.Hp
}
func (a *Attribute) SetHp(Hp int) {
	a.Hp = Hp
}
func (a *Attribute) Skillattract(target commend.Beattracter) {
	fmt.Printf("you attract,-%d\n", a.GetAttract())
	target.SetHp(target.GetHp() - a.GetAttract())
	fmt.Printf("slime now have %d Hp\n", target.GetHp())
	fmt.Printf("you now have %d Hp\n", a.GetHp())
}
func (a *Attribute) GetskillAttract() int {
	return a.Skill
}
func (a *Attribute) Gethealth() int {
	return a.Health
}
