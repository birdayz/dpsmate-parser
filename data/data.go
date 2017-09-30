package data

type listing struct {
	segments []Segment
}

type Segment struct {
	playerStats []playerStats
	// Time
}

type playerStats struct {
	data []*genericData
}

type genericData interface {
	toString()
}

type Damage struct {
	TotalDamage     int
	DamageByAbility map[string][]int
}

// Metric - Segment - Source / Affected Unit - ? - ?
