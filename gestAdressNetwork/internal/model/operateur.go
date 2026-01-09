package model

type Operateur struct {
	name string
	x    float64
	y    float64
	_2G  int
	_3G  int
	_4G  int
}


func  New(name string, x float64, y float64, deuxG int, troisG int, quatreG int) (*Operateur, error) {

	o := &Operateur{
		name: name,
		x:    x,
		y:    y,
		_2G:  deuxG,
		_3G:  troisG,
		_4G:  quatreG,
	}
	return o, nil
}
