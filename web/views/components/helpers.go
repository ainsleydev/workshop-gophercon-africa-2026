package components

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"workshop/internal/domain/sighting"
	"workshop/internal/domain/species"
)

// Name returns the species name a reader would recognise, capitalised.
func Name(sp species.Species) string {
	name := cmp.Or(sp.VernacularName, sp.CanonicalName, sp.ScientificName, "Unknown species")
	r, size := utf8.DecodeRuneInString(name)
	return string(unicode.ToUpper(r)) + name[size:]
}

// Where joins the locality and county, leaving the country out as every sighting is in Kenya.
func Where(l sighting.Location) string {
	parts := slices.DeleteFunc(
		[]string{l.Locality, l.StateProvince},
		func(s string) bool { return s == "" },
	)
	return cmp.Or(strings.Join(parts, ", "), l.Country, "Somewhere in Kenya")
}

// Photo returns the sighting's first still image, or the zero Media if it has none.
func Photo(s sighting.Sighting) sighting.Media {
	for _, m := range s.Media {
		if m.Type == "StillImage" && m.URL != "" {
			return m
		}
	}
	return sighting.Media{}
}

// RecordURL links to the original record, falling back to its GBIF page.
func RecordURL(s sighting.Sighting) string {
	if strings.HasPrefix(s.ReferenceURL, "https://") || strings.HasPrefix(s.ReferenceURL, "http://") {
		return s.ReferenceURL
	}
	return "https://www.gbif.org/occurrence/" + strconv.FormatInt(s.GBIFKey, 10)
}

// Thousands formats n with commas, such as 1,498.
func Thousands(n int) string {
	s := strconv.Itoa(n)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}

// individuals renders how many animals were seen in a single sighting.
func individuals(n int) string {
	if n == 1 {
		return "1 animal"
	}
	return fmt.Sprintf("%d animals", n)
}

// Sized swaps an iNaturalist photo's original for a smaller copy, such as "medium" (500px) or "large" (1024px).
func Sized(url, size string) string {
	if !strings.Contains(url, "inaturalist") {
		return url
	}
	return strings.Replace(url, "/original.", "/"+size+".", 1)
}
