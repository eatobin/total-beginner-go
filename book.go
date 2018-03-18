package totalbeginnergo

type Book struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Borrower Borrower `json:"borrower"`
}

func MakeBook(t string, a string) Book {
	bk := Book{
		Title:    t,
		Author:   a,
		Borrower: Borrower{Name: "NoName", MaxBooks: -1},
	}
	return bk
}

func (bk *Book) SetBorrower(br Borrower) {
	bk.Borrower = br
	return
}

func (bk *Book) availableString() string {
	if bk.Borrower == (Borrower{Name: "NoName", MaxBooks: -1}) {
		return "Available"
	}
	return "Checked out to " +
		bk.Borrower.Name
}

func (bk *Book) BookToString() string {
	return bk.Title +
		" by " + bk.Author +
		"; " + bk.availableString()
}
