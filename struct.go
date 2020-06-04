package krpos

import "fmt"

// POS represents Korean Part Of Speech; 품사
type POS struct {
	Pos         string
	Category    string
	Description string
}

func (p POS) String() string {
	return fmt.Sprintf("%s(%s-%s)", p.Pos, p.Category, p.Description)
}
