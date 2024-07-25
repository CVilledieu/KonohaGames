package Phasmo

type Evidence struct {
	EvidenceName string
}

type Ghost struct {
	GhostName string
	Tips      string
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
	names := []string{"Spirit", "Wraith", "Phantom", "Poltergeist", "Banshee", "Jinn", "Oni"}
	tip := []string{"Spooky", "Spooky", "Spooky", "Spooky", "Spooky", "Spooky", "Spooky"}
	AllGhost := make([]Ghost, len(names))
	for i := range AllGhost {

		AllGhost[i] = Ghost{GhostName: names[i], Tips: tip[i]}
	}
	return AllGhost
}

func GetBaseInfo() BaseInfo {
	return BaseInfo{Evidences: GetEvidence(), Ghosts: GetGhosts()}
}
