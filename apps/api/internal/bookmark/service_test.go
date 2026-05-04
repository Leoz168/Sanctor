package bookmark

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
    bmUser = uuid.MustParse("11111111-1111-1111-1111-111111111111")
    bmPost = uuid.MustParse("22222222-2222-2222-2222-222222222222")
)

type bmStubRepo struct{
    createErr error
    exists bool
    existsErr error
    existsUser map[uuid.UUID]bool
    existsPost map[uuid.UUID]bool
    created *Bookmark
}

func (r *bmStubRepo) Create(b *Bookmark) error { r.created = b; return r.createErr }
func (r *bmStubRepo) Delete(userID, postID uuid.UUID) error { return nil }
func (r *bmStubRepo) FindByUserID(userID uuid.UUID) ([]*Bookmark, error) { return []*Bookmark{r.created}, nil }
func (r *bmStubRepo) FindPostsByUserID(userID uuid.UUID) ([]*BookmarkedPost, error) { return nil, nil }
func (r *bmStubRepo) Exists(userID, postID uuid.UUID) (bool, error) { return r.exists, r.existsErr }
func (r *bmStubRepo) ExistsPost(postID uuid.UUID) bool { if r.existsPost==nil { return postID!=uuid.Nil }; return r.existsPost[postID] }
func (r *bmStubRepo) ExistsUser(userID uuid.UUID) bool { if r.existsUser==nil { return userID!=uuid.Nil }; return r.existsUser[userID] }

func TestCreateBookmarkValidationAndSuccess(t *testing.T){
    cases := []struct{ name string; req CreateBookmarkRequest; wantErr string }{
        {"empty-user", CreateBookmarkRequest{UserID: "", PostID: bmPost.String()}, "user ID is required"},
        {"invalid-user", CreateBookmarkRequest{UserID: "not-uuid", PostID: bmPost.String()}, "invalid user ID format"},
        {"empty-post", CreateBookmarkRequest{UserID: bmUser.String(), PostID: ""}, "post ID is required"},
        {"invalid-post", CreateBookmarkRequest{UserID: bmUser.String(), PostID: "bad"}, "invalid post ID format"},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T){
            svc := NewService(&bmStubRepo{})
            _, err := svc.CreateBookmark(tc.req)
            if err==nil || err.Error()!=tc.wantErr { t.Fatalf("expected %s got %v", tc.wantErr, err) }
        })
    }

    // missing user/post
    repo := &bmStubRepo{existsUser: map[uuid.UUID]bool{bmUser:false}, existsPost: map[uuid.UUID]bool{bmPost:true}}
    svc := NewService(repo)
    _, err := svc.CreateBookmark(CreateBookmarkRequest{UserID: bmUser.String(), PostID: bmPost.String()})
    if err==nil || err.Error()!="user not found" { t.Fatalf("expected user not found got %v", err) }

    // successful create
    repo = &bmStubRepo{existsUser: map[uuid.UUID]bool{bmUser:true}, existsPost: map[uuid.UUID]bool{bmPost:true}}
    svc = NewService(repo)
    start := time.Now()
    b, err := svc.CreateBookmark(CreateBookmarkRequest{UserID: bmUser.String(), PostID: bmPost.String()})
    if err!=nil { t.Fatalf("expected success got %v", err) }
    if b.UserID!=bmUser || b.PostID!=bmPost { t.Fatalf("ids mismatch") }
    if b.CreatedAt.Before(start) { t.Fatalf("created at not set") }
}

func TestIsBookmarkedValidation(t *testing.T){
    svc := NewService(&bmStubRepo{})
    _, err := svc.IsBookmarked(uuid.Nil, bmPost)
    if err==nil || err.Error()!="user ID is required" { t.Fatalf("expected user ID required") }
}
