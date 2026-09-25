package smoothing

import (
	r "CODE/PATHPLANNING/RRT_STAR"
	ty "CODE/PATHPLANNING/TYPES"
	wm "CODE/WORLDMODEL/V1"
)

// path pruning
// path smoothing per Bézierkurve
// dann Projektion auf eine gewölbte Fläche

func PathSmoothing(goal *ty.RRTNode) []ty.PointRRT {
	return hilfpathsmoothing(r.ExtractPath(goal))
}

func hilfpathsmoothing(startpath []ty.PointRRT) []ty.PointRRT {
	path := pruning(startpath)
	return path
}

func pruning(p []ty.PointRRT) []ty.PointRRT { // filter out unnecessary points

	if len(p) == 0 { // error handling
		return nil
	}

	pruned := []ty.PointRRT{p[0]}

	current := 0

	for current < len(p)-1 {

		farthest := current + 1 // farthest point to the starting point where a connection exists

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
