package main
// ======================= Job Interface =======================

type Job interface {
	GetDescription() string
	GetSalary() int
}

// ======================= PainterJOB =======================

type Painter struct {
}

func (p *Painter) GetDescription() string {
	return "Painting walls and houses"
}

func (p *Painter) GetSalary() int {
	return 5
}

// ======================= TaxiDriverJOB =======================

type TaxiDriver struct {
}

func (p *TaxiDriver) GetDescription() string {
	return "Driving people from A to B"
}
func (p *TaxiDriver) GetSalary() int {
	return 3
}
