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

	KeyLeft  = 0
	KeyDown  = 1
	KeyUp    = 2
	KeyRight = 3

	TotGlyphSets = 2  // total number of glyph sets
	TotGlyphs    = 15 // total number of glyphs per set
	TotKeyGlyphs = 4  // Number of key glyphs
)

var (
	// The glyph sets we support.
	glyphs = [TotGlyphSets][TotGlyphs]string{
		ASCII:   {"*", "1", "2", "3", "4", "5", "6", "7", "8", "?", "@", "X", ":-(", ":-)", ":-O"},
		UNICODE: {"✱", "1", "2", "3", "4", "5", "6", "7", "8", "?", "⚑", "✘", "(╥‸╥)", "(^ᴗ^)", "(⊙_⊙)"},
	}

	// The key glyphs.
	keyGlyphs = [TotGlyphSets][TotKeyGlyphs]string{
		ASCII:   {"left", "down", "up", "right"},
		UNICODE: {"←", "↓", "↑", "→"},
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

// Get returns a single glyph from the default set.
func Get(glyph int) string {
	return glyphs[defaultSet][glyph]
}

// GetKey returns a key glyph from the default set.
func GetKey(glyph int) string {
	return keyGlyphs[defaultSet][glyph]
}
