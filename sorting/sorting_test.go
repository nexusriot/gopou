package sorting

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Album struct {
	Artist string
	Name   string
	Year   int
}

type SorterTestSuite struct {
	suite.Suite
}

func (s *SorterTestSuite) TestSorted() {

	albums := []*Album{
		{"Nightwish", "Oceanborn", 1998},
		{"Nightwish", "Wishmaster", 2000},
		{"Turmion Kätilöt", "Technodiktator", 2013},
		{"Turmion Kätilöt", "Diskovibrator", 2015},
		{"Turmion Kätilöt", "Global Warning", 2020},
		{"Amaranthe", "Manifest", 2020},
	}
	artists := func(a []Sortable) []string {
		var result []string
		for _, v := range a {
			album := interface{}(v).(*Album)
			result = append(result, album.Artist)
		}
		return result
	}

	years := func(a []Sortable) []int {
		var result []int
		for _, v := range a {
			album := interface{}(v).(*Album)
			result = append(result, album.Year)
		}
		return result
	}

	criteria := []*Criteria{
		{"Year", false},
		{"Artist", true},
	}
	sorted, _ := Sorted(criteria, albums)
	expectedYears := []int{2020, 2020, 2015, 2013, 2000, 1998}
	expectedArtists := []string{"Amaranthe", "Turmion Kätilöt", "Turmion Kätilöt", "Turmion Kätilöt", "Nightwish", "Nightwish"}
	albumYears := years(sorted)
	albumArtists := artists(sorted)
	s.Require().Equal(expectedYears, albumYears)
	s.Require().Equal(expectedArtists, albumArtists)

	criteria = []*Criteria{
		{"Artist", true},
		{"Year", true},
	}
	sorted, _ = Sorted(criteria, albums)
	expectedArtists = []string{"Amaranthe", "Nightwish", "Nightwish", "Turmion Kätilöt", "Turmion Kätilöt", "Turmion Kätilöt"}
	expectedYears = []int{2020, 1998, 2000, 2013, 2015, 2020}
	albumYears = years(sorted)
	albumArtists = artists(sorted)
	s.Require().Equal(expectedYears, albumYears)
	s.Require().Equal(expectedArtists, albumArtists)
}

func TestSortingTestSuite(t *testing.T) {
	suite.Run(t, new(SorterTestSuite))
}
