package trydo

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

func (it Block) Do() {
	if it.Finally != nil {
		defer it.Finally()
	}

	if it.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				it.Catch(r)
			}
		}()
	}

	it.Try()
}
