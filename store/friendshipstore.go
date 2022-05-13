package store

import (
	"fmt"
	"time"

	"telegraph/model"

	"github.com/satori/go.uuid"
)

func (f *FriendshipStore) Create(friendship *model.Friendship) (*model.Friendship, error) {
	if friendship, err := f.isDuplicate(friendship); friendship != nil {
		return friendship, err
	}

	if friendship.Sender == friendship.Recipient {
		return nil, fmt.Errorf("Invalid friendship request.")
	}

	friendship.ID = uuid.NewV4().String()
	friendship.Status = model.FriendshipStatusEnum.PENDING

	err := f.db.Create(friendship).Error
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

	friendship.Status = model.FriendshipStatusEnum.ACCEPTED
	err = f.db.Where("sender = ? AND recipient = ?", friendship.Sender, friendship.Recipient).Select("updated_at", "status").Updates(friendship).Error

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

// FriendshipStore helper methods.

func (f *FriendshipStore) isDuplicate(friendship *model.Friendship) (*model.Friendship, error) {
	invertedFriendship := friendship.Invert()
	if f.isPresent(invertedFriendship) {
		switch invertedFriendship.Status {
		case model.FriendshipStatusEnum.ACCEPTED:
			return invertedFriendship, nil
		case model.FriendshipStatusEnum.PENDING:
			return f.Accept(invertedFriendship.ID)
		default:
			return invertedFriendship, fmt.Errorf("%v status is not implemented.", invertedFriendship.Status)
		}
	}

	if f.isPresent(friendship) {
		return friendship, nil
	}

	// No duplicate friendship is found, create a new one.
	return nil, nil
}

func (f *FriendshipStore) isPresent(friendship *model.Friendship) bool {
	return f.db.Where("sender = ? AND recipient = ?", friendship.Sender, friendship.Recipient).Find(friendship).RowsAffected != 0
}
