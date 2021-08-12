package identity

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/pkg/errors"
	"github.com/timoth-y/chainmetric-contracts/shared/core"
	"github.com/timoth-y/chainmetric-contracts/shared/infrastructure/repository"
	"github.com/timoth-y/chainmetric-contracts/shared/model/user"
)

// Register performs users initial registration.
func Register(request user.RegistrationRequest) (*user.User, error) {
	var (
		user = &user.User{
			ID:        uuid.NewString(),
			Firstname: request.Firstname,
			Lastname:  request.Lastname,
			Email:     request.Email,
		}
		err error
	)

	if user.EnrollmentSecret, err = client.Register(&msp.RegistrationRequest{
		Name: user.IdentityName(),
		Type: "client",
	}); err != nil {
		return nil, errors.Wrap(err, "failed to register user")
	}

	if err := repository.NewUserMongo(core.MongoDB).Upsert(*user); err != nil {
		return nil, errors.Wrap(err, "failed to store user")
	}

	return user, nil
}

// Enroll generates msp.SigningIdentity for user and confirms one.
func Enroll(req user.EnrollmentRequest) error {
	var (
		repo = repository.NewUserMongo(core.MongoDB)
	)

	user, err := repo.GetByID(req.UserID)
	if err != nil {
		return errors.Wrap(err, "failed to found user registration")
	}

	if err = client.Enroll(user.IdentityName(), msp.WithSecret(user.EnrollmentSecret)); err != nil {
		return errors.Wrap(err, "failed to enroll user")
	}

	si, err := client.GetSigningIdentity(user.IdentityName())
	if err != nil {
		return errors.Wrap(err, "failed to get signing identity for new user")
	}

	user.Confirmed = true
	user.Role = req.Role
	user.ExpireAt = req.ExpireAt

	pk, _ := si.PrivateKey().Bytes()
	cert := si.PublicVersion().EnrollmentCertificate()

	fmt.Println(string(cert))
	fmt.Println(string(pk))

	if err = repo.Upsert(*user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
