package status_checker

type StatusChecker interface {
	CheckStatus(url string) (status bool, err error)
}
