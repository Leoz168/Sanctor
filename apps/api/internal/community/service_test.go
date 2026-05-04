package community

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

var (
    commCreator = uuid.MustParse("aaaaaaaa-1111-2222-3333-aaaaaaaaaaaa")
    instID = uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
)

type commStub struct{
    created *Community
    addUserErr error
    addInstErr error
    deleteCalled bool
}
func (r *commStub) Create(c *Community) error { r.created = c; return nil }
func (r *commStub) FindByID(id uuid.UUID) (*Community, error) { if r.created!=nil && r.created.ID==id { return r.created, nil }; return nil, errors.New("not found") }
func (r *commStub) FindAll() []*Community { return []*Community{r.created} }
func (r *commStub) Update(c *Community) error { r.created = c; return nil }
func (r *commStub) Delete(id uuid.UUID) error { r.deleteCalled = true; return nil }
func (r *commStub) AddUserToCommunity(uc *UserCommunity) error { return r.addUserErr }
func (r *commStub) AddCommunityToInstitution(ci *CommunityInstitution) error { return r.addInstErr }
func (r *commStub) RemoveUserFromCommunity(userID, communityID uuid.UUID) error { return nil }
func (r *commStub) GetCommunityMembers(communityID uuid.UUID) ([]*UserCommunity, error) { return []*UserCommunity{{UserID: commCreator, Role: "owner"}}, nil }
func (r *commStub) GetUserCommunities(userID uuid.UUID) []*UserCommunity { return nil }
func (r *commStub) IsUserInCommunity(userID, communityID uuid.UUID) bool { return false }
func (r *commStub) GetMemberCount(communityID uuid.UUID) int { return 1 }
func (r *commStub) GetUserRole(userID, communityID uuid.UUID) (string, error) { if userID==commCreator { return "owner", nil }; return "member", nil }

func TestCreateCommunityRollbacksOnAddUserFail(t *testing.T){
    repo := &commStub{addUserErr: errors.New("fail add user")}
    svc := NewService(repo)
    req := CreateCommunityRequest{Name: "My Comm", CreatedBy: commCreator.String(), InstitutionID: instID.String()}
    _, err := svc.CreateCommunity(req)
    if err==nil || err.Error()!="failed to add creator to community" { t.Fatalf("expected failure on add user, got %v", err) }
}

func TestCreateCommunityRollbacksOnAddInstitutionFail(t *testing.T){
    repo := &commStub{addInstErr: errors.New("fail inst")}
    svc := NewService(repo)
    req := CreateCommunityRequest{Name: "My Comm", CreatedBy: commCreator.String(), InstitutionID: instID.String()}
    _, err := svc.CreateCommunity(req)
    if err==nil || err.Error()!="failed to link community to institution" { t.Fatalf("expected failure on link institution, got %v", err) }
}

func TestGetCommunityValidation(t *testing.T){
    svc := NewService(&commStub{})
    _, err := svc.GetCommunity(uuid.Nil)
    if err==nil || err.Error()!="community ID is required" { t.Fatalf("expected id required") }
}

func TestAddUserToCommunityValidations(t *testing.T){
    repo := &commStub{}
    svc := NewService(repo)
    err := svc.AddUserToCommunity(AddUserToCommunityRequest{UserID: "", CommunityID: ""})
    if err==nil { t.Fatalf("expected validation error") }
    // invalid role
    repo.created = &Community{ID: commCreator}
    err = svc.AddUserToCommunity(AddUserToCommunityRequest{UserID: commCreator.String(), CommunityID: commCreator.String(), Role: "bad"})
    if err==nil || err.Error()!="invalid role: must be member, admin, or owner" { t.Fatalf("expected invalid role") }
}
