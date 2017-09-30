package main

import (
	"fmt"
	"log"

	"github.com/birdayz/dpsmate-parser/data"

	"flag"

	lua "github.com/yuin/gopher-lua"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "Path to DPSMate.lua")
}

func main() {
	flag.Parse()

	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(path); err != nil {
		panic(err)
	}

	val := L.GetGlobal("DPSMateDamageDone")
	switch x := val.(type) {
	case *lua.LTable:
		// each of those is a segment
		segments := parseDamageDone(x)

		log.Printf("Found %v segments", len(segments))
		// var segments []data.Segment

	}
}

func parseDamageDone(table *lua.LTable) (segments []data.Segment) {
	table.ForEach(func(a lua.LValue, b lua.LValue) { // a = number, b = table
		if players, ok := b.(*lua.LTable); ok {
			fmt.Println("Found", players.Len(), " players in segment")
			players.ForEach(func(playerNumber lua.LValue, abilitiesTable lua.LValue) {
				// each one is a player
				if damageOfPlayerTable, ok := abilitiesTable.(*lua.LTable); ok {
					totalDmgOfPlayer := getDamageOfPlayer(damageOfPlayerTable)
					fmt.Println("Player deals ", totalDmgOfPlayer.TotalDamage, " damage")
				}

			})

		}

	})
	return
}

// parse segment function

// operates on player's table
func getDamageOfPlayer(table *lua.LTable) (dmg data.Damage) { // Damage from ONE player
	totalDamage := 0
	//var damageByAbility map[string]int
	table.ForEach(func(abilityId lua.LValue, abilityDamages lua.LValue) {
		if abilityDamagesTable, ok := abilityDamages.(*lua.LTable); ok {
			_, totalDamageOfAbility := extractDamageFromAbilityTable(abilityDamagesTable)
			totalDamage += totalDamageOfAbility
			//fmt.Println("Ability", abilityId.String(), "deals ", damagesOfAbility, " damage, total=", totalDamageOfAbility)

		} else {
			// IGNORE for now	fmt.Println("No table...must be i total dmg to xx")
		}
	})
	dmg.TotalDamage = totalDamage
	return
}

func extractDamageFromAbilityTable(table *lua.LTable) (dmg []int, total int) {
	numberOfEntries := table.RawGetInt(1)
	// TODo check if it's a number
	count := int(lua.LVAsNumber(numberOfEntries))

	if count < 1 {
		return
	}

	maxNumber := 2 + count

	for i := 2; i < maxNumber; i++ {
		damageValue := int(lua.LVAsNumber(table.RawGetInt(i)))
		dmg = append(dmg, damageValue)
	}

	total = int(lua.LVAsNumber(table.RawGetInt(maxNumber)))

	return
}
