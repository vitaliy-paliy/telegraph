package store

import (
	"fmt"

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

func (f *FriendshipStore) Accept(friendshipID string) (*model.Friendship, error) {
	friendship, err := f.Get(friendshipID)
	if err != nil {
		return nil, fmt.Errorf("An error occured while trying to accept a friendship. Error: %s.", err)
	}

	friendship.Status = model.FriendshipStatusEnum.ACCEPTED
	err = f.db.Where("sender = ? AND recipient = ?", friendship.Sender, friendship.Recipient).Select("updated_at", "status").Updates(friendship).Error

	return friendship, err
}

func (f *FriendshipStore) Get(friendshipID string) (*model.Friendship, error) {
	friendship := &model.Friendship{ID: friendshipID}

	err := f.db.Where("id = ?", friendshipID).Find(friendship).Error

	return friendship, err
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
