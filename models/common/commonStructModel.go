package common

type Location struct {
	Location_name string `db:"location_name" json:"location_name"`
	Floor         int    `db:"floor" json:"floor"`
}

type Position struct {
	X int	`db:"x" json:"x"`
	Y int	`db:"y" json:"y"`
}