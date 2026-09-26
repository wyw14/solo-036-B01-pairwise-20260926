package launcher

func CombineLaunchResults(valveErr, routeErr error) error {
	if valveErr != nil {
		return valveErr
	}
	if routeErr != nil {
		return routeErr
	}
	return nil
}
