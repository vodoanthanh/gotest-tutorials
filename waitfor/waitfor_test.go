package waitfor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test function
func TestUpdateProfile(t *testing.T) {
	currentProfile := &Profile{
		name: "Name 1",
		age:  28,
	}
	updatedProfile := Profile{
		name: "Name 2",
		age:  30,
	}

	updateProfile(currentProfile, updatedProfile)

	assert.Eventually(t, func() bool {
		return currentProfile.name == updatedProfile.name
	}, 1*time.Second, 100*time.Millisecond)
}
