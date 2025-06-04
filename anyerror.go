package chanutils

// AnyError returns first non-nil error from provided channels (if any), or nil otherwise
func AnyError(cs ...<-chan error) error {
	if len(cs) == 0 {
		return nil
	}

	errsChan := Merge(cs...)
	for err := range errsChan {
		if err != nil {
			return err
		}
	}

	return nil
}
