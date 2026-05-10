package dm

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	dmUserA = uuid.MustParse("10000000-0000-0000-0000-000000000000")
	dmUserB = uuid.MustParse("20000000-0000-0000-0000-000000000000")
)

type dmStub struct {
	group        *DMGroup
	err          error
	members      map[uuid.UUID]bool
	messages     []*DMMessage
	foundMessage *DMMessage
}

func (r *dmStub) CreateGroup(g *DMGroup) error                            { r.group = g; return r.err }
func (r *dmStub) AddUserToGroup(gu *DMGroupUser) error                    { return r.err }
func (r *dmStub) GetGroupUsers(groupID uuid.UUID) ([]*DMGroupUser, error) { return nil, nil }
func (r *dmStub) GetUserGroups(userID uuid.UUID) ([]*DMGroup, error)      { return nil, nil }
func (r *dmStub) FindDirectGroupByUsers(userA, userB uuid.UUID) (*DMGroup, error) {
	return nil, ErrDMGroupNotFound
}
func (r *dmStub) IsUserInGroup(userID, groupID uuid.UUID) bool { return r.members[userID] }
func (r *dmStub) SaveMessage(m *DMMessage) error               { r.messages = append(r.messages, m); return r.err }
func (r *dmStub) GetGroupMessages(groupID uuid.UUID, limit int) ([]*DMMessage, error) {
	return r.messages, nil
}
func (r *dmStub) GetLatestGroupMessage(groupID uuid.UUID) (*DMMessage, error) { return nil, nil }
func (r *dmStub) FindMessageByID(messageID uuid.UUID) (*DMMessage, error) {
	if r.foundMessage == nil || r.foundMessage.ID != messageID {
		return nil, errors.New("message not found")
	}
	return r.foundMessage, nil
}
func (r *dmStub) UpdateMessage(message *DMMessage) error {
	r.foundMessage = message
	return r.err
}
func (r *dmStub) DeleteMessage(messageID uuid.UUID) error {
	if r.foundMessage != nil && r.foundMessage.ID == messageID {
		r.foundMessage = nil
		return nil
	}
	return errors.New("message not found")
}

func TestCreateOrGetDirectGroupValidations(t *testing.T) {
	svc := NewService(&dmStub{})
	_, err := svc.CreateOrGetDirectGroup(dmUserA, dmUserA)
	if err == nil || err.Error() != "direct message requires two different users" {
		t.Fatalf("expected different users error")
	}
	_, err = svc.CreateOrGetDirectGroup(uuid.Nil, dmUserB)
	if err == nil {
		t.Fatalf("expected id required")
	}
}

func TestSendMessageValidationsAndSuccess(t *testing.T) {
	repo := &dmStub{members: map[uuid.UUID]bool{dmUserA: true}}
	svc := NewService(repo)
	_, err := svc.SendMessage(uuid.Nil, SendMessageRequest{GroupID: "", Content: "hi"})
	if err == nil || err.Error() != "groupId and userId are required" {
		t.Fatalf("expected required")
	}

	_, err = svc.SendMessage(dmUserA, SendMessageRequest{GroupID: "bad", Content: "hi"})
	if err == nil {
		t.Fatalf("expected invalid format")
	}

	repo.members = map[uuid.UUID]bool{}
	_, err = svc.SendMessage(dmUserA, SendMessageRequest{GroupID: uuid.New().String(), Content: "hi"})
	if err == nil || !errors.Is(err, ErrNotDMMember) {
		t.Fatalf("expected not member")
	}

	gid := uuid.New()
	repo.members = map[uuid.UUID]bool{dmUserA: true}
	repo.messages = []*DMMessage{}
	msg, err := svc.SendMessage(dmUserA, SendMessageRequest{GroupID: gid.String(), Content: "hello"})
	if err != nil {
		t.Fatalf("expected success got %v", err)
	}
	if msg.ID == uuid.Nil {
		t.Fatalf("expected id set")
	}
	if msg.CreatedAt.IsZero() || msg.UpdatedAt.IsZero() {
		t.Fatalf("expected timestamps set")
	}
}

func TestUpdateAndDeleteMessageAuthorChecks(t *testing.T) {
	message := &DMMessage{
		ID:          uuid.New(),
		GroupID:     uuid.New(),
		UserID:      dmUserA,
		Content:     "before",
		MessageTime: time.Now(),
	}
	repo := &dmStub{foundMessage: message}
	svc := NewService(repo)

	updated, err := svc.UpdateMessage(dmUserA, message.ID, "after")
	if err != nil || updated.Content != "after" {
		t.Fatalf("expected successful update, got updated=%v err=%v", updated, err)
	}

	if _, err := svc.UpdateMessage(dmUserB, message.ID, "hijack"); !errors.Is(err, ErrNotDMMember) {
		t.Fatalf("expected non-author update to be forbidden, got %v", err)
	}

	if err := svc.DeleteMessage(dmUserB, message.ID); !errors.Is(err, ErrNotDMMember) {
		t.Fatalf("expected non-author delete to be forbidden, got %v", err)
	}

	if err := svc.DeleteMessage(dmUserA, message.ID); err != nil {
		t.Fatalf("expected author delete to succeed, got %v", err)
	}
}
