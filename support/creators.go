package support

import (
	"github.com/stephenalexbrowne/zoom"
	"github.com/stephenalexbrowne/zoom/util"
	"strconv"
)

func CreatePersons(num int) ([]*Person, error) {
	results := make([]*Person, num)
	for i := 0; i < num; i++ {
		p := &Person{
			Name: "person_" + strconv.Itoa(i),
			Age:  i,
		}
		if err := zoom.Save(p); err != nil {
			return results, err
		}
		results[i] = p
	}
	return results, nil
}

func CreateArtists(num int) ([]*Artist, error) {
	results := make([]*Artist, num)
	for i := 0; i < num; i++ {
		a := &Artist{
			Name: "artist_" + strconv.Itoa(i),
		}
		if err := zoom.Save(a); err != nil {
			return results, err
		}
		results[i] = a
	}
	return results, nil
}

func CreateColors(num int) ([]*Color, error) {
	results := make([]*Color, num)
	for i := 0; i < num; i++ {
		val := i % 255
		c := &Color{
			R: val,
			G: val,
			B: val,
		}
		if err := zoom.Save(c); err != nil {
			return results, err
		}
		results[i] = c
	}
	return results, nil
}

func CreatePetOwners(num int) ([]*PetOwner, error) {
	results := make([]*PetOwner, num)
	for i := 0; i < num; i++ {
		p := &PetOwner{
			Name: "petOwner_" + strconv.Itoa(i),
		}
		if err := zoom.Save(p); err != nil {
			return results, err
		}
		results[i] = p
	}
	return results, nil
}

func CreatePets(num int) ([]*Pet, error) {
	results := make([]*Pet, num)
	for i := 0; i < num; i++ {
		p := &Pet{
			Name: "pet_" + strconv.Itoa(i),
		}
		if err := zoom.Save(p); err != nil {
			return results, err
		}
		results[i] = p
	}
	return results, nil
}

func CreateFriends(num int) ([]*Friend, error) {
	results := make([]*Friend, num)
	for i := 0; i < num; i++ {
		f := &Friend{
			Name: "friend_" + strconv.Itoa(i),
		}
		if err := zoom.Save(f); err != nil {
			return results, err
		}
		results[i] = f
	}
	return results, nil
}

func CreateConnectedFriends(num int) ([]*Friend, error) {
	friends, err := CreateFriends(num)
	if err != nil {
		return friends, err
	}

	// randomly connect the friends to one another
	for i1, f1 := range friends {
		numFriends := util.RandInt(0, len(friends)-1)
		for n := 0; n < numFriends; n++ {
			i2 := util.RandInt(0, len(friends)-1)
			f2 := friends[i2]
			for i2 == i1 || friendListContains(f2, f1.Friends) {
				i2 = util.RandInt(0, len(friends)-1)
				f2 = friends[i2]
			}
			f1.Friends = append(f1.Friends, f2)
		}
		// resave the model
		if err := zoom.Save(f1); err != nil {
			return friends, err
		}
	}

	return friends, nil
}

func friendListContains(f *Friend, list []*Friend) bool {
	for _, e := range list {
		if e == f {
			return true
		}
	}
	return false
}
