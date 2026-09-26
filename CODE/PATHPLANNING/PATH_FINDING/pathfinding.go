package pathfinding

import (
	s "CODE/PATHPLANNING/PATH_SMOOTHING"
	r "CODE/PATHPLANNING/RRT_STAR"
	ty "CODE/PATHPLANNING/TYPES"
)

// the robot will give its current position and the position where he wants to be
// it will get a path to a point which was the closest possible to the goal
func Finding(currobot, goal ty.PointRRT, N int, rad, deltaQ, goalRadius, min, max, t float64) []ty.PointRRT {
	_, bestgoal := r.RRTstar(currobot, goal, N, rad, deltaQ, goalRadius, min, max)
	return s.PathSmoothing(bestgoal, N, rad, deltaQ, goalRadius, min, max, t)
}
