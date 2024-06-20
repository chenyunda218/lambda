package lambda

func Map[T, O any](arr []T, cb func(T) O) []O {
	var output []O = make([]O, len(arr))
	for i := 0; i < len(arr); i++ {
		output[i] = cb(arr[i])
	}
	return output
}

func MapI[T, O any](arr []T, cb func(int, T) O) []O {
	var output []O = make([]O, len(arr))
	for i := 0; i < len(arr); i++ {
		output[i] = cb(i, arr[i])
	}
	return output
}

func Filter[T any](arr []T, cb func(T) bool) []T {
	var output []T
	for i := 0; i < len(arr); i++ {
		if cb(arr[i]) {
			output = append(output, arr[i])
		}
	}
	return output
}

func FilterI[T any](arr []T, cb func(int, T) bool) []T {
	var output []T
	for i := 0; i < len(arr); i++ {
		if cb(i, arr[i]) {
			output = append(output, arr[i])
		}
	}
	return output
}

func Reduce[T, O any](arr []T, init O, cb func(O, T) O) O {
	for i := 0; i < len(arr); i++ {
		init = cb(init, arr[i])
	}
	return init
}

func ReduceI[T, O any](arr []T, init O, cb func(int, O, T) O) O {
	for i := 0; i < len(arr); i++ {
		init = cb(i, init, arr[i])
	}
	return init
}

func Head[T any](arr []T) T {
	return arr[0]
}

func Tail[T any](arr []T) []T {
	return arr[1:]
}

func HeadTail[T any](arr []T) (T, []T) {
	return arr[0], arr[1:]
}

func Reverse[T any](arr []T) []T {
	var output []T = make([]T, len(arr))
	for i := 0; i < len(arr); i++ {
		output[i] = arr[len(arr)-1-i]
	}
	return output
}

func ForEach[T any](arr []T, cb func(T)) {
	for i := 0; i < len(arr); i++ {
		cb(arr[i])
	}
}
