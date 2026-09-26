package smoothing

import (
	r "CODE/PATHPLANNING/RRT_STAR"
	ty "CODE/PATHPLANNING/TYPES"
	wm "CODE/WORLDMODEL/V1"
)

// path pruning
// path smoothing per B Spline curve
// dann Projektion auf eine gewölbte Fläche

// zu steile kurven noch beachten

func PathSmoothing(goal *ty.RRTNode, N int, rad, deltaQ, goalRadius, min, max, t float64) []ty.PointRRT {
	return hilfpathsmoothing(r.ExtractPath(goal), N, rad, deltaQ, goalRadius, min, max, t)
}

func hilfpathsmoothing(startpath []ty.PointRRT, N int, rad, deltaQ, goalRadius, min, max, t float64) []ty.PointRRT {
	purgedpath := pruning(startpath)
	smoothedpath := bspline(purgedpath, t)
	projectedpath := projection(smoothedpath)
	return correct(projectedpath, N, rad, deltaQ, goalRadius, min, max)
}

// can be replaced with random purging
func pruning(p []ty.PointRRT) []ty.PointRRT { // filter out unnecessary points

	if len(p) == 0 { // error handling
		return nil
	}
	pruned := []ty.PointRRT{p[0]}
	current := 0
	for current < len(p)-1 {
		farthest := current + 1                 // farthest point to the starting point where a connection exists
		for i := len(p) - 1; i > current; i-- { // counting downwards from the farthest
			if wm.ObstacleFree(p[current], p[i]) {
				farthest = i
				break
			}
		}
		pruned = append(pruned, p[farthest]) // append it to the new paths
		current = farthest
	}

	return pruned
}

// uses the B Spline to smooth the path and projects it onto the surface

// c(t) = sum n i=0 Bik(t) * pi; t e [0,1]

// Der Knotenvektor wird T bezeichnet, Knoten aus dem Knotenvektor werden ab jetzt mit u bezeichnet, ansonsten verliere ich den Verstand

// gibt eine Liste an Werten aus
func bspline(p []ty.PointRRT, t float64) []ty.PointRRT {

	if len(p) < 4 {
		return p
	}

	bspp := []ty.PointRRT{}

	for anf := 0.0; anf < 1; anf += t {
		bspp = append(bspp, bsplinehilf(p, anf))
	}
	bspp = append(bspp, bsplinehilf(p, 1))

	bspp[0] = p[0]
	bspp[len(bspp)-1] = p[len(p)-1] // the start and goal have to always be the same

	return bspp
}

// gibt einen Wert aus
func bsplinehilf(p []ty.PointRRT, t float64) ty.PointRRT {

	var k int = 4
	n := len(p) - 1

	T := knotenvektoren(n, k)

	s := ty.PointRRT{X: 0, Y: 0, Z: 0}

	for i := 0; i <= n; i++ { // sum
		cur := ty.VectorScaling(b(i, k, t, T), p[i])
		s = ty.AddTwoPoints(s, cur)
	}
	return s
}

func b(i, k int, t float64, T []float64) float64 {

	if k == 1 {
		if t >= T[i] && t < T[i+1] {
			return 1
		}
		if t == 1 && T[i+1] == 1 {
			return 1
		} else {
			return 0
		}
	}
	result := 0.0

	if (T[i+k-1] - T[i]) != 0 {
		result += (t - T[i]) / (T[i+k-1] - T[i]) * b(i, k-1, t, T)
	}
	if T[i+k]-T[i+1] != 0 {
		result += (T[i+k] - t) / (T[i+k] - T[i+1]) * b(i+1, k-1, t, T)
	}

	return result
}

func knotenvektoren(n, k int) []float64 {

	T := []float64{}

	for i := 0; i <= n+k; i++ {
		if i < k {
			T = append(T, 0)
		} else if i > n {
			T = append(T, 1)
		} else {
			T = append(T, float64(i-k+1)/float64(n-k+2))
		}
	}

	return T
}

// graph needs to be projected on the real 3D map in order to calculate possible movements
func projection(t []ty.PointRRT) []ty.PointRRT {
	return t
}

// it has to be made sure that the final path is also valid
func correct(p []ty.PointRRT, N int, rad, deltaQ, goalRadius, min, max float64) []ty.PointRRT {

	final := []ty.PointRRT{}

	for i := 0; i < len(p)-1; i++ {
		if wm.ObstacleFree(p[i], p[i+1]) { // if the path is valid, append the point
			final = append(final, p[i])
		} else {
			_, newpoint := r.RRTstar(p[i], p[i+1], N, rad, deltaQ, goalRadius, min, max)
			if newpoint == nil {
				continue
			}
			newpath := r.ExtractPath(newpoint)
			final = append(final, newpath[1:]...) // schauen ob stimmt
		}
	}

	if len(p) > 0 {
		final = append(final, p[len(p)-1])
	}

	return final
}
