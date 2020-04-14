package expand

import "strings"

/* Constructs matchers that use regex and word assocs to expand and connect words */
/* The end-product of this package is an Association, which builds a lists of the lists of words that are associated */
/* Ping API (maybe localized, but potentially Python) to check if words are commonly associated in recipes */

// Associations associations
type Associations struct {
	Matched []string
	Roots   []*AssociationNode
}

// NewAssociations constructs assoc obj
func NewAssociations(alphabet []string, dataset []string) *Associations {
	var AllMatchers []string

	for _, sentence := range alphabet {
		words := strings.Split(sentence, " ")
		matchers := ""
		for _, word := range words {
			matcher := NewMatcher(word, "", dataset...)
			if matcher.Matched != "" {
				matchers = matchers + matcher.Matched + " "
			}
		}
		if matchers != "" {
			matchers = strings.TrimSpace(matchers)
			AllMatchers = append(AllMatchers, matchers)
		}
	}

	return &Associations{
		Matched: AllMatchers,
	}
}

//Associate builds association from Association object and includes link of all ingredients
func (a *Associations) Associate(threshold float64) int {
	if len(a.Matched) == 0 {
		return 0
	}
	var roots []*AssociationNode
	for i := 0; i < len(a.Matched); i++ {
		root := NewAssociationNode(a.Matched[i])
		var associatedWords []string

		associatedWords = append(associatedWords, root.Word)
		root.AssociatedWords = associatedWords

		curr := root
		for j := i + 1; j < len(a.Matched); j++ {

			next := NewAssociationNode(a.Matched[j])
			if curr.isAssociated(next) < threshold {
				root.AssociatedWords = append(root.AssociatedWords, a.Matched[j])
				curr = next
			}
		}
		roots = append(roots, root)
	}

	completeSetRoot := NewAssociationNode(a.Matched[0])
	var associatedWords []string

	associatedWords = append(associatedWords, completeSetRoot.Word)
	completeSetRoot.AssociatedWords = associatedWords

	for i := 1; i < len(a.Matched); i++ {
		completeSetRoot.AssociatedWords = append(completeSetRoot.AssociatedWords, a.Matched[i])
	}
	roots = append(roots, completeSetRoot)

	a.Roots = roots
	return 1
}

//AssociationNode node
type AssociationNode struct {
	Word            string
	AssociatedWords []string
}

//NewAssociationNode node
func NewAssociationNode(word string) *AssociationNode {
	return &AssociationNode{
		Word: word,
	}
}

// The way this will work: matchers to point to the word they are most greatly associated with
// For example, beef -> butter
// Potential issues: words like peanut -> butter because of the prevalence of peanut butter

// Matcher constructs matching object
type Matcher struct {
	Word     *Mapper
	WordSize int

	Checkers []string
	Checks   []*Mapper

	Matched string
}

//NewMatcher inits matcher obj
func NewMatcher(word, matcher string, checkers ...string) *Matcher {
	newWord := NewMapper(word)
	newWord.InitMap()

	newCheckers := []*Mapper{}

	for _, checker := range checkers {
		newChecker := NewMapper(checker)
		newCheckers = append(newCheckers, newChecker)
	}
	return &Matcher{
		Word:     newWord,
		WordSize: len(newWord.Map),
		Checks:   newCheckers,
		Checkers: checkers,
		Matched:  matcher,
	}
}

// Matcher checks if words matches with food words; if not, discard word - want to implement this more efficiently
func (m *Matcher) Matcher() {
	for _, check := range m.Checks {
		if m.WordSize == 0 {
			return
		}
		var rgxChk [][]bool
		rgxChk[len(check.Map)+1][m.WordSize+1] = true
		if m.isMatch(rgxChk, check.Map) {
			m.Matched = check.Map
		}
	}
}

func (m *Matcher) isMatch(rgxChk [][]bool, check string) bool {
	for i := len(check); i >= 0; i-- {
		for j := m.WordSize - 1; j >= 0; j-- {
			matches := m.charMatch(check[i], m.Word.Map[j]) && i < len(check)
			if m.Word.Map[j+1] == '*' && j+1 < m.WordSize {
				rgxChk[i][j] = rgxChk[i][j+2] || (matches && rgxChk[i+1][j])
			} else {
				rgxChk[i][j] = matches && rgxChk[i+1][j+1]
			}
		}
	}
	return rgxChk[0][0]
}

func (an *AssociationNode) isAssociated(an2 *AssociationNode) float64 {
	return association(an.Word, an2.Word)
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

//Mapper represents matching
type Mapper struct {
	Map string
}

// NewMapper Constructs RgxMpr obj
func NewMapper(word string) *Mapper {
	return &Mapper{
		Map: word,
	}
}

// InitMap sets up map
func (mapping *Mapper) InitMap() {
	var curr []byte
	j := 0
	curr[j] = mapping.Map[0]
	j++
	for i := 1; i < len(mapping.Map); i++ {
		curr[j] = '.'
		j++
		curr[j] = mapping.Map[i]
		j++
		curr[j] = '*'
		j++
	}
	mapping.Map = string(curr)
}

func association(word1 string, word2 string) float64 {
	return 0
}
