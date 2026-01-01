package main

import "Pokedexcli/internal/pokeapi"

type Pokedex struct {
	byId   map[int]*pokeapi.Pokemon
	byName map[string]*pokeapi.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		byId:   make(map[int]*pokeapi.Pokemon),
		byName: make(map[string]*pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(pk pokeapi.Pokemon) {
	// We store a pointer so both maps point to the SAME object in memory
	ref := &pk
	p.byId[pk.ID] = ref
	p.byName[pk.Name] = ref
}

func (p *Pokedex) GetByName(name string) (*pokeapi.Pokemon, bool) {
	val, ok := p.byName[name]
	return val, ok
}

func (p *Pokedex) GetByID(id int) (*pokeapi.Pokemon, bool) {
	val, ok := p.byId[id]
	return val, ok
}

func (p *Pokedex) Delete(name string) {
	if pk, ok := p.byName[name]; ok {
		delete(p.byId, pk.ID)
		delete(p.byName, pk.Name)
	}
}
