
type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *Reader) Read(p []byte) (n int, err error) {

}

func NewReader(s string) *Reader { return &Reader{s, 0, -1} }


