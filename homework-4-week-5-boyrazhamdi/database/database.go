package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"homework-3-week-5-boyrazhamdi/author"
	"homework-3-week-5-boyrazhamdi/book"
)

func Connect() (*gorm.DB, error) {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("cannot open database %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

type Repositories struct {
	Book   book.BookRepository
	Author author.AuthorRepository
	DB     *gorm.DB
}

func (repo Repositories) Migration() error {
	return repo.DB.AutoMigrate(book.Book{}, author.Author{})
}

func NewRepositories() (*Repositories, error) {
	db, err := Connect()
	if err != nil {
		return &Repositories{}, err
	}
	return &Repositories{
		Book:   book.NewBookRepository(db),
		Author: author.NewAuthorRepository(db),
		DB:     db,
	}, nil
}
