package comment

import (
    "errors"
    "testing"
    "time"

    "github.com/google/uuid"
)

var (
    cPost = uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
    cUser = uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")
)

type cStubRepo struct{
    createErr error
    existsPost map[uuid.UUID]bool
    existsUser map[uuid.UUID]bool
    stored *Comment
}

func (r *cStubRepo) Create(comment *Comment) error { r.stored = comment; return r.createErr }
func (r *cStubRepo) FindByID(id uuid.UUID) (*Comment, error) { if r.stored==nil { return nil, errors.New("not found") }; return r.stored, nil }
func (r *cStubRepo) FindByPostID(postID uuid.UUID) []*Comment { return []*Comment{r.stored} }
func (r *cStubRepo) Update(comment *Comment) error { r.stored = comment; return nil }
func (r *cStubRepo) Delete(id uuid.UUID) error { r.stored = nil; return nil }
func (r *cStubRepo) ExistsPost(postID uuid.UUID) bool { if r.existsPost==nil { return true }; return r.existsPost[postID] }
func (r *cStubRepo) ExistsUser(userID uuid.UUID) bool { if r.existsUser==nil { return true }; return r.existsUser[userID] }

func TestCreateCommentValidationAndSuccess(t *testing.T){
    svc := NewService(&cStubRepo{})
    _, err := svc.CreateComment(CreateCommentRequest{PostID: "", CreatedByUserID: cUser.String(), Content: "hi"})
    if err==nil || err.Error()!="post ID is required" { t.Fatalf("expected post id required") }

    _, err = svc.CreateComment(CreateCommentRequest{PostID: cPost.String(), CreatedByUserID: "", Content: "hi"})
    if err==nil || err.Error()!="created by user ID is required" { t.Fatalf("expected created by user id required") }

    long := ""
    for i:=0;i<2001;i++{ long += "x" }
    _, err = svc.CreateComment(CreateCommentRequest{PostID: cPost.String(), CreatedByUserID: cUser.String(), Content: long})
    if err==nil || err.Error()!="content must be 2000 characters or fewer" { t.Fatalf("expected content length err") }

    // missing post
    repo := &cStubRepo{existsPost: map[uuid.UUID]bool{cPost:false}}
    svc = NewService(repo)
    _, err = svc.CreateComment(CreateCommentRequest{PostID: cPost.String(), CreatedByUserID: cUser.String(), Content: "ok"})
    if err==nil || err.Error()!="post not found" { t.Fatalf("expected post not found") }

    // success
    repo = &cStubRepo{existsPost: map[uuid.UUID]bool{cPost:true}, existsUser: map[uuid.UUID]bool{cUser:true}}
    svc = NewService(repo)
    before := time.Now()
    cm, err := svc.CreateComment(CreateCommentRequest{PostID: cPost.String(), CreatedByUserID: cUser.String(), Content: "hello"})
    if err!=nil { t.Fatalf("expected success got %v", err) }
    if cm.ID==uuid.Nil || cm.CreatedAt.Before(before) { t.Fatalf("expected timestamps set") }
}

func TestUpdateCommentValidation(t *testing.T){
    stored := &Comment{ID: uuid.New(), Content: "old"}
    repo := &cStubRepo{stored: stored}
    svc := NewService(repo)
    _, err := svc.UpdateComment(stored.ID, UpdateCommentRequest{Content: ""})
    if err==nil || err.Error()!="content is required" { t.Fatalf("expected content required") }
}
