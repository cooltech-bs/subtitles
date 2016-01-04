package subtitles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterHTML(t *testing.T) {

	var in = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"<b>GO NINJA!</b>", "NINJA&nbsp;GO!"},
		},
	}

	var expected = []caption{
		{
			1,
			makeTime(0, 0, 4, 630),
			makeTime(0, 0, 6, 18),
			[]string{"GO NINJA!", "NINJA GO!"},
		},
	}

	assert.Equal(t, expected, filterHTML(in))
}
