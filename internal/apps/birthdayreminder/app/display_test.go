package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadNextBirthday(t *testing.T) {

	const (
		odog  = "nils_21-02.png"
		monki = "monki_25-10.png"
		oehm  = "oehm_17-08.png"
	)

	cases := []struct {
		today time.Time
		want  string
	}{
		{
			today: time.Date(2010, time.Month(10), 28, 12, 0, 0, 0, time.UTC),
			want:  odog,
		},
		{
			today: time.Date(2010, time.Month(10), 18, 8, 0, 0, 0, time.UTC),
			want:  monki,
		},
		{
			today: time.Date(2010, time.Month(2), 22, 2, 0, 0, 0, time.UTC),
			want:  oehm,
		},
	}

	files := []string{oehm, monki, odog}

	for _, testCase := range cases {

		assert.Equal(t, testCase.want, loadNextBirthday(testCase.today, files))
	}

}
