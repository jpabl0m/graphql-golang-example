package graphql

type User struct {
	ID   string
	Name string
}

type Post struct {
	ID      string
	Title   string
	Content string
}

type Comment struct {
	ID      string
	Content string
}
