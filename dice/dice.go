package dice

import "github.com/nboughton/go-roll"

// D5 used in a few roll tables
var D5 = roll.NewDie(roll.Faces{{N: 1, Value: "1"}, {N: 2, Value: "2"}, {N: 3, Value: "3"}, {N: 4, Value: "4"}, {N: 5, Value: "5"}})
