package Phasmo

const (
	GHOSTCOUNT = 24
)

type Evidence struct {
	Id    string
	Name  string
	Ghost []Ghost
}

// Each Ghost has 3 Evidences that are a unique set. No 2 ghosts have the same set of Evidence
// Tip is extra info about the behaviour of the ghost
type Ghost struct {
	Id       string
	Name     string
	Evidence []Evidence
	Tip      []string
}

type Investigation struct {
	Evidence []Evidence
	Ghosts   []Ghost
}

// Used if reset button is pressed or a new page is loaded.
// Sets all ghosts/evidence as possible and clears all ruled out evidence/ghosts
func StartJob() Investigation {
	return Investigation{}
}

func (inv *Investigation) RuleOutEvidence(evid Evidence) {

}

func (inv *Investigation) SetPossibleGhosts() {

}
