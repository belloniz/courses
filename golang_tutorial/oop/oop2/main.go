package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName string
	bio string
}

type post struct {
	title string
	content string
	author
}

type website struct {
	posts []post
}

func (a author) fullName() string  {
	return a.firstName + " " + a.lastName
}

func (p post) details()  {
	fmt.Println("Title", p.title)
	fmt.Println("Content", p.content)
	fmt.Println("Author", p.fullName())
}

func (w website) contents()  {
	fmt.Println("Contents of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main()  {
	author1 := author{
		"Gabriel",
		"Belloni",
		"Golag Enthusiast",
	}
	post1 := post{
		"Inheritance in Go",
		"Go support composition instead of inheritance",
		author1,
	}
	post2 := post{
		"Structs instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent lenguage and not a parallel one",
		author1,
	}
	w := website{
		posts: []post{post1, post2, post3},
	}
	w.contents()
}
