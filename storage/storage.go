package storage

import (
	"golang-gin/models"
)

var eclipses *[]models.Eclipse

func GetAllEclipses() *[]models.Eclipse {

	db.Find(&eclipses)
	return eclipses
}

func GetEclipseByID(id int) *models.Eclipse {
	var eclipse models.Eclipse
	if result := db.First(&eclipse, id); result.Error != nil {
		return nil // Возвращение nil, если заметка с указанным ID не найдена
	}
	return &eclipse
}

// Функция для создания новой заметки
func AddEclipse(description string, duration float64) models.Eclipse {
	newEclipse := models.Eclipse{
		Description: description,
		Duration:    duration,
	}
	// Использование GORM для выполнения SQL-запроса INSERT и сохранения новой заметки в базе данных
	db.Create(&newEclipse)
	db.Find(&eclipses)
	return newEclipse // Возвращение созданной заметки
}

// Функция для обновления существующей заметки по ID
func UpdateEclipseByID(id int, description string, duration float64) *models.Eclipse {
	var eclipse models.Eclipse
	// Использование GORM для выполнения SQL-запроса SELECT с условием WHERE id = id заметки
	if result := db.First(&eclipse, id); result.Error != nil {
		return nil // Возвращение nil, если заметка с указанным ID не найдена
	}
	eclipse.Description = description
	eclipse.Duration = duration
	// Использование GORM для выполнения SQL-запроса UPDATE и сохранения обновленной заметки в базе данных
	db.Save(eclipse)
	return &eclipse // Возвращение обновленной заметки
}

// Функция для удаления заметки по ID
func DeleteEclipseByID(id int) bool {
	// Использование GORM для выполнения SQL-запроса DELETE с условием WHERE id = id заметки
	if result := db.Delete(&models.Eclipse{}, id); result.Error != nil {
		return false // Возвращение false, если удаление не удалось
	}
	return true // Возвращение true при успешном удалении заметки
}
