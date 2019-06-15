package DBController

func (dbc *DBCnt) INSERT() {
	dbc.DB.Create("hoge")
}
