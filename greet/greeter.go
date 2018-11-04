package greet

import "io"

type Greeter struct {
	Message string
}

func (g *Greeter) Greet(w io.Writer) error {
	_, err := w.Write([]byte(g.Message))
	return err
}
