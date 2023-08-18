package practice

import (
	"fmt"

	"github.com/jaminwebdev/practice-go/utils"
)

func FunPractice() {
	// funWithMaps()
	// printSeparator()
	// funWithSlices()
	// printSeparator()
	// intermediateMapExample()
	// printSeparator()
	// fmt.Println(addToGroceryList("bread", "zucchini"))
	// printSeparator()
	// getCostsByDay([]cost{
	// 	{day: 0, value: 1.0},
	// 	{day: 1, value: 2.0},
	// 	{day: 1, value: 3.1},
	// 	{day: 2, value: 4},
	// 	{day: 8, value: 3.2},
	// })
	runExternalAppPkg()
}

func funWithMaps() {
	fmt.Println("** FUN WITH MAPS **")
	userEmails := map[int]string{}

	userEmails[1] = "firstuser@something.com"
	userEmails[2] = "seconduser@another.com"

	firstUserEmail := userEmails[1]
	fmt.Println(firstUserEmail)
}

func funWithSlices() {
	fmt.Println("** FUN WITH SLICES **")
	originalNames := []string{"Dave", "Steve", "Mary"}

	newNames := append(originalNames, "Kevin", "Meredith")

	fmt.Println("fetching user names:")
	for _, name := range newNames {
		fmt.Println("Going to get", name, "from the API")
	}
}

func printSeparator() {
	fmt.Println("<------------->")
}

var allMyPets map[string]string = map[string]string{
	"fido":  "dog",
	"stacy": "cat",
	"greg":  "fish",
}

func doesPetExist(petName string) bool {
	_, ok := allMyPets[petName]
	return ok
}

func intermediateMapExample() {
	fmt.Println(doesPetExist("fido"))
}

var initialGroceries = []string{"beans", "bananas", "lemons", "peanut butter"}

func addToGroceryList(newGroceries ...string) []string {
	foods := initialGroceries

	for _, grocery := range newGroceries {
		foods = append(foods, grocery)
	}
	fmt.Println(initialGroceries)
	return foods
}

func runExternalAppPkg() {
	myNum := utils.Add(3, 5, 1, 29, 2)
	utils.PrintNum(myNum)
}

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) {
	costsByDay := []float64{}

	for _, singleCost := range costs {
		// while day >= length of costsByDay slice, continue appending/filling with 0.0
		for singleCost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[singleCost.day] += singleCost.value
	}

	fmt.Println(costsByDay)
}
