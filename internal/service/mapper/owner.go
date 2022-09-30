package mapper

import (
	"context"
	openapi "github.com/Interhyp/metadata-service/api/v1"
	"sort"
)

func (s *Impl) GetSortedOwnerAliases(_ context.Context) ([]string, error) {
	fileInfos, err := s.Metadata.ReadDir("owners/")
	if err != nil {
		return []string{}, err
	}

	result := make([]string, 0)
	for i := range fileInfos {
		alias := fileInfos[i].Name()
		if fileInfos[i].IsDir() {
			// check presence of owner.info.yaml to be sure
			_, err := s.Metadata.Stat("owners/" + alias + "/owner.info.yaml")
			if err == nil {
				if s.OwnerRegex.MatchString(alias) {
					result = append(result, alias)
				}
			}
		}
	}

	sort.Strings(result)
	return result, nil
}

func (s *Impl) GetOwner(ctx context.Context, ownerAlias string) (openapi.OwnerDto, error) {
	result := openapi.OwnerDto{}

	fullPath := "owners/" + ownerAlias + "/owner.info.yaml"
	err := GetT[openapi.OwnerDto](ctx, s, &result, fullPath)

	return result, err
}

func (s *Impl) WriteOwner(ctx context.Context, ownerAlias string, owner openapi.OwnerDto) (openapi.OwnerDto, error) {
	err := s.Metadata.Pull(ctx)
	if err != nil {
		return owner, err
	}

	path := "owners/" + ownerAlias
	fileName := "owner.info.yaml"
	description := "owner " + ownerAlias
	err = WriteT[openapi.OwnerDto](ctx, s, &owner, path, fileName, description, owner.JiraIssue)

	return owner, err
}

func (s *Impl) DeleteOwner(ctx context.Context, ownerAlias string, jiraIssue string) (openapi.OwnerPatchDto, error) {
	result := openapi.OwnerPatchDto{}

	err := s.Metadata.Pull(ctx)
	if err != nil {
		return result, err
	}

	fullPath := "owners/" + ownerAlias + "/owner.info.yaml"
	description := "owner " + ownerAlias
	err = DeleteT[openapi.OwnerPatchDto](ctx, s, &result, fullPath, description, jiraIssue)

	return result, err
}

func (s *Impl) IsOwnerEmpty(_ context.Context, ownerAlias string) bool {
	s.muOwnerCaches.Lock()
	defer s.muOwnerCaches.Unlock()

	for _, owner := range s.serviceOwnerCache {
		if owner == ownerAlias {
			return false
		}
	}

	for _, owner := range s.repositoryOwnerCache {
		if owner == ownerAlias {
			return false
		}
	}

	return true
}
