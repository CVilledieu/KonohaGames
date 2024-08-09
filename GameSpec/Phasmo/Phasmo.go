package Phasmo

/*
Default is ruled out. Info passed from server to client updates the options to no longer be ruled out

When phasmo Page is loaded StartJob func should be run.
  This sends all ghosts and evidence to the client as possible options
When the client removes an option it updates and sends the updated Item back to the server
  All info is then updated with the info sent back by the server

This should mean that only the possible options need to be tracked instead of the state of all options
*/

type Evidence struct {
	Id      int
	Name    string
	GhostId []int
}

// Each Ghost has 3 Evidences that are a unique set. No 2 ghosts have the same set of Evidence
// Tip is extra info about the behaviour of the ghost
type Ghost struct {
	Id         int
	Name       string
	EvidenceId []int
	Tip        []string
}

//Use IDs to search and deal with everything
type Investigation struct {
	Ghosts   []int
	Evidence []int
}

// Used if reset button is pressed or a new page is loaded.
// Sets all ghosts/evidence as possible and clears all ruled out evidence/ghosts
func StartNewJob() Investigation {
	return Investigation{Ghosts: GetAllGhostIds(), Evidence: GetAllEvidenceIds()}
}

//Provides a clean slate. Used to set up a new Investigation or if a reset button is added
func GetAllGhostIds() []int {
	ghosts := make([]int, 24)
	for i, _ := range ghosts {
		ghosts[i] = i
	}
	return ghosts
}
func GetAllEvidenceIds() []int {
	evidence := make([]int, 7)
	for i, _ := range evidence {
		evidence[i] = i
	}
	return evidence
}

//
func getGhostById(id int) Ghost {
	panic("Id passed in didn't match any Ghost")
}

func getEvidenceById(id int) Evidence {
	panic("Id passed in doesnt match any Evidence")
}

// Returns all the Ghosts that are listed by the passed in evidence
func GetGhostsByEvidence(id int) []int {
	return getEvidenceById(id).GhostId
}

// Returns all the Ghosts that are listed by the passed in evidence
func getEvidenceByGhosts(id int) []int {
	return getGhostById(id).EvidenceId
}

// Returns Possible Ghosts based on an array of evidence
// The input array is the evidence either the player or the server has either ruled out
func GetPossibleGhosts(EVarr []Evidence) (pGhost []Ghost) {
	notPossible := make(map[int]int)
	for _, EV := range EVarr {
		for _, gId := range EV.GhostId {
			notPossible[gId] += 1
		}
	}
	for key := range notPossible {
		pGhost = append(pGhost, getGhostById(key))
	}
	return pGhost
}

func (job Investigation) AddEvidence(id int) {
	evidence := getEvidenceById(id)
	checkIfInArr()
}

func (job Investigation) AddGhost(id int) {

}

func (job Investigation) RemoveEvidence(id int) {

}

func (job Investigation) RemoveGhost(id int) {

}

func checkIfInArr(arr []any, id int) {

}
