package data

type Account struct {
	Id        string
	Email     string
	Role      string
	Campus_id string
	Bde_id    string
}

type AccountList struct {
	Accounts []Account
}
