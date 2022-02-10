package helper

type Error struct {
	Reason      string
	Instruction string
}

func (e Error) Error() string {
	return "reason: " + e.Reason
}
