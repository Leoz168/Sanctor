package institution

import (
    "testing"
    "github.com/google/uuid"
)

type instStub struct{
    created *Institution
    exists bool
}
func (r *instStub) Create(i *Institution) error { r.created = i; return nil }
func (r *instStub) FindByID(id uuid.UUID) (*Institution, error) { if r.created!=nil && r.created.ID==id { return r.created, nil }; return nil, errNotFound }
func (r *instStub) FindAll() []*Institution { return []*Institution{r.created} }
func (r *instStub) Update(i *Institution) error { r.created = i; return nil }
func (r *instStub) Delete(id uuid.UUID) error { r.created = nil; return nil }
func (r *instStub) ExistsByName(name string) bool { return r.exists }

var errNotFound = errorString("not found")
type errorString string
func (e errorString) Error() string { return string(e) }

func TestCreateInstitutionValidationAndDuplicate(t *testing.T){
    svc := NewService(&instStub{})
    _, err := svc.CreateInstitution(CreateInstitutionRequest{Name: ""})
    if err==nil || err.Error()!="institution name is required" { t.Fatalf("expected name required") }

    svc = NewService(&instStub{exists:true})
    _, err = svc.CreateInstitution(CreateInstitutionRequest{Name: "Name"})
    if err==nil || err.Error()!="institution already exists" { t.Fatalf("expected already exists") }
}

func TestGetInstitutionValidation(t *testing.T){
    svc := NewService(&instStub{})
    _, err := svc.GetInstitution(uuid.Nil)
    if err==nil || err.Error()!="institution ID is required" { t.Fatalf("expected id required") }
}
