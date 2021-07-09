package user_test

import (
	"app/match"
	"app/mocks"
	"app/user"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUse(t *testing.T)  {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := user.User{mockDoer}

	//mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)
	//mockDoer.EXPECT().DoSomething(123, gomock.Eq("Hello GoMock")).Return(dummyError).Times(1)

	//mockDoer.EXPECT().
	//	DoSomething(123, match.OfType("string")).
	//	Return(dummyError).
	//	Times(1)


	//callFirst := mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1)
	//mockDoer.EXPECT().DoSomething(2, "2nd hello").Return(dummyError).Times(1).After(callFirst)

	//gomock.InOrder(
	//	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(dummyError).Times(1),
	//	mockDoer.EXPECT().DoSomething(2, "2nd hello").Return(dummyError).Times(1),
	//	)

	//mockDoer.EXPECT().
	//	DoSomething(gomock.Any(), gomock.Any()).
	//	Return(dummyError).
	//	Do(func(x int, y string) {
	//		fmt.Println("Called with x=",x," and y =",y)
	//})

	mockDoer.EXPECT().DoSomething(match.OfType("int"), match.OfType("string")).Return(dummyError).
		Do(func(x int, y string) {
			if x > len(y) {
				t.Fail()
			}
	})

	err := testUser.Use()
	if err != dummyError {
		t.Fail()
	}

}
