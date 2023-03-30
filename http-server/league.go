package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}

func (league League) Find(name string) *Player {
	for i, p := range league {
		if p.Name == name {
			return &league[i]
		}
	}
	return nil
}
