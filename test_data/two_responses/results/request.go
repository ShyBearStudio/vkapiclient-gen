package vkapiclient

type Request struct {
	MethodName string
}

func NewRequest(methodName string) *Request {
	return nil
}

func (r *Request) AddIntParam(str string, p int) {

}

func (r *Request) AddStrParam(str string, p string) {

}

func (r *Request) AddBoolParam(str string, p bool) {

}

func (r *Request) AddStrArrParam(str string, p []string) {

}

func (r *Request) Send() []byte {
	var arr []byte
	return arr
}
