package content

import (
	"testing"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/helpers/helptest"
	"github.com/stretchr/testify/require"
)

func Test_ContentFor(t *testing.T) {
	r := require.New(t)

	in := "<button>hi</button>"
	hc := helptest.NewContext()
	hc.BlockContextFn = func(c hctx.Context) (string, error) {
		return in, nil
	}

	cf := hc.New().(*helptest.HelperContext)
	ContentFor("buttons", hc)
	s, err := ContentOf("buttons", hctx.Map{}, cf)
	r.NoError(err)
	r.Contains(s, in)
}
