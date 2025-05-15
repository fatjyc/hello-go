package hello

import "fmt"
import "github.com/fatjyc/hello-go/tools"

type Greeter struct {
	Language string
	Name     string
}

func NewGreeter(lang string, name string) *Greeter {
	return &Greeter{Language: lang, Name: name}
}

func (g *Greeter) Greet() string {
	greeting := Languages(g.Language)
	if g.Name == "" {
		return fmt.Sprintf("%s, World!！", greeting)
	}
	return fmt.Sprintf("%s, %s![%s]", greeting, g.Name, tools.Reverse(g.Name))
}
