package services


type TestService struct {

}

var Tests TestService

func (self TestService) Test()(test string){
	test = "test"
	return
}