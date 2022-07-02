package main

import (
	"bufio"
	"fmt"
	"meal-manager-poc/dbtools"
	"os"
	"strings"
)

const (
	mainMenuPrompt = `
# ######## MAIN MENU ######### #
1) Find me some meals
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
	//fmt.Println(mealDB.String())
}

// mainMenu provides following options:
// 1) Find me Something Good
// 2) Search by protein
// 3) Add a food
// 4) Exit
func mainMenu() {
	var userSelection string

	for {
		fmt.Println(mainMenuPrompt)
		fmt.Printf(">")
		if _, err := fmt.Scanln(&userSelection); err != nil {
			fmt.Println("Invalid selection. Please try again.")
			continue
		}

		switch userSelection {
		case "1":
			get5RandomMeals()
		case "2":
			searchMeals()
		case "3":
			for {
				addMeal()
			}
		case "4":
			fmt.Println("Don't forget to feed Marvin!")
			return
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func get5RandomMeals() {
	rMeals := mealDB.Get5RandomMeals()
	for _, m := range rMeals {
		fmt.Printf("%v\n", m.String())
	}
	if len(rMeals) == 0 {
		fmt.Println("NONE")
	}
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

func addMeal() {
	reader := bufio.NewReader(os.Stdin)

	var err error
	var input string

	m := dbtools.Meal{}

	fmt.Println("\nEnter New Meal Info: ")
	fmt.Print("URL: ")
	if m.Url, err = reader.ReadString('\n'); err != nil || m.Url == "" {
		return
	}
	m.Url = strings.TrimSpace(m.Url)

	fmt.Print("Title: ")
	if m.Title, err = reader.ReadString('\n'); err != nil || m.Title == "" {
		return
	}
	m.Title = strings.TrimSpace(m.Title)

	fmt.Print("Protein: ")
	input, err = reader.ReadString('\n')
	if err != nil || input == "" {
		return
	}
	input = strings.TrimSpace(strings.ToLower(input))
	m.Protein = dbtools.StringToProtein(input)

	//fmt.Print("Rating: ")
	//fmt.Scanln(&m.Url)

	//fmt.Print("Notes: ")
	//fmt.Scanln(&m.Notes)

	mealDB.AddMeal(m)
	mealDB.Save()
	fmt.Println("added and saved!")
}
