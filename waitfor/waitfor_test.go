package waitfor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test function
func TestUpdateProfile(t *testing.T) {
	firstName := "Name 1"

	currentProfile := &Profile{
		name: firstName,
		age:  28,
	}
	updatedProfile := Profile{
		name: "Name 2",
		age:  30,
	}

	updateProfile(currentProfile, updatedProfile)

	// saveName takes 3 seconds to complete
	// we have to wait more than 3 seconds
	// if waitFor is less then or equal 3 seconds then it fails

	assert.Eventually(t, func() bool {
		if currentProfile.name != firstName {
			assert.Equal(t, updatedProfile.name, currentProfile.name)
			return true
		}
		return false
	}, 4*time.Second, 100*time.Millisecond)
}
