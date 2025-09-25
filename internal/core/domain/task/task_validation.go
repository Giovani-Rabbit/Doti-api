package taskdomain

type TaskPositionParams struct {
	Id       int32 `json:"id"`
	Position int32 `json:"position"`
}

func HasRepeatedPositions(tasks []TaskPositionParams) bool {
	seen := make(map[int32]bool)

	for _, p := range tasks {
		if seen[p.Position] {
			return true // has repeated number
		}
		seen[p.Position] = true // position seen
	}

	return false
}
