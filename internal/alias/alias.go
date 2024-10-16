package alias

import "fmt"

type Alias struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

func (a *Alias) ToSh() string {
	return fmt.Sprintf("alias %s=%s", a.Name, a.Command)
}
