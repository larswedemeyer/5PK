// --> VERSION 1, RIGID RULES WILL BE USED

// VERSION 2, HYBRID-MODEL
// VERSION 3, REINFORCEMENT LEARNING
// THIS MODULE DETERMINES WHAT THE ROBOT SHOULD DO, BUT NOT HOW IT SHOULD EXECUTE THE ACTION

package verhaltensweise

import (
	wm "CODE/WORLDMODEL/V1"
)

// SEARCHING THE BALL
// GOING TO THE BALL
// PICKING UP THE BALL (Search is over)

// NEEDS A BALL EVERSEEN
// DISTANCE TOBALL

func Decisions(w *wm.World) {
	if !w.Ball.EverSeen {
		SearchBall()
	} else {
		abs, _ := w.DistanceToBall()
		if abs < 0.5 {
			PickBallUp()
		} else {
			GoToBall()
		}
	}
}

func PickBallUp() {}

func SearchBall() {}

func GoToBall() {

}

// REFERS TO THE RRT* PATH FOR MOVEMENT
