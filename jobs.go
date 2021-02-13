package main

type Job interface {
	GetDescription() string
	GetSalary() int
	GetUserCriterias() []Criteria
}

type Painter struct {
}

func (p *Painter) GetDescription() string {
	return "Painting walls and houses"
}

func (p *Painter) GetUserCriterias() []Criteria {
	return []Criteria{
		&AgeCriteria{Age: 20},
	}
}
func (p *Painter) GetSalary() int {
	return 5
}

type TaxiDriver struct {
}

func (p *TaxiDriver) GetDescription() string {
	return "Driving people from A to B"
}
func (p *TaxiDriver) GetSalary() int {
	return 3
}
