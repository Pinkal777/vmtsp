package mtsp

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x float64
	y float64
}

type Load struct {
	LoadNumber         string
	Pickup             coords
	Dropoff            coords
	DisatnceToDelivery float64
	AssignedTo         *Driver
}

type Problems []Load

func getxy(coordstr string) coords {
	pickupCoord := strings.Split(coordstr, ",")
	x, _ := strconv.ParseFloat(strings.TrimSpace(pickupCoord[0]), 64)
	y, _ := strconv.ParseFloat(strings.TrimSpace(pickupCoord[1]), 64)
	return coords{x: x,
		y: y}
}

func getCoords(record []string) []coords {
	pickupCoordsStr := record[1][1 : len(record[1])-1]
	dropOffCoordsStr := record[2][1 : len(record[2])-1]
	coords := make([]coords, 2)
	coords[0] = getxy(pickupCoordsStr)
	coords[1] = getxy(dropOffCoordsStr)
	return coords
}

// Ls->Le //calculate new from //sqrt((x2-x1)^2 + (y2-y1)^2)
func calculateLoadDistance(load Load) float64 {
	return math.Sqrt((math.Pow((load.Dropoff.x-load.Pickup.x), 2) + math.Pow((load.Dropoff.y-load.Pickup.y), 2)))
}

// load Problem from file
func LoadProblems(path string) (Problems, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading the file [%v]", err)
	}
	// Closes the file
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ' '

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error while reading the problem file[%v]", err)
	}
	loads := make(Problems, len(records)-1)

	for i, record := range records[1:] { //skip header
		coords := getCoords(record)
		loads[i] = Load{LoadNumber: record[0],
			Pickup:  coords[0],
			Dropoff: coords[1]}
		dis := calculateLoadDistance(loads[i])
		loads[i].DisatnceToDelivery = dis
	}
	return loads, nil
}

// func CalDistance123333(loads []Load) []Load {//will remove this
// 	for i, load := range loads {
// 		dis := calculateLoadDistance(load)
// 		loads[i].DisatnceToDelivery = dis
// 	}
// 	return loads
// }

// D->Ls
func GetDistanceFromDepoet(l Load) float64 {
	firstLoad := Load{LoadNumber: "First", Pickup: coords{x: 0, y: 0}, Dropoff: coords{x: l.Pickup.x, y: l.Pickup.y}}
	return calculateLoadDistance(firstLoad)
}

// Le->D
func GetDistanceToDepoet(l Load) float64 {
	lastLoad := Load{LoadNumber: "last", Pickup: coords{x: l.Dropoff.x, y: l.Dropoff.y}, Dropoff: coords{x: 0, y: 0}}
	return calculateLoadDistance(lastLoad)
}

// L1e->l2s
func GetDistanceFromLoad1toLoad2(load1, load2 Load) float64 {
	load1ToLoad2 := Load{LoadNumber: "gap", Pickup: coords{x: load1.Dropoff.x, y: load1.Dropoff.y}, Dropoff: coords{x: load2.Pickup.x, y: load2.Pickup.y}}
	return calculateLoadDistance(load1ToLoad2)
}

// L1s->L1e->L2s->L2e
func distance(load1, load2 Load) float64 {
	load1ToLoad2 := Load{LoadNumber: "gap", Pickup: coords{x: load1.Dropoff.x, y: load1.Dropoff.y}, Dropoff: coords{x: load2.Pickup.x, y: load2.Pickup.y}}
	distanceLoad1ToLoad2 := calculateLoadDistance(load1ToLoad2)
	return load1.DisatnceToDelivery + distanceLoad1ToLoad2 + load2.DisatnceToDelivery
}

// D->L1s->L1e->......->Lns->Lne->d
func CalculateRouteDistance(p *Problems, loadNumbers ...string) float64 {
	totalRouteDistance := 0.0
	ln, _ := strconv.Atoi(loadNumbers[0])
	totalRouteDistance += GetDistanceFromDepoet((*p)[ln-1]) //start from Depo
	for i := range loadNumbers[:len(loadNumbers)-1] {
		lnI, _ := strconv.Atoi(loadNumbers[i])
		lnJ, _ := strconv.Atoi(loadNumbers[i+1])
		totalRouteDistance += distance((*p)[lnI-1], (*p)[lnJ-1])
	}
	ln, _ = strconv.Atoi(loadNumbers[len(loadNumbers)-1])
	totalRouteDistance += GetDistanceToDepoet((*p)[ln-1]) // return to Depo
	return totalRouteDistance
}
