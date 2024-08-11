package Phasmo

import (
	"encoding/json"
	"os"
)

/*
items = Evidence
NPC = Ghosts
Start New Job sets all items/NPC to 0
If an item/NPC is ruled out then it becomes -1
If an item/NPC is found or selected it becoems 1
*/

const (
	GHOSTS   = 24
	EVIDENCE = 7
)

type Evidence struct {
	Id   int
	Name string
}

type Ghost struct {
	Id         int
	Name       string
	EvidenceId []int
	Notes      []string
}

// Key is the Id of the item/NPC
// Value is the state of the item/NPC
type Investigation struct {
	GMap map[int]Ghost
	EMap map[int]int
}

// Used if reset button is pressed or a new page is loaded.
// Sets all item/NPC as possible and clears all ruled out item/NPC
func StartNewInvestigation() Investigation {
	return Investigation{GMap: freshMapGhosts(), EMap: freshMapItems()}
}

func freshMapGhosts() map[int]Ghost {
	G := make(map[int]Ghost, GHOSTS)
	for i := range G {
		G[i] = getGhostById(i)
	}
	return G
}

func freshMapItems() map[int]int {
	E := make(map[int]int, EVIDENCE)
	for i := range E {
		E[i] = 0
	}
	return E //Ghost or Evidence depends what const is passed inSS
}

func getGhostById(id int) Ghost {
	return readGhostJson()[id]
}

func readGhostJson() []Ghost {
	gJson, err := os.ReadFile("public/Database/Phasmo/Ghosts.json")
	if err != nil {
		panic(err)
	}
	gArr := make([]Ghost, GHOSTS)
	err = json.Unmarshal(gJson, &gArr)
	if err != nil {
		panic(err)
	}
	return gArr
}

func readEvidenceJson() []Evidence {
	eJson, err := os.ReadFile("public/Database/Phasmo/Evidence.json")
	if err != nil {
		panic(err)
	}
	eArr := make([]Evidence, EVIDENCE)
	err = json.Unmarshal(eJson, &eArr)
	if err != nil {
		panic(err)
	}
	return eArr
}

// state refers to if the player or server has ruled it out or if its still an option
func (job Investigation) UpdateEvidence(id int, state int) {
	job.EMap[id] = state
}

func (job Investigation) UpdateGhost(id int, state bool) {
	job.GMap[id] = state
}

func (job Investigation) GetGhostDisplay() []Ghost {
	var display []Ghost
	ghosts := readGhostJson()
	for key, value := range job.GMap {
		if value <= 0 {
			display = append(display, ghosts[key])
		}
	}
	return display
}

func (job Investigation) GetEvidenceDisplay() (Found, Unselected, CrossedOut []Evidence) {
	evidence := readEvidenceJson()
	for key, value := range job.EMap {
		if value == 0 {
			Unselected = append(Unselected, evidence[key])
		}
		if value == 1 {
			Found = append(Found, evidence[key])
		}
		if value == -1 {
			CrossedOut = append(CrossedOut, evidence[key])
		}
	}
	return Found, Unselected, CrossedOut
}
