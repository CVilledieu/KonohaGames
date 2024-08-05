package Phasmo

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	GHOSTCOUNT = 24
)

type _Evidence struct {
	Name string
}

type _Ghost struct {
	Id       string
	Name     string
	Evidence []string
	Tips     []string
}

type BaseInfo struct {
	Evidences []_Evidence
	Ghosts    []_Ghost
}

func GetAllEvidence() []_Evidence {
	names := []string{"Spirit Box", "Freezing Temps", "Ultraviolet", "Ghost Orbs", "EMF", "Ghost Writing", "D.O.T.S."}
	AllEvidence := make([]_Evidence, len(names))
	for i, n := range names {

		AllEvidence[i] = _Evidence{Name: n}
	}
	return AllEvidence
}

func GetAllGhosts() []_Ghost {
	AllGhost := make([]_Ghost, GHOSTCOUNT)
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

// State indicates if its 1 Neither ruled out or found, 2 Found, or 3 Ruled out
type Evidence struct {
	Name  string
	Ghost []Ghost
	State int
}

type Ghost struct {
	Name     string
	State    int
	Evidence []Evidence
}

type Investigation struct {
	Evidence []_Evidence
	Ghosts   []_Ghost
}
