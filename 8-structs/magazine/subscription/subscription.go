package subscription

type Subcriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
