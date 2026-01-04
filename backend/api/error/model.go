package error

import "fmt"

type ValidationError struct {
	Title string
	Detail string 
}

type NotFoundError struct {
	Instance string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s.", v.Title, v.Detail)
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found.", n.Instance)
}