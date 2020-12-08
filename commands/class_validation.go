package commands

//Class enum type
type Class string

//Defines the possible classes a user can choose
const (
	Rogue  Class = "rogue"
	Knight Class = "knight"
	Wizard Class = "wizard"
)

//IsValid returns true if the class is an existing one
func (class Class) IsValid() bool {
	switch class {
	case Rogue, Knight, Wizard:
		return true
	}
	return false
}
