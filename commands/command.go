package commands

type Command interface {
	Execute() error
	GetPayload() error
}

type CommandPayload interface {}
