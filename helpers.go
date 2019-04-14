package helpers

import (
	"github.com/gobuffalo/helpers/content"
	"github.com/gobuffalo/helpers/debug"
	"github.com/gobuffalo/helpers/encoders"
	"github.com/gobuffalo/helpers/env"
	"github.com/gobuffalo/helpers/escapes"
	"github.com/gobuffalo/helpers/forms"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/helpers/inflections"
	"github.com/gobuffalo/helpers/iterators"
	"github.com/gobuffalo/helpers/meta"
	"github.com/gobuffalo/helpers/text"
)

var Content = content.New()
var Debug = debug.New()
var Encoders = encoders.New()
var Env = env.New()
var Escapes = escapes.New()
var Forms = forms.New()
var Inflections = inflections.New()
var Iterators = iterators.New()
var Meta = meta.New()
var Text = text.New()

var ALL = func() hctx.Map {
	return hctx.Merge(
		Content,
		Debug,
		Encoders,
		Env,
		Escapes,
		Forms,
		Inflections,
		Iterators,
		Meta,
		Text,
	)
}
