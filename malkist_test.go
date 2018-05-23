package malkist_test

import (
	"github.com/pamungkaski/malkist-go"
	"testing"
)

func TestMalkist_ChangeAPIKey(t *testing.T) {
	before := "gOooogleApiKEey"
	m := malkist.Malkist{
		Key: before,
	}
	expected := "bukalapakapikeyhwhhwhwhw"
	m.ChangeAPIKey(expected)

	if m.Key != expected {
		t.Errorf("Change api key from %v to %v failed resulting key to be %v", before, expected, m.Key)
	}
}
