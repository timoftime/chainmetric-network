package identity

import (
	"time"
)


type (
	// EnrollmentOption allows passing parameters for Enroll method.
	EnrollmentOption interface {
		Apply(*enrollArgs)
	}

	// EnrollOptionFunc is a function that mutates model during Enroll execution.
	EnrollOptionFunc func(*enrollArgs)

	enrollArgs struct {
		UserID   string
		Role     string
		ExpireAt *time.Time
	}
)


// Apply calls EnrollOptionFunc on model.
func (f EnrollOptionFunc) Apply(args *enrollArgs) {
	f(args)
}

// WithRole creates user with given `role`.
func WithRole(role string) EnrollmentOption {
	return EnrollOptionFunc(func(args *enrollArgs) {
		args.Role = role
	})
}

// WithExpiration creates user with given `expireAt`.
func WithExpiration(expireAt *time.Time) EnrollmentOption {
	return EnrollOptionFunc(func(args *enrollArgs) {
		args.ExpireAt = expireAt
	})
}
