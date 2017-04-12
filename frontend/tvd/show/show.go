package show

import (
	dlpb "github.com/diektronics/dl-us/protos/dl"
)

type Show struct {
	Name string
	Eps  string
	Blob string
	Down *dlpb.Down
}
