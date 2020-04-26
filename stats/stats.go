package stats

import (
	"sort"
	"strings"
)

type Stats struct {
	PartialName string
	Count       uint64
}

// StatsMap type providing general outline of where the statistics are
// concentrated.
type StatsMap struct {
	Stats     map[string]*Stats
	Threshold uint64
}

// New construct a new StatsMap, provide the threshold to descend into more
// detail in the stats.
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
				// FIXME: we could probably do with tweaking
				// the way we decide to descend
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

// Summary return a slice of the stats sorted by count
// Note that the results are not precise counts.
func (m *StatsMap) Summary() []Stats {
	// copy and sort
	stats := make([]Stats, 0)
	i := 0
	for _, v := range m.Stats {
		if v.Count >= m.Threshold {
			stats = append(stats, *v)
		}
		i++
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].Count < stats[j].Count })
	return stats
}
