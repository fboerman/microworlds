package sdl2canvas

import "math"

func EvenqToCube(q float64, r float64) (x float64, y float64, z float64) {
	x = q
	z = r - (q + math.Mod(math.Abs(q), 2.)) / 2
	y = -x-z

	return
}

func CubeToEvenq(x float64, y float64, z float64) (q float64, r float64) {
	q = x
	r = z + (x + math.Mod(math.Abs(x), 2.)) / 2

	return
}

func RoundCube(x_f float64, y_f float64, z_f float64) (x float64, y float64, z float64) {
	x = math.Round(x_f)
	y = math.Round(y_f)
	z = math.Round(z_f)

	x_diff := math.Abs(float64(x) - x_f)
	y_diff := math.Abs(float64(y) - y_f)
	z_diff := math.Abs(float64(z) - z_f)

	if x_diff > y_diff && x_diff > z_diff {
		x = -y-z
	} else if y_diff > z_diff {
		y = -x-z
	} else {
		z = -x-y
	}

	return
}


func PixelToHex(px int, py int, size int) (q int, r int) {
	size_f := float64(size)
	px_f := float64(px)
	py_f := float64(py)

	// calculate fractional hex coordinates
	q_f := (2./3. * px_f) / size_f
	r_f := (-1./3. * px_f  +  math.Sqrt(3)/3. * py_f) / size_f
	// round them into the correct hex
	// go to cube domain
	x_f, y_f, z_f := EvenqToCube(q_f, r_f)
	// round the hex in cube domain
	x_f, y_f, z_f = RoundCube(x_f, y_f, z_f)
	// move back to offset domain
	q_f, r_f = CubeToEvenq(x_f, y_f, z_f)
	// lastly move to integers
	q = int(q_f)
	r = int(r_f)

	return
}
