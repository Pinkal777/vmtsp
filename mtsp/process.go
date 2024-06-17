package mtsp

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

const (
	DriverCapacity = float64(720.00)
)

type Savings map[string]float64

// checks savings, to determine if continuous delivery rather going to depo
func calculateSavings(p *Problems) Savings {
	//this can be used for optimization later
	//childRouts := make(map[string][]string)
	savings := make(Savings)
	for i := range *p {
		for j := range *p {
			if i != j {
				load1 := (*p)[i]
				load2 := (*p)[j]
				//fmt.Println(DeliveryisSubset(load1, load2))
				//if DeliveryisSubset() returns true then we can add this to childRouts map
				//i.e 3:[5,6,8]-this means delivery route of load-5,6,8 is inbetween delivery route of load-3 route
				//this can be later used to optimize/assign/arrange  loads to perticular driver
				//sij =di0 + d0j − dij
				saving := GetDistanceToDepoet(load1) + GetDistanceFromDepoet(load2) - GetDistanceFromLoad1toLoad2(load1, load2)
				savings[fmt.Sprintf("%d:%d", i, j)] = saving
			}
		}
	}
	return savings
}

func sortReverse(savings Savings) []string {
	keys := make([]string, 0, len(savings))

	for key := range savings {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return savings[keys[i]] > savings[keys[j]]
	})
	return keys
}

func removeDriver(assignedDrivers []Driver, driver *Driver) []Driver {
	result := []Driver{}
	for _, d := range assignedDrivers {
		if !reflect.DeepEqual(d.RouteLoads, driver.RouteLoads) {
			result = append(result, d)
		}
	}
	return result
}

// MTSP
func AssignLoads(p *Problems) []Driver {
	allDrivers := []Driver{}
	loadsSavings := calculateSavings(p)
	//Sort according to maximum savings order
	savings := sortReverse(loadsSavings)
	for _, k := range savings { //start forming route from maximum saving
		key := strings.Split(k, ":")
		l1, _ := strconv.Atoi(key[0])
		l2, _ := strconv.Atoi(key[1])

		load1 := (*p)[l1]
		load2 := (*p)[l2]

		//from here we can start connecting/merging loads to route
		//both loads will assign to same driver if Driver capable as it will return maximim Profit
		if load1.AssignedTo == nil && load2.AssignedTo == nil {

			distance := CalculateRouteDistance(p, []string{load1.LoadNumber, load2.LoadNumber}...)
			if distance <= DriverCapacity { //side constraints to check driver capacity
				newDriver := NewDriver()
				newDriver.RouteLoads = append(newDriver.RouteLoads, load1.LoadNumber, load2.LoadNumber)
				newDriver.DrivenDistance += distance
				allDrivers = append(allDrivers, newDriver)
				(*p)[l1].AssignedTo = &newDriver
				(*p)[l2].AssignedTo = &newDriver
			}
		} else if load1.AssignedTo != nil && load2.AssignedTo == nil { // merge load2 to last if possible

			driver := load1.AssignedTo
			if load1.LoadNumber == driver.RouteLoads[len(driver.RouteLoads)-1] { //if its a last delivery of driver
				distance := CalculateRouteDistance(p, append(driver.RouteLoads, load2.LoadNumber)...)
				if distance <= DriverCapacity { //side constraints to check driver capacity
					allDrivers = removeDriver(allDrivers, driver) //remove old will add updated soon
					driver.RouteLoads = append(driver.RouteLoads, load2.LoadNumber)
					driver.DrivenDistance = distance //reset as we are calculating whole route distance
					(*p)[l2].AssignedTo = driver
					allDrivers = append(allDrivers, *driver)
				}
			}
		} else if load1.AssignedTo == nil && load2.AssignedTo != nil { // merge load1 to first if possible

			driver := load2.AssignedTo
			if load2.LoadNumber == driver.RouteLoads[0] { //if its a First delivery of driver
				distance := CalculateRouteDistance(p, append([]string{load1.LoadNumber}, driver.RouteLoads...)...)
				if distance <= DriverCapacity { //side constraints to check driver capacity
					allDrivers = removeDriver(allDrivers, driver) //remove old will add updated soon
					driver.RouteLoads = append([]string{load1.LoadNumber}, driver.RouteLoads...)
					driver.DrivenDistance = distance //reset as we are calculating whole route distance
					(*p)[l1].AssignedTo = driver
					allDrivers = append(allDrivers, *driver)
				}
			}
		} else { //both loads are already assigned now check if we can merge 2 drivers delivery to one

			driver1 := load1.AssignedTo
			driver2 := load2.AssignedTo
			if driver1 != driver2 && load1.LoadNumber == driver1.RouteLoads[len(driver1.RouteLoads)-1] && load2.LoadNumber == driver2.RouteLoads[0] {
				distance := CalculateRouteDistance(p, append(driver1.RouteLoads, driver2.RouteLoads...)...)
				if distance <= DriverCapacity {
					allDrivers = removeDriver(allDrivers, driver1) //remove old will add updated soon
					driver1.RouteLoads = append(driver1.RouteLoads, driver2.RouteLoads...)
					driver1.DrivenDistance = distance
					for _, loadNumber := range driver2.RouteLoads {
						ln, _ := strconv.Atoi(loadNumber)
						(*p)[ln-1].AssignedTo = driver1
					}
					driver2.DrivenDistance = 0.0
					allDrivers = removeDriver(allDrivers, driver2)
					allDrivers = append(allDrivers, *driver1)
				}
			}
		}
	}

	//assign remaining route to each driver
	for i, load := range *p {
		if load.AssignedTo == nil {
			newDriver := NewDriver()
			newDriver.RouteLoads = append(newDriver.RouteLoads, load.LoadNumber)
			newDriver.DrivenDistance = CalculateRouteDistance(p, newDriver.RouteLoads...)
			allDrivers = append(allDrivers, newDriver)
			(*p)[i].AssignedTo = &newDriver
		}
	}
	return allDrivers
}
