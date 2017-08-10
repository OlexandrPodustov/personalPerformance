package erratum

import (
	"errors"
	"fmt"
)

const testVersion = 2

func Use(f func() (Resource, error), s string) (rr error) {
	var r Resource
	var e error
	var ok bool
	defer func() {
		eRec := recover()
		if eRec != nil {
			if erVal, ok := eRec.(FrobError); ok {
				rr = erVal
				r.Defrob(erVal.defrobTag)
				r.Close()
			}
			rr = eRec.(error)
			r.Close()
		}
	}()

	r, e = f()

	if e != nil {

		if _, ok = e.(TransientError); ok {
			var ee error
			var r1 Resource

			for i := 0; i < 3; i++ {
				r1, ee = f()
			}
			if ee != nil {
				//log.Println("1", e)
				return errors.New(fmt.Sprint(ee))
			}
			r1.Frob(s)
		}

		return nil
	}

	r.Frob(s)
	r.Close()
	return nil
}
