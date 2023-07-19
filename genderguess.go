package ngender_go

import (
	_ "embed"
	"encoding/csv"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Guesser struct {
	MaleTotal   int
	FemaleTotal int
	Freq        map[string]struct{ Female, Male float64 }
	Total       int
}

//go:embed charfreq.csv
var charFreqData string

// NewGuesser creates a new instance of Guesser.
func NewGuesser() (*Guesser, error) {
	guesser := &Guesser{
		Freq: make(map[string]struct{ Female, Male float64 }),
	}

	err := guesser.loadModel()
	if err != nil {
		return nil, err
	}

	return guesser, nil
}

// loadModel loads the character frequency model from the embedded CSV data.
func (g *Guesser) loadModel() error {
	reader := csv.NewReader(strings.NewReader(charFreqData))
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		char := record[0]
		male := parseFloat(record[1])
		female := parseFloat(record[2])

		g.MaleTotal += int(male)
		g.FemaleTotal += int(female)
		g.Freq[char] = struct{ Female, Male float64 }{
			Female: female,
			Male:   male,
		}
	}

	g.Total = g.MaleTotal + g.FemaleTotal

	for char := range g.Freq {
		female := g.Freq[char].Female
		male := g.Freq[char].Male
		g.Freq[char] = struct{ Female, Male float64 }{
			Female: female / float64(g.FemaleTotal),
			Male:   male / float64(g.MaleTotal),
		}
	}

	return nil
}

// Guess predicts the gender for the given name.
// It returns the predicted gender ("male", "female", or "unknown") and the probability.
func (g *Guesser) Guess(name string) (string, float64) {
	firstname := name[3:] // Assuming the name format is "姓名" (e.g., "赵本山")
	for _, char := range firstname {
		if !isChineseChar(char) {
			panic("姓名必须为中文")
		}
	}

	pf := g.probForGender(firstname, 0)
	pm := g.probForGender(firstname, 1)

	if pm > pf {
		return "male", pm / (pm + pf)
	} else if pm < pf {
		return "female", pf / (pm + pf)
	} else {
		return "unknown", 0
	}
}

// probForGender calculates the probability of a given gender for the given firstname.
func (g *Guesser) probForGender(firstname string, gender int) float64 {
	p := float64(g.FemaleTotal) / float64(g.Total)
	for _, char := range firstname {
		if gender == 0 {
			p *= g.Freq[string(char)].Female
		} else if gender == 1 {
			p *= g.Freq[string(char)].Male
		}
	}
	return p
}

// parseFloat converts a string to a float64 value.
func parseFloat(s string) float64 {
	result, _ := strconv.ParseFloat(s, 64)
	return result
}

// isChineseChar checks if a rune represents a Chinese character.
func isChineseChar(char rune) bool {
	return utf8.RuneLen(char) == 3
}
