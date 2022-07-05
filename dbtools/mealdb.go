package dbtools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type MealDB struct {
	JsonPath string

	mealMap map[Protein][]Meal
}

// Init loads the database from memory
func (m *MealDB) Init() error {
	m.mealMap = make(map[Protein][]Meal)
	for i, _ := range acceptedProteins {
		m.mealMap[Protein(i)] = make([]Meal, 0)
	}

	meals := make([]Meal, 0)
	b, err := ioutil.ReadFile(m.JsonPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &meals); err != nil {
		return err
	}

	for _, v := range meals {
		m.mealMap[v.Protein] = append(m.mealMap[v.Protein], v)
	}

	return nil
}

// Save saves the current database to memory - overwriting the existing file
func (m *MealDB) Save() error {
	meals := make([]Meal, 0)
	for i, _ := range acceptedProteins {
		meals = append(meals, m.mealMap[Protein(i)]...)
	}

	b, err := json.MarshalIndent(meals, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(m.JsonPath, b, 0644)
}

func (m *MealDB) String() string {
	ret := "[\n"
	for i, v := range acceptedProteins {
		ret += fmt.Sprintf("\t%v:\n", v)

		for _, m := range m.mealMap[Protein(i)] {
			ret += fmt.Sprintf("\t\t%+v\n", m)
		}
	}
	ret += "]"
	return ret
}

func (m *MealDB) AddMeal(newMeal Meal) {
	// ensure this meal is new (use URL as key)
	for _, m := range m.mealMap[newMeal.Protein] {
		if m.Url == newMeal.Url {
			fmt.Println("This item already exists. Skipping...")
			return
		}
	}

	m.mealMap[newMeal.Protein] = append(m.mealMap[newMeal.Protein], newMeal)
}

func (m *MealDB) GetMealsWithProtein(p Protein) []Meal {
	return m.mealMap[p]
}

func (m *MealDB) Get5RandomMeals() []Meal {
	// slap together list of all meals
	meals := make([]Meal, 0)
	for i, _ := range acceptedProteins {
		meals = append(meals, m.mealMap[Protein(i)]...)
	}

	// choose 5 random ones
	randMeals := make([]Meal, 0)
	rand.Seed(time.Now().UnixNano())
	for len(randMeals) < 5 {
		i := rand.Intn(len(meals))
		newMeal := meals[i]

		// make sure all 5 given are unique
		unique := true
		for _, m := range randMeals {
			if m.Url == newMeal.Url {
				unique = false
			}
		}
		if unique {
			randMeals = append(randMeals, newMeal)
		}
	}
	return randMeals
}

func (m *MealDB) AddDummyData() {
	meal := Meal{
		Protein: chicken,
		Title:   "Chimken Recipe",
		Url:     "www.urmom.com",
		Notes:   "Very Spicy",
		Rating:  2,
	}

	m.mealMap[chicken] = append(m.mealMap[chicken], meal)
	m.mealMap[chicken] = append(m.mealMap[chicken], meal)

	meal2 := Meal{
		Protein: veggie,
		Title:   "veggie Recipe",
		Url:     "www.urmom.com",
		Notes:   "Super Bland",
		Rating:  1,
	}

	m.mealMap[veggie] = append(m.mealMap[veggie], meal2)
}

func (m *MealDB) SaveToCSV(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("unable to create " + path + "...")
		return
	}

	str := fmt.Sprintf("Title,Url,Protein,Rating (1-5),Notes\n")
	_, err = f.WriteString(str)
	if err != nil {
		fmt.Println("unable to write to file")
		return
	}

	for _, p := range acceptedProteins {
		meals := m.mealMap[StringToProtein(p)]

		for _, m := range meals {
			str = fmt.Sprintf("%v,%v,%v,%v,%v\n", m.Title, m.Url, ProteinToString(m.Protein), m.Rating, m.Notes)
			_, err = f.WriteString(str)
			if err != nil {
				fmt.Println("unable to save meal " + m.Title)
				return
			}
		}
	}

}
