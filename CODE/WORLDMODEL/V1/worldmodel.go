// --> VERSION 1, USES A 2-DIMENSIONAL MAP
// FINISHED but the Update function has to be updated and the other modules have to be imbetted

// VERSION 2, IMPLEMENTING A GRID-MODEL WITH SAVED HIGHTS OF OCCUPIED CELLS

// VERSION 3, A REAL 3-DIMENSIONAL MAP

package wordmodel

import (
	ty "CODE/PATHPLANNING/TYPES"
	"encoding/json"
	"math"
	"os"
	"time"
)

type Ball struct {
	Visible    bool
	Position   ty.PointRRT
	LastSeen   time.Time
	Confidence float64
	EverSeen   bool
}

func NewBall() Ball {
	return Ball{Visible: false, Position: ty.PointRRT{X: 0, Z: 0, Y: 0}, LastSeen: time.Now(), Confidence: 0, EverSeen: false}
}

func (b *Ball) Found(pos ty.PointRRT) {
	b.Visible = true
	b.Position = pos
	b.LastSeen = time.Now()
	b.Confidence = 1.0
	if !b.EverSeen {
		b.EverSeen = true
	}
}

func (b *Ball) Lost() {
	b.Visible = false
}

func (b *Ball) UpdateConfidence() {
	if b.Visible {
		b.Confidence = 1.0
		return
	}
	seconds := time.Since(b.LastSeen).Seconds()
	b.Confidence = 1.0 - seconds/10.0
	if b.Confidence < 0 {
		b.Confidence = 0
	}
}

func (w *World) Update() {
	w.Ball.UpdateConfidence()
}

type Robot struct {
	Position ty.PointRRT
	Heading  float64
	Pitch    float64
	Roll     float64
}

func NewRobot() Robot {
	return Robot{Position: ty.PointRRT{X: 0, Z: 0, Y: 0}}
}

func (r *Robot) UpdatePosition(pos ty.PointRRT) {
	r.Position = pos
}

func (r *Robot) UpdateHeading(h float64) {
	r.Heading = h
}

func (r *Robot) UpdatePitch(p float64) {
	r.Pitch = p
}

func (r *Robot) UpdateRoll(rl float64) {
	r.Roll = rl
}

type Obstacle struct {
	ID       string
	Position ty.PointRRT
	Width    float64
	Height   float64
}

func NewObstacle(id string, x, y, z, width, height float64) *Obstacle {
	return &Obstacle{ID: id, Position: ty.PointRRT{X: 0, Z: 0, Y: 0}, Width: width, Height: height}
}

type World struct {
	Ball      Ball
	Robot     Robot
	Obstacles []*Obstacle
}

func NewWorld() *World {
	return &World{Ball: NewBall(), Robot: NewRobot(), Obstacles: []*Obstacle{}}
}

func (w *World) AddObstacle(o *Obstacle) {
	w.Obstacles = append(w.Obstacles, o)
}

// calculating the distance using pythagoras
func (w *World) DistanceToBall() (float64, bool) {
	if !w.Ball.EverSeen { // if the ball hadnt been found yet we have to take that into account
		return 0, false
	}
	return math.Hypot(w.Ball.Position.X-w.Robot.Position.X, w.Ball.Position.Y-w.Robot.Position.Y), true
}

func (w *World) HeadingToBall() (float64, bool) {
	if !w.Ball.EverSeen {
		return 0, false
	}
	dx := w.Ball.Position.X - w.Robot.Position.X
	dy := w.Ball.Position.Y - w.Robot.Position.Y
	return math.Atan2(dy, dx), true
	// Atan2 returns the angle to the x-axis, atan2 uses the Relativvektor to the ball -> the heading is always right
}

func (w *World) TurnAngleToBall() (float64, bool) {

	heading, _ := w.HeadingToBall()

	if !w.Ball.EverSeen {
		return 0, false
	}

	turn := heading - w.Robot.Heading

	for turn > math.Pi {
		turn -= 2 * math.Pi
	}

	for turn < -math.Pi {
		turn += 2 * math.Pi
	}

	// <- makes sure that the value is between -Pi and Pi

	return turn, true
}

// WRITING IN THE JSON
// THE JSON FILE IS TEMPORARY
// IT IS CURRENTLY USED TO PROVIDE AN EASY WAY TO READ THE DATA FOR CONTROL

type WorldDataJSON struct {
}

type DataJSON struct {
	Running bool
	Visible bool
	Counter int
	Score   int
}

func SaveData(x DataJSON) {

	file, _ := os.Create("world.json")
	defer file.Close()

	json.NewEncoder(file).Encode(x)
}

// CHECKS WHETHER AN OBJECT IS IN THE WAY OR WHETHER THE ROBOT IS TOO LARGE TO PASS
func ObstacleFree(a, b ty.PointRRT) bool {
	return true
}
