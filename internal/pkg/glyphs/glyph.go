package glyphs

import "fmt"

const (
	ASCII   = 0 // ASCII glyphs supported on all terminals
	UNICODE = 1 // Unicode glyphs

	GMine  = 0  // Mine glyph index
	GOne   = 1  // Number one glyph index
	GTwo   = 2  // Number two glyph index
	GThree = 3  // Number three glyph index
	GFour  = 4  // Number four glyph index
	GFive  = 5  // Number five glyph index
	GSix   = 6  // Number six glyph index
	GSeven = 7  // Number seven glyph index
	GEight = 8  // Number eight glyph index
	GQMark = 9  // Question mark glyph index
	GCheck = 10 // Check mark glyph index
	GXMark = 11 // X mark glyph index
	SSad   = 12 // Sad status index
	SHappy = 13 // Happy status index
	SDoubt = 14 // Doubt status index

	TotGlyphSets = 2  // total number of glyph sets
	TotGlyphs    = 15 // total number of glyphs per set
)

var (
	// The glyph sets we support.
	glyphs = [TotGlyphSets][TotGlyphs]string{
		ASCII:   {"*", "1", "2", "3", "4", "5", "6", "7", "8", "?", "F", "X", ":-(", ":-)", ":-O"},
		UNICODE: {"✱", "1", "2", "3", "4", "5", "6", "7", "8", "?", "⚑", "✘", "(╥‸╥)", "(^ᴗ^)", "(⊙_⊙)"},
	}

	// The default set. It can be changed through the SetDefault() function.
	defaultSet = ASCII
)

// GetSet returns the full glyph set of the selected type.
func GetSet(gs int) [TotGlyphs]string {
	return glyphs[gs]
}

// SetDefault sets the default glyph set.
func SetDefault(gs int) error {
	if gs < 0 || gs >= TotGlyphSets {
		return fmt.Errorf("glyph set %d not supported", gs)
	}
	defaultSet = gs
	return nil
}

// Get returns a single glyph from the default set. If the argument is not
// supported, an empty string is returned.
func Get(glyph int) string {
	if glyph < 0 || glyph >= TotGlyphs {
		return "N"
	}
	return glyphs[defaultSet][glyph]
}
