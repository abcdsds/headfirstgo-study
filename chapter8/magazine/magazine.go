package magazine

type Employee struct {
	Name string
	Salary float64
	address Address
}


type employee struct {
	Name string
	Salary float64
	address Address
	//익명 필드를 정의 할수 있다.
	Address
}

type Address struct {
	Street string
	City string
	State string
	PostalCode string
}
