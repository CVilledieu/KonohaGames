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
	EvidenceName string
}

type Ghost struct {
	Id            string
	GhostName     string
	GhostEvidence []string
	Tips          []string
}

type BaseInfo struct {
	Evidences []Evidence
	Ghosts    []Ghost
}

func GetEvidence() []Evidence {
	names := []string{"Spirit Box", "Freezing Temps", "Ultraviolet", "Ghost Orbs", "EMF", "Ghost Writing", "D.O.T.S."}
	AllEvidence := make([]Evidence, len(names))
	for i, n := range names {

		AllEvidence[i] = Evidence{EvidenceName: n}
	}
	return AllEvidence
}

func GetGhosts() []Ghost {
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
	return BaseInfo{Evidences: GetEvidence(), Ghosts: GetGhosts()}
}
