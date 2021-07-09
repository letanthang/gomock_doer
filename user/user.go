package user

import "app/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	//u.Doer.DoSomething(123, "Hello GoMock")

	return u.Doer.DoSomething(2, "2nd hello")
}