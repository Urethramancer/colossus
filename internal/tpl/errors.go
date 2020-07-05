package tpl

type errNoName struct{}

func (e errNoName) Error() string {
	return "no template name specified"
}
