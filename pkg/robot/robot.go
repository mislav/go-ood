package robot

import (
	"github.com/ronnas/go-ood/pkg/maze"
	"github.com/ronnas/go-ood/pkg/maze/travel"
)

// Step represents a movement in a maze, it has a target cell and direction
type Step struct {
	C maze.Cell
	D maze.Direction
}

// Robot records its steps as it roams though a maze
type Robot struct {
	travel.Travel
	steps []Step
}

// New returns a new gopher
func New(t travel.Travel) Robot {
	r := Robot{Travel: t, steps: []Step{{t.Maze.CellFromCoords(t.Coords), t.Dir}}}
	return r
}

// Steps returns a copy of the gopher's steps
func (r Robot) Steps() []Step {
	steps := make([]Step, len(r.steps))
	copy(steps, r.steps)
	return steps
}

// Finished returns true or false if the gopher is done going through the maze
func (r Robot) Finished() bool {
	return r.Coords[maze.X] == r.Maze.DimX-1 && r.Coords[maze.Y] == r.Maze.DimY-1
}

// Move moves the gopher in its current direction. Returns an error if impossible to move.
func (r *Robot) Move() error {
	if err := r.Travel.Move(); err != nil {
		return err
	}
	r.steps = append(r.steps, Step{C: r.Maze.CellFromCoords(r.Coords), D: r.Dir})
	return nil
}

// TurnLeft turns the Robot left
func (r *Robot) TurnLeft() {
	r.Travel.TurnRight()
	r.Travel.TurnRight()
	r.Travel.TurnRight()
	r.steps = append(r.steps, Step{C: r.Maze.CellFromCoords(r.Coords), D: r.Dir})
}

// TurnRight turns the Robot right
func (r *Robot) TurnRight() {
	r.Travel.TurnRight()
	r.steps = append(r.steps, Step{C: r.Maze.CellFromCoords(r.Coords), D: r.Dir})
}