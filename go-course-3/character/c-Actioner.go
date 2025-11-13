package character

type C_actioner interface {
	Restore(Health int)
	Gethealth() int
}
