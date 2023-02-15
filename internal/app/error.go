package app

func (a *statsApp) Error() chan error {
	return a.errCh
}
