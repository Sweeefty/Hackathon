package data

type Request struct {
	Id         string
	Account_id string
	Campus_id  string
	Comment    string
	Title      string
	Anonymous  string
}

type RequestList struct {
	RequestList []Request
}
