package compra

import "time"

type Compras struct {
	Mercado string
	Data    time.Time
	Item    []Items
}

type Items struct {
	Name  string
	Preco float64
}
