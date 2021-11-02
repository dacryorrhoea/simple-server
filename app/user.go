package main

type Item struct {
	Number int
	Name   string
}

func getAllItems() (items []Item, err error) {
	query := `select * from items`
	err = Db.Select(&items, query)
	return
}
