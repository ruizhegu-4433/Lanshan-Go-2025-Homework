package monster

import (
	"fmt"
	"go-course-3/commend"
)

type Dagon struct {
	Att   Attributes
	Skill int
}

func (d *Dagon) Attract(target commend.Beattracter) {
	fmt.Printf("Dangon attract,-%d\n", d.GetAttract())
	target.SetHp(target.GetHp() - d.GetAttract())
	fmt.Printf("you now have %d Hp\n", target.GetHp())
}
func (d *Dagon) Skillattract(target commend.Beattracter) {
	fmt.Printf("Dangon Skillattract,-%d\n", d.GetAttract())
	target.SetHp(target.GetHp() - d.GetskillAttract())
	fmt.Printf("you now have %d Hp\n", target.GetHp())
}
func (d *Dagon) GetAttract() int {
	return d.Att.Attraction
}
func (d *Dagon) GetHp() int {
	return d.Att.Hp
}
func (d *Dagon) SetHp(Hp int) {
	d.Att.Hp = Hp
}
func (d *Dagon) GetskillAttract() int {
	return d.Skill
}
