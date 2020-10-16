package commands

type Command interface {
	execute() error
}
