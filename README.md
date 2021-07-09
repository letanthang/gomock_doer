add mockgen to go.mod
```
go get github.com/golang/mock/mockgen/model
```
Generate code
```
mockgen -destination=mocks/mock_doer.go -package=mocks app/doer Doer
```
Or use go:generate
```
add to doer/doer.go this commment
....
//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks app/doer Doer
```
and run
```
go generate ./...
```
Run the test
```test
go test 
```