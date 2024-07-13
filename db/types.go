package db

type Product struct {
	Name   string
	Slug   string
	Detail string
	CategoryId int
}

type Category struct {
	Id int
	Name string
	Slug string
}