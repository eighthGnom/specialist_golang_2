package models

func FindBookById(id int) (Book, bool) {
	for _, book := range DB {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}

func FindAndDeleteBookByID(id int) (Book, bool) {
	l := len(DB)
	for index, book := range DB {
		if book.ID == id {
			if index == l-1 {
				DB = DB[:index]
				return book, true
			}
			DB = append(DB[:index], DB[index+1:]...)
			return book, true
		}
	}
	return Book{}, false
}

func FindAndUpdateBookByID(id int, newBook Book) (Book, bool) {
	for index, book := range DB {
		if book.ID == id {
			DB[index] = newBook
			DB[index].ID = id
			return DB[index], true
		}
	}
	return Book{}, false
}
