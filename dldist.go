package main

// Compute the Damerau–Levenshtein distance between a and b.
// Reference: https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance#Distance_with_adjacent_transpositions
func DLDist(a, b string) int {
	ra := make([]rune, 0, len(a))
	rb := make([]rune, 0, len(b))
	for _, r := range a {
		ra = append(ra, r)
	}
	for _, r := range b {
		rb = append(rb, r)
	}
	var i, j int
	// Pre-process to remove the common prefix and suffix.
	for i = 0; i < len(ra) && i < len(rb) && ra[i] == rb[i]; i++ {
	}
	if i > 0 {
		ra = ra[i:]
		rb = rb[i:]
	}
	for i, j = len(ra)-1, len(rb)-1; i >= 0 && j >= 0 && ra[i] == rb[j]; i, j = i-1, j-1 {
	}
	if i < len(ra)-1 {
		ra = ra[:i+1]
		rb = rb[:j+1]
	}
	na, nb := len(ra)+2, len(rb)+2

	da := make(map[rune]int)
	d := make([][]int, na)
	for i = range d {
		d[i] = make([]int, nb)
	}

	maxDist := len(ra) + len(rb)
	d[0][0] = maxDist
	for i = 1; i < na; i++ {
		d[i][0] = maxDist
		d[i][1] = i - 1
	}
	for j = 1; j < nb; j++ {
		d[0][j] = maxDist
		d[1][j] = j - 1
	}

	var db, cost, k, l int
	for i = 2; i < na; i++ {
		db = 0
		for j = 2; j < nb; j++ {
			k = da[rb[j-2]]
			l = db
			if ra[i-2] == rb[j-2] {
				cost = 0
				db = j - 1
			} else {
				cost = 1
			}
			d[i][j] = min4int(
				d[i-1][j-1]+cost,          // substitution
				d[i][j-1]+1,               // insertion
				d[i-1][j]+1,               // deletion
				d[k][l]+(i-k-2)+1+(j-l-2), // transposition
			)
		}
		da[ra[i-2]] = i - 1
	}

	return d[na-1][nb-1]
}

func min4int(a, b, c, d int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	if a > d {
		a = d
	}
	return a
}
