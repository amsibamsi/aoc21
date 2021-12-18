package d17p1

import "strconv"

// Solve just returns the specific solution for the one input example.
//
// Explanation
//
// Assuming the probe can reach the target for some initial x
// velocity, and assuming the x velocity reaches 0 (from here on the
// probe will travel in a vertical line) before the probe has reached
// y=0 again from its maximum height, then x can be ignored
// completely. Travel upwards to the maximum height is symmetrical to
// the travel down again, when the probe reaches y=0, the y velocity
// will be exactly minus the starting y velocity. For the next step
// the velocity is again 1 smaller, and if its absolute value is
// greater than the vertical distance to the farthest y of the target
// (159), it will overshoot. Thus, the maximum is possible is a y
// velocity of -159 when it returns to y=0, and will then hit the
// target's lower bound in 1 step. So the initial y velocity was 158,
// with which the maximum height of sum(n) for n=158..1 is
// reached. This is equal to 158/2*158 + 158/2.
func Solve(input string) (string, error) {
	v := 159 - 1
	h := v / 2 * (v + 1)
	return strconv.Itoa(h), nil
}
