package pages

import (
	"strings"

	"workshop/internal/domain/sighting"
	"workshop/web/views/components"
)

type (
	homeView struct {
		Summary  sighting.Summary
		Featured []animal
		Floating []animal
		Places   []sighting.Tally
		Hero     sighting.Sighting
		Counties int
	}
	animal struct {
		Name     string
		Count    int
		Sighting sighting.Sighting
	}
)

// newHomeView works out the home page's figures and photos in one pass over the sightings.
func newHomeView(sightings []sighting.Sighting) homeView {
	withPhoto := make(map[string]sighting.Sighting)
	for _, s := range sightings {
		name := components.Name(s.Species)
		if _, ok := withPhoto[name]; !ok && components.Photo(s).URL != "" {
			withPhoto[name] = s
		}
	}

	var animals []animal
	counts := sighting.CountBy(sightings, func(s sighting.Sighting) string { return components.Name(s.Species) })
	for _, t := range sighting.Top(counts, len(counts)) {
		if s, ok := withPhoto[t.Name]; ok {
			animals = append(animals, animal{Name: t.Name, Count: t.Count, Sighting: s})
		}
	}

	counties := sighting.CountBy(sightings, func(s sighting.Sighting) string { return s.Location.StateProvince })
	v := homeView{
		Summary:  sighting.Summarise(sightings),
		Featured: animals[:min(3, len(animals))],
		Floating: animals[min(3, len(animals)):min(6, len(animals))],
		Places:   sighting.Top(counties, 6),
		Counties: len(counties),
	}
	v.Hero = heroSighting(animals)

	return v
}

// heroSighting prefers an elephant for the banner, then any animal with a photo.
func heroSighting(animals []animal) sighting.Sighting {
	for _, want := range []string{"elephant", "giraffe", "zebra"} {
		for _, a := range animals {
			if strings.Contains(strings.ToLower(a.Name), want) {
				return a.Sighting
			}
		}
	}
	if len(animals) > 0 {
		return animals[0].Sighting
	}
	return sighting.Sighting{}
}

// featuredIcon gives each tall card its own icon.
func featuredIcon(i int) string {
	return [...]string{"paw", "eye", "camera"}[i%3]
}

// yearSpan reads like "Seen from 1987 to 2026."
func yearSpan(s sighting.Summary) string {
	if s.Earliest.IsZero() {
		return ""
	}
	return "Seen from " + s.Earliest.Format("2006") + " to " + s.Latest.Format("2006") + "."
}
