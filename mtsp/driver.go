package mtsp

type Driver struct {
	DrivenDistance float64
	RouteLoads     []string
	
}

func NewDriver() Driver {
	return Driver{DrivenDistance: 0.0, RouteLoads: []string{}}
}
