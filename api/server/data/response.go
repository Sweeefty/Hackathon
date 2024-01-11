package data

type Response struct {
	Status string
	Code   string
	Data   interface{}
}

type ResponseDataLess struct {
	Status string
	Code   string
}
