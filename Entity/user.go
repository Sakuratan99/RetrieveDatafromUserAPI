package entity

type User struct {
	Id         int64   `json:"id"`
	Uid        string  `json:"uid"`
	First_name string  `json:"first_name"`
	Last_name  string  `json:"last_name"`
	Username   string  `json:"username"`
	Address    Address `json:"address"`
}

type Address struct {
	City           string      `json:"city"`
	Street_name    string      `json:"street_name"`
	Street_address string      `json:"street_address"`
	Zip_code       string      `json:"zip_code"`
	State          string      `json:"state"`
	Country        string      `json:"country"`
	Coordinates    Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
