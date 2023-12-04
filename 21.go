package main

import "fmt"

// Пользователь, которому передадим интерфейс с методом
type user struct {
}

// Интерфейс с методом
type computer interface {
	connectToWiFi()
}

// Структура для которой не нужен адаптер
type compWithModule struct {
}

// Структура, которой нужен адаптер и нет метода из интерфейса
type compWithoutModule struct {
}

// Структура с адаптером и методом из интерфейса
type compWithoutModuleWithAdapter struct {
	comp *compWithoutModule
}

// Метод пользователя, внутри которого метод устройства
func (u *user) turnOnWiFi(comp computer) {
	fmt.Println("Вайфай включен")
	comp.connectToWiFi()
}

// Метод для устройства, которому адаптер не нужен
func (with *compWithModule) connectToWiFi() {
	fmt.Println("Компьютер с модулем подключен к вайфаю")
}

// Метод для устройсва, которому нужен адаптер
func (without *compWithoutModule) connectToWiFiWithountModule() {
	fmt.Println("Компьютер без модуля подключен к вайфаю")
}

// Добавление адаптера
func (withAdapter *compWithoutModuleWithAdapter) connectToWiFi() {
	fmt.Println("Вставляем сетевой адаптер Вайфая")
	withAdapter.comp.connectToWiFiWithountModule()
}

func main() {
	user := &user{}
	// Работа с устройством, у которому не нужен адаптер(С этим все понятно)
	compWith := &compWithModule{}
	user.turnOnWiFi(compWith)
	fmt.Println("------------------------------------------------------------------------------")
	// Тут создаем устройство без адаптрера
	compWithout := &compWithoutModule{}
	// И передаем структуре, являющейся адаптером
	compWithAdapter := &compWithoutModuleWithAdapter{
		comp: compWithout,
	}
	// Теперь устройство без метода через адаптер смогло воспользоваться методом
	user.turnOnWiFi(compWithAdapter)
}
