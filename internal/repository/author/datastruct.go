package author

type Author struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	MiddleName  string `db:"middle_name"`
	BirthYear   int    `db:"birth_year"`
	Description string `db:"description"`
}
