package models

//WriterInfo is how WriterInfo or Claims structured. It is a helper struct
type WriterInfo struct {
	ID        int64
	AvatarURL string
	Name      string
	Bio       string
}
