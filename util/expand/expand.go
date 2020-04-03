package expand

/* Constructs matchers that use regex and word assocs to expand and connect words */
/* Ping API (maybe localized, but potentially Python) to check if food word */

// An association is the end product - constructed of words from an ingredient that are food and logically go together
// For example, peanut butter is an association
//Essentially a linked list in which assocations are pointers to the next word

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

// The way this will work: matchers have next pointers to point to the word they are most greatly associated with
// For example, beef -> butter
// Potential issues: words like peanut -> butter because of the prevalence of peanut butter

// Matcher constructs matching object
type Matcher struct {
	Word     *RgxMpr
	WordSize int

	Checker   string
	Check     *RgxMpr
	CheckSize int

	Next *Matcher
	Prev *Matcher

	Matched string
}

//NewMatcher inits matcher obj
func NewMatcher(word string, checker string, next *Matcher, prev *Matcher) *Matcher {
	newWord := NewRgxMpr(word)
	newWord.InitMap(word)
	newChecker := NewRgxMpr(checker)
	newChecker.InitMap(checker)
	return &Matcher{
		Word:      newWord,
		WordSize:  len(newWord.Rgxmpr),
		Check:     newChecker,
		Checker:   checker,
		CheckSize: len(newChecker.Rgxmpr),
		Next:      next,
		Prev:      prev,
	}
}

// Matcher checks if words regex match to food words; if not, discard word
func (m *Matcher) Matcher() {
	if m.WordSize == 0 || m.CheckSize == 0 {
		return
	}
	var rgxChk [][]bool
	rgxChk[m.CheckSize+1][m.WordSize+1] = true
	if m.isMatch(rgxChk) {
		m.Matched = m.Checker
	}
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

func (m *Matcher) isAssociated(m2 *Matcher) bool {
	return associationAPI(m.Word.Rgxmpr, m2.Word.Rgxmpr)
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

//RgxMpr represents matching
type RgxMpr struct {
	Rgxmpr string
}

// NewRgxMpr Constructs RgxMpr obj
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
	root.Rgxmpr = string(curr)
}

func associationAPI(word1 string, word2 string) bool {
	return false
}
