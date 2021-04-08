package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	//Checks that first element in list is Fargo
	allmovies := Movies
	movie := allmovies[0]
	assert.Contains("Fargo", movie.Title)

	//Checks that there is 5 titles in the movie list
	var nr int
	for k := range allmovies {
		nr = k + 1
	}
	assert.Equal(5, nr)
}
