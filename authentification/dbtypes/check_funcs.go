package dbtypes

func (reqAccess *AccessServices) In(targetAccess *AccessServices) bool {
	for _, reqAgent := range reqAccess.Task.Agents {
		in := false
		for _, targetAgent := range targetAccess.Task.Agents {
			if reqAgent == targetAgent {
				in = true
				break
			}
		}
		if !in {
			return false
		}
	}
	return true
}
