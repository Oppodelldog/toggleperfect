package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadNextBirthday(t *testing.T) {
	const (
		odog  = "data/nils_21-02.png"
		monki = "data/monki_25-10.png"
		oehm  = "data/oehm_17-08.png"
	)

	cases := []struct {
		today time.Time
		want  Birthday
	}{
		{
			today: time.Date(2010, time.Month(10), 28, 12, 0, 0, 0, time.UTC),
			want:  Birthday{Filename: "nils_21-02.png", FilePath: odog, Name: "nils", Day: 21, Month: 2},
		},
		{
			today: time.Date(2010, time.Month(10), 18, 8, 0, 0, 0, time.UTC),
			want:  Birthday{Filename: "monki_25-10.png", FilePath: monki, Name: "monki", Day: 25, Month: 10},
		},
		{
			today: time.Date(2010, time.Month(2), 22, 2, 0, 0, 0, time.UTC),
			want:  Birthday{Filename: "oehm_17-08.png", FilePath: oehm, Name: "oehm", Day: 17, Month: 8},
		},
	}

	files := []string{oehm, monki, odog}

	for _, testCase := range cases {
		assert.Equal(t, testCase.want, loadNextBirthday(testCase.today, files))
	}
}
