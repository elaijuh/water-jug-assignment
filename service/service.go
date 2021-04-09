package service

import (
	"fmt"
)

type ProblemResolver struct {
	x int
	y int
	z int
}

type Actions struct {
	steps   int
	actions []string
}

func NewProblemResolver(x, y, z int) *ProblemResolver {
	return &ProblemResolver{x, y, z}
}

func (r *ProblemResolver) Resolvable() bool {
	x := r.x
	y := r.y
	if y > x {
		x, y = y, x
	}
	if r.z > x {
		return false
	}
	if r.z%r.gcd(x, y) != 0 {
		return false
	}
	return true
}

func (r *ProblemResolver) GetActions() []string {
	actx := r.act('x')
	acty := r.act('y')
	if actx.steps > acty.steps {
		return acty.actions
	}
	return actx.actions
}

// Transfer water start from 'from'(either x or y), store the steps and actions
func (r *ProblemResolver) act(from byte) Actions {
	a := r.x
	b := 0
	to := 'y'
	acap := r.x
	bcap := r.y

	if from == 'y' {
		a = r.y
		to = 'x'
		acap = r.y
		bcap = r.x
	}

	ans := Actions{steps: 0, actions: make([]string, 0)}
	ans.steps++
	ans.actions = append(ans.actions, fmt.Sprintf("Fill bucket %c -> %c:%d,%c:%d", from, from, a, to, b))

	for a != r.z && b != r.z {
		if a < bcap-b {
			ans.steps++
			b += a
			a = 0
			ans.actions = append(ans.actions, fmt.Sprintf("Transfer bucket %c to bucket %c -> %c:%d,%c:%d", from, to, from, a, to, b))
			if b == r.z {
				break
			}
			ans.steps++
			a = acap
			ans.actions = append(ans.actions, fmt.Sprintf("Fill bucket %c -> %c:%d,%c:%d", from, from, a, to, b))
		} else {
			ans.steps++
			a -= bcap - b
			b = bcap
			ans.actions = append(ans.actions, fmt.Sprintf("Transfer bucket %c to bucket %c -> %c:%d,%c:%d", from, to, from, a, to, b))
			if a == r.z || b == r.z {
				break
			}
			ans.steps++
			b = 0
			ans.actions = append(ans.actions, fmt.Sprintf("Empty bucket %c -> %c:%d,%c:%d", to, from, a, to, b))
		}
	}

	return ans
}

func (r *ProblemResolver) gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return r.gcd(b, a%b)
}
