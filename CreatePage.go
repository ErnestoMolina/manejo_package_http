// package main

// import (
// 	"os"
// )

// type Page struct {
// 	Title string
// 	Body  []byte
// }

// func (p *Page) save() error {
// 	filename := p.Title + ".txt"
// 	return os.WriteFile(filename, p.Body, 0600)
// }

// func main() {
// 	p1 := &Page{Title: "index", Body: []byte("Hola Mundo")}
// 	p1.save()
// }