package gocamp

type PlayerInteracter interface {
	ProcessMessage(message string) (response string)
}
