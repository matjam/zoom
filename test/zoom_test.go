package test

import (
	"github.com/stephenalexbrowne/zoom"
	. "launchpad.net/gocheck"
	"testing"
)

// We'll define a person struct as the basis of all our tests
// Throughout these, we will try to save, edit, relate, and delete
// Persons in the database
type Person struct {
	Name string
	Age  int
}

// A convenient constructor for our Person struct
func NewPerson(name string, age int) *Person {
	p := &Person{
		Name: name,
		Age:  age,
	}
	return p
}

// Gocheck setup...
func Test(t *testing.T) {
	TestingT(t)
}

type MainSuite struct{}

var _ = Suite(&MainSuite{})

func (s *MainSuite) SetUpSuite(c *C) {
	_, err := zoom.InitDb()
	if err != nil {
		c.Error(err)
	}

	err = zoom.Register(&Person{}, "person")
	if err != nil {
		c.Error(err)
	}
}

func (s *MainSuite) TearDownSuite(c *C) {
	zoom.UnregisterName("person")
	zoom.CloseDb()
}

func (s *MainSuite) TestSave(c *C) {
	p := NewPerson("Bob", 25)
	err := zoom.Save(p)
	if err != nil {
		c.Error(err)
	}
	c.Assert(p.Name, Equals, "Bob")
	c.Assert(p.Age, Equals, 25)
}

// func (s *MainSuite) TestFindById(c *C) {
// 	// Create and save a new model
// 	p1 := NewPerson("Jane", 26)
// 	p1.Save()

// 	// find the model using FindById
// 	result, err := zoom.FindById("person", p1.Id)
// 	if err != nil {
// 		c.Error(err)
// 	}
// 	p2 := result.(*Person)

// 	// Make sure the found model is the same as original
// 	c.Assert(p2.Name, Equals, p1.Name)
// 	c.Assert(p2.Age, Equals, p1.Age)
// }