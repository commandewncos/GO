package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

// New campaign for send e-mails
func NewCampaign(name, content string, emails []string) (*Campaign, error) {

	switch {
	case name == "":
		return nil, errors.New("field 'name' is required")
	case content == "":
		return nil, errors.New("field 'content' is required")
	case len(emails) == 0:
		return nil, errors.New("e-mail is required")
	}

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        int(xid.New().Counter()),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}, nil

}
