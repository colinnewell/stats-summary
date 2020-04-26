package stats

import (
	"sort"
	"strings"
)

type Stats struct {
	PartialName string
	Count       uint64
}

type StatsMap struct {
	Stats     map[string]*Stats
	Threshold uint64
}

func New(threshold uint64) StatsMap {
	return StatsMap{Stats: make(map[string]*Stats), Threshold: threshold}
}

func (m *StatsMap) RecordStats(stat string) {
	if !m.incrementStat(stat) {
		parts := strings.Split(stat, ".")
		var incremented bool
		for i := len(parts) - 1; i >= 0; i-- {
			part := strings.Join(parts[0:i], ".")
			if m.incrementStat(part) {
				incremented = true
				if m.Stats[part].Count >= m.Threshold {
					part = strings.Join(parts[0:i+1], ".")
					m.addStat(part)
				}
				// check if we've hit the threshold
				// if we have, store the next level too
				break
			}
		}
		if !incremented {
			m.addStat(parts[0])
		}
	}
}

func (m *StatsMap) incrementStat(stat string) bool {
	_, ok := m.Stats[stat]
	if ok {
		m.Stats[stat].Count++
	}
	return ok
}

func (m *StatsMap) addStat(stat string) {
	m.Stats[stat] = &Stats{PartialName: stat, Count: 1}
}

func (m *StatsMap) Summary() []Stats {
	// copy and sort
	stats := make([]Stats, len(m.Stats))
	i := 0
	for _, v := range m.Stats {
		stats[i] = *v
		i++
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].Count < stats[j].Count })
	return stats
}
