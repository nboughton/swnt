// Package dice utilises the standard "nds" notation where n = number of die and s = number of sides; i.e 1d6, 3d10, 8d8 etc
// in order to create individual sets of a single type of die or bags of mixed collections that can then be manipulated or rolled
package dice

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Dice represents a set of 1 type of dice, i.e: 3d20 OR 2d4 OR 1d6
type Dice struct {
	number, sides int
}

// NewDice takes the common notation "nds" where n is the number of dice and s is the number of sides;
// i.e 1d6 and returns a new Dice set. Returns error if s is not a valid dice string
func NewDice(s string) (*Dice, error) {
	number, sides, err := strToVal(s)
	if err != nil {
		return new(Dice), err
	}

	return &Dice{number, sides}, nil
}

// Add adds n die to a single set
func (d *Dice) Add(n int) {
	d.number += n
}

// Remove removes n die from a single set to a minimum of 1
func (d *Dice) Remove(n int) {
	if d.number-n < 1 {
		d.number = 1
	} else {
		d.number -= n
	}
}

// Min returns the minimume possible roll
func (d *Dice) Min() int {
	return d.number
}

// Max returns the maximum possible roll
func (d *Dice) Max() int {
	return d.number * d.sides
}

// Roll all dice in set and return the aggregate result and an array of individual results
func (d *Dice) Roll() (int, []int) {
	t, a := 0, []int{}

	for i := 0; i < d.number; i++ {
		n := rng.Intn(d.sides) + 1
		t += n
		a = append(a, n)
	}

	return t, a
}

// String satisfies the Stringer interface for Dice
func (d *Dice) String() string {
	return fmt.Sprintf("%dd%d", d.number, d.sides)
}

// Bag is a collection of different types of Dice; i.e [3d20, 2d4, 1d6]
type Bag struct {
	dice []*Dice
}

// NewBag returns a new Bag object. A bag can be created with a collection of
// dice specified in string form for convenience. I.e b := NewBag("2d20", "1d6", "8d8").
// Returns error if any item in dice is not a valid dice string
func NewBag(dice ...string) (*Bag, error) {
	b := new(Bag)

	for _, a := range dice {
		if err := b.Add(a); err != nil {
			return b, err
		}
	}

	return b, nil
}

// Add puts more dice in the bag, adding to existing sets where possible.
// Returns error if s is not a valid dice string
func (b *Bag) Add(s string) error {
	d, err := NewDice(s)
	if err != nil {
		return err
	}

	// increment existing set if it exists
	for _, set := range b.dice {
		if set.sides == d.sides {
			set.number += d.number
			return nil
		}
	}

	// Otherwise add a new set
	b.dice = append(b.dice, d)
	return nil
}

// Remove reduces the number of dice by the specified s string if s exists in the bag.
// Returns error if s is not a valid dice string
func (b *Bag) Remove(s string) error {
	number, sides, err := strToVal(s)
	if err != nil {
		return err
	}

	// Remove specified dice from set
	for _, set := range b.dice {
		if set.sides == sides {
			// ensure no < 0 values
			if set.number-number < 0 {
				set.number = 0
			} else {
				set.number -= number
			}
			break
		}
	}

	return nil
}

// Min returns the minimum possible roll
func (b *Bag) Min() int {
	t := 0
	for _, d := range b.dice {
		t += d.Min()
	}
	return t
}

// Max returns the maximum possible roll
func (b *Bag) Max() int {
	t := 0
	for _, d := range b.dice {
		t += d.Max()
	}
	return t
}

// Roll returns aggregate rolls of all Dice in the bag and a map set of results
func (b *Bag) Roll() (int, map[string][]int) {
	t, a := 0, make(map[string][]int)

	for _, d := range b.dice {
		n, s := d.Roll()
		t += n

		a[d.String()] = s
	}

	return t, a
}

// String satisfies the Stringer interface for Bags
func (b *Bag) String() string {
	v := make([]string, len(b.dice))

	for i, d := range b.dice {
		v[i] = fmt.Sprint(d)
	}

	return strings.Join(v, ", ")
}

// returns int values for numbers, sides
func strToVal(a string) (number, sides int, err error) {
	_, err = fmt.Sscanf(a, "%dd%d", &number, &sides)
	if err != nil {
		return number, sides, fmt.Errorf("%s is not a valid dice string", a)
	}

	if number > 0 && sides > 0 {
		return number, sides, nil
	}

	return number, sides, fmt.Errorf("%s is not a valid dice string", a)
}
