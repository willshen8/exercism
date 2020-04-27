package erratum

// Use open a resource and handles errors if resources not successful
func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()
	if _, ok := err.(TransientError); ok {
		return Use(o, input)
	} else if err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			v, ok := r.(FrobError)
			if ok {
				resource.Defrob(v.defrobTag)
				err = v.inner
			} else {
				err = r.(error)
			}

		}
		resource.Close()
	}()
	resource.Frob(input)
	defer resource.Close()

	return
}
