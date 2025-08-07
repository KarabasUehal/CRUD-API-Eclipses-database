package storage

import (
	"golang-gin/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Функция инициализации базы данных
func InitDatabase() error {
	var err error
	// Открытие подключения к базе данных SQLite с использованием GORM

	db, err = gorm.Open(sqlite.Open("eclipse.db"), &gorm.Config{})
	if err != nil {
		// Возвращение ошибки, если подключение не удалось
		return err
	}
	// Автоматическое создание таблицы для модели Eclipse, если она еще не существует
	return db.AutoMigrate(models.Eclipse{})

}

// Функция для получения экземпляра базы данных
func GetDB() *gorm.DB {
	// Возвращение глобальной переменной db, содержащей подключение к базе данных
	return db
}
