package monster

import (
	"fmt"
	"go-course-3/commend"
)

type Slime struct {
	Att   Attributes
	Skill int
}

func (s *Slime) Attract(target commend.Beattracter) {
	fmt.Printf("slime attract,-%d\n", s.GetAttract())
	target.SetHp(target.GetHp() - s.GetAttract())
	fmt.Printf("you now have %d Hp\n", target.GetHp())
}
func (s *Slime) Skillattract(target commend.Beattracter) {
	fmt.Printf("slime Skillattract,-%d\n", s.GetAttract())
	target.SetHp(target.GetHp() - s.GetskillAttract())
	fmt.Printf("you now have %d Hp\n", target.GetHp())
}
func (s *Slime) GetAttract() int {
	return s.Att.Attraction
}
func (s *Slime) GetHp() int {
	return s.Att.Hp
}
func (s *Slime) SetHp(Hp int) {
	s.Att.Hp = Hp
}
func (s *Slime) GetskillAttract() int {
	return s.Skill
}
