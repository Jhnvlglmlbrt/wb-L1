package main

import "fmt"

// Адаптируемый класс
type Adaptee struct {
	Data string
}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's Specific Request: " + a.Data
}

// Целевой интерфейс
type Target interface {
	Request() string
}

// Адаптер
type Adapter struct {
	Adaptee *Adaptee
}

// Реализация целевого интерфейса через адаптер
func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

func main() {
	adapter := &Adapter{
		Adaptee: &Adaptee{Data: "some data"},
	}

	result := adapter.Request()
	fmt.Println(result)
}
