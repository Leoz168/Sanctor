package post

import (
    "context"
    "errors"
    "testing"
    "github.com/google/uuid"
)

type postStubRepo struct{
    created *Post
    findErr error
}
func (r *postStubRepo) CreateWithLinks(post *Post, communityIDs []uuid.UUID, institutionIDs []uuid.UUID) (*Post, error){ r.created = post; return post, nil }
func (r *postStubRepo) FindByID(id uuid.UUID) (*Post, error){ if r.findErr!=nil { return nil, r.findErr }; return r.created, nil }
func (r *postStubRepo) FindAll() ([]*Post, error){ return []*Post{}, nil }
func (r *postStubRepo) Search(filters PostSearchFilters) ([]*Post, error){ return []*Post{}, nil }
func (r *postStubRepo) Update(post *Post) error { return nil }
func (r *postStubRepo) Delete(id uuid.UUID) error { return nil }

type stubPublisher struct{}
func (p *stubPublisher) PublishPostCreated(ctx context.Context, event PostCreatedEvent) error { return nil }
func (p *stubPublisher) PublishPostUpdated(ctx context.Context, event PostUpdatedEvent) error { return nil }
func (p *stubPublisher) PublishPostDeleted(ctx context.Context, event PostDeletedEvent) error { return nil }
func (p *stubPublisher) PublishPostViewed(ctx context.Context, event PostViewedEvent) error { return nil }

func TestParseUUIDsInvalid(t *testing.T){
    _, err := parseUUIDs([]string{"not-uuid"})
    if err==nil { t.Fatalf("expected invalid uuid error") }
}

func TestUniqueIDsFilters(t *testing.T){
    out := uniqueIDs([]string{"a","a","","b"})
    if len(out)!=2 { t.Fatalf("expected 2 unique entries got %v", out) }
}

func TestCreatePostValidationUser(t *testing.T){
    svc := NewService(&postStubRepo{}, &stubPublisher{})
    _, err := svc.CreatePost(context.Background(), &CreatePostRequest{UserID: ""})
    if err==nil || err.Error()!="userId is required" { t.Fatalf("expected userId required") }
    _, err = svc.CreatePost(context.Background(), &CreatePostRequest{UserID: "bad-uuid"})
    if err==nil { t.Fatalf("expected invalid user id") }
}

func TestGetPostNotFoundAndPublish(t *testing.T){
    repo := &postStubRepo{findErr: errors.New("not found")}
    svc := NewService(repo, &stubPublisher{})
    _, err := svc.GetPost(context.Background(), uuid.New())
    if err==nil { t.Fatalf("expected not found") }
}
