package hasher

type Hasher interface {
	Compare(pw, hashPw string) bool
	Hash(pw string) (string, error)
}
