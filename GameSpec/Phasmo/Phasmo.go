package Phasmo

/*
Default is ruled out. Info passed from server to client updates the options to no longer be ruled out

When phasmo Page is loaded StartJob func should be run.
  This sends all ghosts and evidence to the client as possible options
When the client removes an option it updates and sends the updated Item back to the server
  All info is then updated with the info sent back by the server

This should mean that only the possible options need to be tracked instead of the state of all options
*/

const (
	GHOSTS   = 24
	EVIDENCE = 7
)

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
	Ghosts   map[int]bool
	Evidence map[int]bool
}

// Used if reset button is pressed or a new page is loaded.
// Sets all ghosts/evidence as possible and clears all ruled out evidence/ghosts
func StartNewJob() Investigation {
	return Investigation{Ghosts: SetUpInvestigation(GHOSTS), Evidence: SetUpInvestigation(EVIDENCE)}
}

//Provides a clean slate. Used to set up a new Investigation or if a reset button is added
//The Ids are simply 0-23 for ghosts and 0-6 for Evidence
func SetUpInvestigation(t int) map[int]bool {
	GE := make(map[int]bool, t)
	for i := 0; i < t; i++ {
		GE[i] = true
	}
	return GE //Ghost or Evidence depends what const is passed inSS
}

func getGhostById(id int) Ghost {
	panic("Id passed in didn't match any Ghost")
}

func getEvidenceById(id int) Evidence {
	panic("Id passed in doesnt match any Evidence")
}

// Returns all the Ghost Ids that point to that Evidence id
func GetGhostsByEvidence(id int) []int {
	return getEvidenceById(id).GhostId
}

// Returns all the Evidence Ids that point to that ghost id
func getEvidenceByGhosts(id int) []int {
	return getGhostById(id).EvidenceId
}

//state refers to if the player or server has ruled it out or if its still an option
func (job Investigation) UpdateEvidence(id int, state bool) {
	job.Evidence[id] = state
}

func (job Investigation) UpdateGhost(id int, state bool) {
	job.Ghosts[id] = state
}
