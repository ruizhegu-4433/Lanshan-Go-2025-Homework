package commend

type Attracter interface {
	Attract(target Beattracter)
	Skillattract(target Beattracter)
	GetskillAttract() int
	GetAttract() int
}
type Beattracter interface {
	SetHp(Hp int)
	GetHp() int
}
type A_B interface {
	Attracter
	Beattracter
}
