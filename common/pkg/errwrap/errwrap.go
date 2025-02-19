package errwrap

import "fmt"

func Wrap(msg string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", msg, err)
}

func Wrapf(format string, err error, a ...any) error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, a...)
	return fmt.Errorf("%s: %w", msg, err)
}
