package clue

import "context"

type Meta interface {
	GetMessage() string
	SetMessage(string)
	SetCode(string)
	GetCode() string
	GetInfo() interface{}
	Templating(ctx context.Context, clue *Clue) *Clue
	MarshalJSON(ctx context.Context, clue *Clue) ([]byte, error)
}
