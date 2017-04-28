package sliceint

type Sliceint []int

func NewSliceint(param ...int) Sliceint {
	slice := make(Sliceint, 0)
	if len(param) == 0 {
		return slice
	} else {
		for _, v := range param {
			slice = append(slice, v)
		}
		return slice
	}
}
func (this *Sliceint) Add(elem ...int) {
	*this = append(*this, elem...)
}
func (this *Sliceint) Delete(elem ...int) {
	if len(*this) == 0 {
		return
	}
	for _, v := range elem {
		this.delete(v)
	}
}
func (this *Sliceint) Check(elem int) bool {
	if len(*this) == 0 {
		return false
	}
	for _, v := range *this {
		if v == elem {
			return true
		}
	}
	return false
}
func (this *Sliceint) delete(elem int) {
	if len(*this) == 0 {
		return
	}
	for k, v := range *this {
		if v == elem {
			switch k {
			case 0:
				*this = (*this)[1:]
			case len(*this) - 1:
				*this = (*this)[:len(*this)-1]
			default:
				*this = append((*this)[:k], (*this)[k+1:]...)
			}
		}
	}
}
