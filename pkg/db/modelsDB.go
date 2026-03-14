package db

import "time"

// Item представляет запись в таблице items
type Item struct {
	ID       int       `db:"id"`       // уникальный идентификатор записи
	Category string    `db:"category"` // категория ("Продукты", "Транспорт", "Развлечения", "Здоровье", "Другое")
	Amount   float64   `db:"amount"`   // сумма
	Date     time.Time `db:"date"`     // дата в формате YYYY-MM-DD
}
