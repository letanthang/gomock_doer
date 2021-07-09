package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks app/doer Doer

type Doer interface {
	DoSomething(int, string) error
}