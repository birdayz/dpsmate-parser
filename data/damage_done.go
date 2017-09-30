package data

type damageDoneData struct {
	hits         []damageDoneAbility
	totalDamage  uint64
	numberOfHits int
}

type damageDoneAbility struct {
	damage  uint64
	ability int
}
