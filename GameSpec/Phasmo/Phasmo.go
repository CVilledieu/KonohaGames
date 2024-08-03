package Phasmo

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	GHOSTCOUNT = 24
)

type Evidence struct {
	Name string
}

type Ghost struct {
	Id       string
	Name     string
	Evidence []string
	Tips     []string
}

type BaseInfo struct {
	Evidences []Evidence
	Ghosts    []Ghost
}

func GetAllEvidence() []Evidence {
	names := []string{"Spirit Box", "Freezing Temps", "Ultraviolet", "Ghost Orbs", "EMF", "Ghost Writing", "D.O.T.S."}
	AllEvidence := make([]Evidence, len(names))
	for i, n := range names {

		AllEvidence[i] = Evidence{Name: n}
	}
	return AllEvidence
}

func GetAllGhosts() []Ghost {
	AllGhost := make([]Ghost, GHOSTCOUNT)
	JReader, err := os.ReadFile("public/json/ghosts.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(JReader, &AllGhost)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return AllGhost
}

func GetBaseInfo() BaseInfo {
	return BaseInfo{Evidences: GetAllEvidence(), Ghosts: GetAllGhosts()}
}

type Hunt struct {
	PossibleEvidence []Evidence
	RuledOutEvidence []Evidence
	PossibleGhosts   []Ghost
	RuledOutGhosts   []Ghost
}

func getNewHunt() Hunt {
	return Hunt{PossibleEvidence: GetAllEvidence(), PossibleGhosts: GetAllGhosts()}
}

func (h Hunt) RuleOutEvidence(ev Evidence) {
	h.RuledOutEvidence = append(h.RuledOutEvidence, ev)
}

func (h Hunt) RuleOutGhost(g Ghost) {
	h.RuledOutGhosts = append(h.RuledOutGhosts, g)
}
