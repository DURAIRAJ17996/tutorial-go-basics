package mock

func IsEven(num int) (bool, error) {
	if num <= 0 {
		return false, NewCustomError("Can't determince for 0", 400)
	}

	if num%2 == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
