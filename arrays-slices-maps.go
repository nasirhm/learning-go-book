package main

import "fmt"

func main(){
	// Initializing an Array
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// Average of 5 elements of an array
	var y [5]float64
	y[0] = 98
	y[1] = 99
	y[2] = 85
	y[3] = 56
	y[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)

	// Cleaning the above code by
	total = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	// More cleaning, _ defines we don't need this variable to be used later
	total = 0
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// Shorter syntax for creating arrays.
	a := [5]float64{98,45,43,48,56}
	fmt.Println(a)

	// slices, same as array but with no length
	var j []float64
	// The other clause in make represents the capacity of underlying array
	j = make([]float64, 5,10)
	fmt.Println(j)

	// another way to creat slices is by [low : high] expression:
	arr := [5]float64{1,2,3,4,5}
	e := arr[1:4] // some jargon : arr[0:] equals to arr [0:len(arr)], arr[:5] same as [0:5]
	fmt.Println(e)

	// some fun with slices
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// some more
	slice3 := []int{1,2,3}
	slice4 := make([]int,2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	// Maps
	m := make (map[string]int)
	m["key"] = 10
	fmt.Println(m["key"])

	mi := make(map[int]int)
	mi[1] = 100
	fmt.Println(mi[1])

	// We can delete things from a map as well
	delete(mi,1)
	fmt.Println(mi)

	// Maps with some Elements
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	// Element not found in a map
	fmt.Println((elements["Un"]))
	name, ok := elements["Un"]
	fmt.Println(name, ok) // False
	if name, ok := elements["Un"]; ok{
		fmt.Println(name,ok)
	}

	// Shorter way to make a Map
	elements = map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	// Nesting Maps
	elementss := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := elementss["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// Accessing the 4th element of the Array
	var fourth_element[5] int
	fourth_element[3] = 123
	fmt.Println(fourth_element[3])

	// Length of a slice created using make([]int, 3, 9)
	var sliceee = make([]int,3,9)
	fmt.Println(len(sliceee))

	// array ranges
	r := [6]string{"a","b","c","d","e","f"}
	fmt.Println(r[2:5])

	// Write program to find the smallest number in the list
	sNumber := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97,9,17,
	}
	for l := 0; l<len(sNumber); l++{

	}
}
