package bitbucketmock

import (
	"context"
	"github.com/Interhyp/metadata-service/acorns/repository"
	auacornapi "github.com/StephanHCB/go-autumn-acorn-registry/api"
	"github.com/pkg/errors"
)

const FILTER_FAILED_USERNAME = "filterfailedusername"

type BitbucketMock struct {
}

func New() auacornapi.Acorn {
	return &BitbucketMock{}
}

// implement acorn

func (b *BitbucketMock) AcornName() string {
	return repository.BitbucketAcornName
}

func (b *BitbucketMock) AssembleAcorn(registry auacornapi.AcornRegistry) error {
	return nil
}

func (b *BitbucketMock) SetupAcorn(registry auacornapi.AcornRegistry) error {
	return nil
}

func (b *BitbucketMock) TeardownAcorn(registry auacornapi.AcornRegistry) error {
	return nil
}

// implement bitbucket interface

func (b *BitbucketMock) IsBitbucket() bool {
	return true
}

func (b *BitbucketMock) Setup(ctx context.Context) error {
	return nil
}

func (b *BitbucketMock) GetBitbucketUser(ctx context.Context, username string) (repository.BitbucketUser, error) {
	return repository.BitbucketUser{
		Name: username,
	}, nil
}

func (b *BitbucketMock) GetBitbucketUsers(ctx context.Context, usernames []string) ([]repository.BitbucketUser, error) {
	result := []repository.BitbucketUser{}
	for _, username := range usernames {
		result = append(result, repository.BitbucketUser{
			Name: username,
		})
	}
	return result, nil
	//return []repository.BitbucketUser{
	//	{
	//		Name: "reviewer-one",
	//	},
	//	{
	//		Name: "reviewer-two",
	//	},
	//}, nil
}

func (b *BitbucketMock) FilterExistingUsernames(ctx context.Context, usernames []string) ([]string, error) {
	if usernames[0] == FILTER_FAILED_USERNAME {
		return []string{}, errors.New("error")
	}
	return usernames, nil
	// return []string{"approver-one", "approver-two"}, nil
}
