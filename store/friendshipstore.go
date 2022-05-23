package store

import (
	"fmt"
	"time"

	"telegraph/model"

	"github.com/satori/go.uuid"
)

func (f *FriendshipStore) Create(friendship *model.Friendship) (*model.Friendship, error) {
	// Check for self friend request.
	if friendship.Sender == friendship.Recipient {
		return nil, fmt.Errorf("Invalid friendship request.")
	}

	err := f.db.Where(friendship).Or(friendship.Invert()).
		Attrs(model.Friendship{ID: uuid.NewV4().String(), Status: model.FriendshipStatusEnum.PENDING}).
		FirstOrCreate(friendship).Error

	return friendship, err
}

func (f *FriendshipStore) Delete(ID string) (friendship *model.Friendship, err error) {
	friendship, err = f.Get(ID)
	if err != nil {
		return
	}

	friendship.DeletedAt = time.Now()
	err = f.db.Unscoped().Where("id = ?", ID).Delete(friendship).Error

	return
}

func (f *FriendshipStore) Accept(ID string) (friendship *model.Friendship, err error) {
	friendship, err = f.Get(ID)
	if err != nil {
		return
	}

	err = f.db.Model(friendship).Where("id = ?", ID).Update("status", model.FriendshipStatusEnum.ACCEPTED).Error
	if err != nil {
		return
	}

	// Create RBAC reference.

	return
}

func (f *FriendshipStore) Get(ID string) (friendship *model.Friendship, err error) {
	err = f.db.Where("id = ?", ID).First(&friendship).Error

	return
}

func (f *FriendshipStore) GetMany(ID string, status *string) (friendships []*model.Friendship, err error) {
	query := f.db.Where(f.db.Where("sender = ?", ID).Or("recipient = ?", ID))
	if status != nil {
		query.Where("status = ?", status)
	}
	err = query.Find(&friendships).Error

	return
}
