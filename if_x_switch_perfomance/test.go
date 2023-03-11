package test

//If ..
func If() {
	var test bool = true

	if test {
		println("test")
	}

}

//Swich ..
func Swich() {
	var test bool = true

	switch test {
	case true:
		println("test")
	}

}

//Swich1 ..
func Swich1() {
	var test bool = true

	switch {
	case test == true:
		println("test")
	}

}
