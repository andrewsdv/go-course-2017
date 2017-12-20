package models

type Figure interface {
	Name() string
	Weight() int
}