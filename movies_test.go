package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	//Checks that first element in list is Fargo
	allmovies := Movies
	movie := allmovies[0]
	assert.Contains(t, movie.Title, "Fargo")

}
