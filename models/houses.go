package models

type AllRaashis string

type Nakshatra string

type RaashiAspect struct {
	House AllRaashis
}

type SpecialAspectsMap map[AllRaashis][]RaashiAspect

type Raashyadhipati string

type SpecialPosition map[Raashyadhipati]AllRaashis

type ReverseSpecialPosition map[AllRaashis]Raashyadhipati

type RaashyadhipatiMapping map[string]Raashyadhipati

type ReverseRaashyadhipatiMapping map[Raashyadhipati][]AllRaashis
