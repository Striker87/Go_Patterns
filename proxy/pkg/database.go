package pkg

type DataBase struct {
}

func (db DataBase) GetData(user string) ([]string, error) {
	return []string{
		"first row",
		"last row",
	}, nil
}
