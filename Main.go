package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"log"
	city "weather/Input"
	w "weather/weatherfunc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	p := tea.NewProgram(city.InitialModel())
	Fmodel, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	if m, ok := Fmodel.(city.Model); ok {
		w.Weather(m.TextInput.Value())
	}
}
