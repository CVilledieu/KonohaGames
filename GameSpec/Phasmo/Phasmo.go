package Phasmo

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

type Investigation struct {
	Evidence []Evidence
	Ghosts   []Ghost
}

// Used if reset button is pressed or a new page is loaded.
// Sets all ghosts/evidence as possible and clears all ruled out evidence/ghosts
func StartJob() Investigation {
	return Investigation{Ghosts: getAllGhosts()}
}

func getAllGhosts() []Ghost {
	return []Ghost{
		{
			Id:         0,
			Name:       "Spirit",
			EvidenceId: []int{0, 1, 2},
			Tip:        []string{"Shy Boy"},
		},
	}
}

func getAllEvidence() []Evidence {
	return []Evidence{
		{
			Id:      0,
			Name:    "Spirit Box",
			GhostId: []int{0, 1, 2},
		},
	}
}

func getGhostById(id int) Ghost {
	allGhost := getAllGhosts()
	for _, val := range allGhost {
		if val.Id == id {
			return val
		}
	}
	panic("Id passed in didn't match any Ghost")
}

func getEvidenceById(id int) Evidence {
	allEvidence := getAllEvidence()
	for _, val := range allEvidence {
		if val.Id == id {
			return val
		}
	}
	panic("Id passed in doesnt match any Evidence")
}

// Returns all the Ghosts that are listed by the passed in evidence
func getGhostsByEvidence(EV Evidence) []Ghost {
	ghosts := []Ghost{}

	for _, val := range EV.GhostId {
		ghosts = append(ghosts, getGhostById(val))
	}
	return ghosts
}

// Returns all the Ghosts that are listed by the passed in evidence
func getEvidenceByGhosts(g Ghost) []Evidence {
	ev := []Evidence{}

	for _, val := range g.EvidenceId {
		ev = append(ev, getEvidenceById(val))
	}
	return ev
}


//Returns Possible Ghosts based on an array of evidence
//The input array is the evidence either the player or the server has either ruled out 
func Get