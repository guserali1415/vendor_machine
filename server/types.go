package main

type User struct {
	ID    int
	Token string
}

type VendingMachine struct {
	ID                   int
	IDLE                 bool
	AssignedUserID       int
	Name                 string
	Products             []Product
	CoinsInventory       int
	SelectedProductIndex int
}

type Product struct {
	ID    int
	Price int
	Name  string
	Stock int
}

const (
	NOT_FOUND = -1
)
