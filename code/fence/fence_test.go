package fence

import (
	"testing"
)

func TestIsInCircleFence(t *testing.T) {
	if IsInCircleFence(5000, 114.04788, 22.59758,
		114.01683, 22.68777) {
		t.Log("In circle!")
	} else {
		t.Log("Not in circle")
	}
}
