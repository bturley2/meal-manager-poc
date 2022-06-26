package main

import (
	"fmt"
	"meal-manager-poc/dbtools"
	"strings"
)

const (
	mainMenuPrompt = `
# ######## MAIN MENU ######### #
1) Find me Something Good
2) Search by protein
3) Add a food
4) Exit
# ############################ #`

	mealSearchPrompt = `
# ########## MEAL SEARCH ########## #
What protein type would you like?
(Options: chicken, beef, turkey, pork, fish, veggie, other)
Type "exit" when finished.
# ################################# #`

	dbPath = "db.json"
)

var (
	mealDB dbtools.MealDB
)

func main() {
	initDb()
	mainMenu()
}

func initDb() {
	mealDB = dbtools.MealDB{
		JsonPath: dbPath,
	}
	if err := mealDB.Init(); err != nil {
		fmt.Printf("Initializing DB may have failed: %v", err)
	}

	fmt.Println(mealDB.String())
}

// mainMenu provides following options:
// 1) Find me Something Good
// 2) Search by protein
// 3) Add a food
// 4) Exit
func mainMenu() {
	var userSelection string

	fmt.Println(mainMenuPrompt)

	for {
		fmt.Printf(">")
		if _, err := fmt.Scanln(&userSelection); err != nil {
			fmt.Println("Invalid selection. Please try again.")
			continue
		}

		switch userSelection {
		case "1":
			getRandomMeal()
		case "2":
			searchMeals()
		case "3":
			addFood()
		case "4":
			fmt.Println("Don't forget to feed Marvin!")
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func getRandomMeal() {

}

func searchMeals() {
	var userSelection string

	for {
		fmt.Println(mealSearchPrompt)
		fmt.Printf(">")
		if _, err := fmt.Scanln(&userSelection); err != nil {
			fmt.Println("Invalid selection. Please try again.")
			continue
		}

		userSelection = strings.TrimSpace(strings.ToLower(userSelection))

		if dbtools.IsValidProtein(userSelection) {
			p := dbtools.StringToProtein(userSelection)
			printMealByProtein(p)
		} else if userSelection == "exit" {
			return
		} else {
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func printMealByProtein(p dbtools.Protein) {
	fmt.Printf("\nHere's all stored meals for %v:\n", dbtools.ProteinToString(p))

	meals := mealDB.GetMealsWithProtein(p)
	for _, m := range meals {
		fmt.Printf("\t%v\n", m.String())
	}
	if len(meals) == 0 {
		fmt.Println("NONE")
	}
}

func addFood() {

}
