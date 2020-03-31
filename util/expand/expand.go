package expand

/* Constructs matchers that use regex and word assocs to expand and connect words */
/* Ping API to check if food word */

// Associations constructs associations, final form of functionality
type Associations struct {
	Words string
	Root  *Matcher
}

// NewAssociations constructs assoc obj
func NewAssociations(allWords string, root *Matcher) *Associations {
	return &Associations{
		Words: allWords,
		Root:  root,
	}
}

// Matcher constructs matching object
type Matcher struct {
	Word      *RgxMpr
	WordSize  int
	Check     *RgxMpr
	CheckSize int
	Next      *Matcher
}

//NewMatcher inits matcher obj
func NewMatcher(word string, checker string, next *Matcher) *Matcher {
	newWord := NewRgxMpr(word)
	newWord.InitMap(word)
	newChecker := NewRgxMpr(checker)
	newChecker.InitMap(checker)
	return &Matcher{
		Word:      newWord,
		WordSize:  len(newWord.Rgxmpr),
		Check:     newChecker,
		CheckSize: len(newChecker.Rgxmpr),
		Next:      next,
	}
}

// Matcher checks if words regex match to food words; if not, discard word
func (m *Matcher) Matcher() string {
	if m.WordSize == 0 || m.CheckSize == 0 {
		return ""
	}
	var rgxChk [][]bool
	rgxChk[m.CheckSize+1][m.WordSize+1] = true
	if m.isMatch(rgxChk) {
		return m.Check.Rgxmpr
	}
	return ""
}

func (m *Matcher) isMatch(rgxChk [][]bool) bool {
	for i := m.CheckSize; i >= 0; i-- {
		for j := m.WordSize - 1; j >= 0; j-- {
			matches := m.charMatch(m.Check.Rgxmpr[i], m.Word.Rgxmpr[j]) && i < m.CheckSize
			if m.Word.Rgxmpr[j+1] == '*' && j+1 < m.WordSize {
				rgxChk[i][j] = rgxChk[i][j+2] || (matches && rgxChk[i+1][j])
			} else {
				rgxChk[i][j] = matches && rgxChk[i+1][j+1]
			}
		}
	}
	return rgxChk[0][0]
}

func (m *Matcher) charMatch(c1 byte, c2 byte) bool {
	if c1 == c2 {
		return true
	}
	if c2 == '.' {
		return true
	}
	return false
}

//RgxMpr nodes to represent letters
type RgxMpr struct {
	Rgxmpr string
}

// NewRgxMpr Sets root to starting letter of word
func NewRgxMpr(rgxmpr string) *RgxMpr {
	return &RgxMpr{
		Rgxmpr: rgxmpr,
	}
}

// InitMap sets up DAG
func (root *RgxMpr) InitMap(word string) {
	var curr []byte
	j := 0
	curr[j] = word[0]
	j++
	for i := 1; i < len(word); i++ {
		curr[j] = '.'
		j++
		curr[j] = word[i]
		j++
		curr[j] = '*'
		j++
	}
}
