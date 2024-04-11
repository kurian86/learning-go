package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstTeam, secondTeam string, firstScore, secondScore int) {
	if firstScore == secondScore {
		return
	}

	if firstScore > secondScore {
		l.Wins[firstTeam]++
	} else {
		l.Wins[secondTeam]++
	}
}

func (l League) Ranking() []string {
	teamNames := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		teamNames = append(teamNames, team.Name)
	}

	sort.Slice(teamNames, func(i, j int) bool {
		return l.Wins[teamNames[i]] > l.Wins[teamNames[j]]
	})

	return teamNames
}

func main() {
	league := League{
		Name: "Premier League",
		Teams: []Team{
			{
				Name:    "Manchester City",
				Players: []string{"Erling Haaland", "Rodri Hernández", "Oscar Bobb"},
			},
			{
				Name:    "Manchester United",
				Players: []string{"Antony", "Amad Diallo", "Casemiro"},
			},
			{
				Name:    "Chelsea",
				Players: []string{"Cole Palmer", "Enzo Fernández", "Conor Gallagher"},
			},
			{
				Name:    "Arsenal",
				Players: []string{"Bukayo Saka", "Declan Ride", "Grabiel Jesus"},
			},
		},
		Wins: make(map[string]int, 4),
	}

	league.MatchResult(league.Teams[0].Name, league.Teams[1].Name, 1, 0)
	league.MatchResult(league.Teams[0].Name, league.Teams[2].Name, 0, 2)
	league.MatchResult(league.Teams[0].Name, league.Teams[3].Name, 0, 3)
	league.MatchResult(league.Teams[1].Name, league.Teams[0].Name, 1, 0)
	league.MatchResult(league.Teams[1].Name, league.Teams[2].Name, 2, 0)
	league.MatchResult(league.Teams[1].Name, league.Teams[3].Name, 0, 3)
	league.MatchResult(league.Teams[2].Name, league.Teams[0].Name, 1, 0)
	league.MatchResult(league.Teams[2].Name, league.Teams[1].Name, 2, 0)
	league.MatchResult(league.Teams[2].Name, league.Teams[3].Name, 3, 0)
	league.MatchResult(league.Teams[3].Name, league.Teams[0].Name, 1, 0)
	league.MatchResult(league.Teams[3].Name, league.Teams[1].Name, 2, 0)
	league.MatchResult(league.Teams[3].Name, league.Teams[2].Name, 3, 0)

	fmt.Println(league.Ranking())
}
