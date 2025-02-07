package support

// StringRef in LLVM is a borrowed immutable reference to an underlying string
// owned by some other code.
type StringRef struct {
	data string
}

func NewStringRef2(data string) StringRef {
	return StringRef{data: data}
}

func (s StringRef) String() string {
	return s.data
}