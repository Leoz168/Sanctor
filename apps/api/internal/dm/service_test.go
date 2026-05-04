package dm

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

var (
    dmUserA = uuid.MustParse("10000000-0000-0000-0000-000000000000")
    dmUserB = uuid.MustParse("20000000-0000-0000-0000-000000000000")
)

type dmStub struct{
    group *DMGroup
    err error
    members map[uuid.UUID]bool
    messages []*DMMessage
}
func (r *dmStub) CreateGroup(g *DMGroup) error { r.group = g; return r.err }
func (r *dmStub) AddUserToGroup(gu *DMGroupUser) error { return r.err }
func (r *dmStub) GetGroupUsers(groupID uuid.UUID) ([]*DMGroupUser, error) { return nil, nil }
func (r *dmStub) GetUserGroups(userID uuid.UUID) ([]*DMGroup, error) { return nil, nil }
func (r *dmStub) FindDirectGroupByUsers(userA, userB uuid.UUID) (*DMGroup, error) { return nil, ErrDMGroupNotFound }
func (r *dmStub) IsUserInGroup(userID, groupID uuid.UUID) bool { return r.members[userID] }
func (r *dmStub) SaveMessage(m *DMMessage) error { r.messages = append(r.messages, m); return r.err }
func (r *dmStub) GetGroupMessages(groupID uuid.UUID, limit int) ([]*DMMessage, error) { return r.messages, nil }

func TestCreateOrGetDirectGroupValidations(t *testing.T){
    svc := NewService(&dmStub{})
    _, err := svc.CreateOrGetDirectGroup(dmUserA, dmUserA)
    if err==nil || err.Error()!="direct message requires two different users" { t.Fatalf("expected different users error") }
    _, err = svc.CreateOrGetDirectGroup(uuid.Nil, dmUserB)
    if err==nil { t.Fatalf("expected id required") }
}

func TestSendMessageValidationsAndSuccess(t *testing.T){
    repo := &dmStub{members: map[uuid.UUID]bool{dmUserA:true}}
    svc := NewService(repo)
    _, err := svc.SendMessage(SendMessageRequest{GroupID: "", UserID: dmUserA.String(), Content: "hi"})
    if err==nil || err.Error()!="groupId and userId are required" { t.Fatalf("expected required") }

    // invalid UUIDs
    _, err = svc.SendMessage(SendMessageRequest{GroupID: "bad", UserID: "bad", Content: "hi"})
    if err==nil { t.Fatalf("expected invalid format") }

    // not a member
    repo.members = map[uuid.UUID]bool{}
    _, err = svc.SendMessage(SendMessageRequest{GroupID: uuid.New().String(), UserID: dmUserA.String(), Content: "hi"})
    if err==nil || !errors.Is(err, ErrNotDMMember) { t.Fatalf("expected not member") }

    // success path
    gid := uuid.New()
    repo.members = map[uuid.UUID]bool{dmUserA:true}
    repo.messages = []*DMMessage{}
    // stub will accept SaveMessage
    msg, err := svc.SendMessage(SendMessageRequest{GroupID: gid.String(), UserID: dmUserA.String(), Content: "hello"})
    if err!=nil { t.Fatalf("expected success got %v", err) }
    if msg.ID==uuid.Nil { t.Fatalf("expected id set") }
    if msg.CreatedAt.IsZero() { t.Fatalf("expected timestamps set") }
}
