package dbtools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (m *MealDB) GetMealsWithProtein(p Protein) []Meal {
	return nil
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
