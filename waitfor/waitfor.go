package waitfor

import "time"

type Profile struct {
	name string
	age  int
}

func saveName(currentProfile *Profile, name string) {
	time.Sleep(3 * time.Second) // Simulate work
	currentProfile.name = name
}

func saveAge(currentProfile *Profile, age int) {
	currentProfile.age = age
}

func updateProfile(currentProfile *Profile, updatedProfile Profile) {
	go saveName(currentProfile, updatedProfile.name)
	saveAge(currentProfile, updatedProfile.age)
}
