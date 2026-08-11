package boardgame

import "errors"

// must be capitalized to export the struct, but
// also exported fields must also be capitlized
// so below name and plays are not exported so in another file
// you cannot do the variable.name or variable.plays
// we don't even necessarily have to export the struct and just
// export the constructor function.. something to note

type Boardgame struct {
	name  string
	plays int
}

// !! when the constructor function is in its own package
// it is not uncommong to just call the construct func `New`
// this is bc of how it is called
// e.g., `boardgame.New()`

// !! struct embedding
// a struct has no inheritance or classes
// however, you can still build a struct type that builds on another struct
type Expansion struct {
	standalone bool
	// this is the embedd
	Boardgame
	// you can name the embedd and refer to it in other parts by the name, e.g.:
	// bg Boardgame
}

func NewBoardgame(name string, plays int) (*Boardgame, error) {
	if name == "" {
		return nil, errors.New("game must have a name")
	}

	return &Boardgame{
		name:  name,
		plays: plays,
	}, nil
}

func NewExpansion(standalone bool) Expansion {
	return Expansion{
		standalone: false,
		User: User{
			name: "",
			plays: 0
		}
	}
}
