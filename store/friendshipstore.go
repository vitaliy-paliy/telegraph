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
		return nil, fmt.Errorf("Error: invalid friendship request.")
	}

	err := f.db.Where(friendship).Or(friendship.Invert()).
		Attrs(model.Friendship{ID: uuid.NewV4().String(), Status: model.FriendshipStatusEnum.PENDING}).
		FirstOrCreate(friendship).Error

	// Create request policy
	senderPolicy := []string{friendship.Sender, friendship.ID, model.FriendshipPolicyEnum.SENDER}
	if hasPolicy := f.enf.HasPolicy(senderPolicy); !hasPolicy {
		f.enf.AddPolicy(senderPolicy)
	}

	recipientPolicy := []string{friendship.Recipient, friendship.ID, model.FriendshipPolicyEnum.RECIPIENT}
	if hasPolicy := f.enf.HasPolicy(recipientPolicy); !hasPolicy {
		f.enf.AddPolicy(recipientPolicy)
	}

	return friendship, err
}

func (f *FriendshipStore) Delete(ID string) (friendship *model.Friendship, err error) {
	friendship, err = f.Get(ID)
	if err != nil {
		return
	}

	friendship.DeletedAt = time.Now()
	err = f.db.Unscoped().Where("id = ?", ID).Delete(friendship).Error
	if err != nil {
		return
	}

	// Delete casbin policy.
	for _, userID := range []string{friendship.Sender, friendship.Recipient} {
		if hasPolicy := f.enf.HasPolicy(userID, ID, model.FriendshipPolicyEnum.FRIEND); hasPolicy {
			f.enf.RemovePolicy(userID, ID, model.FriendshipPolicyEnum.FRIEND)
		}
	}

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

	// Create friendship policy.
	for _, userID := range []string{friendship.Sender, friendship.Recipient} {
		if hasPolicy := f.enf.HasPolicy(userID, ID, model.FriendshipPolicyEnum.FRIEND); !hasPolicy {
			f.enf.AddPolicy(userID, ID, model.FriendshipPolicyEnum.FRIEND)
		}
	}

	f.deleteRequestPolicies(friendship)

	return
}

func (f *FriendshipStore) Cancel(ID string) (friendship *model.Friendship, err error) {
	friendship, err = f.Get(ID)
	if err != nil {
		return
	}

	if friendship.Status != model.FriendshipStatusEnum.PENDING {
		err = fmt.Errorf("Error: cannot cancel non-pending friend request.")
		return
	}

	err = f.db.Unscoped().Where("id = ?", ID).Delete(friendship).Error
	if err != nil {
		return
	}

	f.deleteRequestPolicies(friendship)

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

// Handlers

func (f *FriendshipStore) deleteRequestPolicies(friendship *model.Friendship) {
	if hasPolicy := f.enf.HasPolicy(friendship.Sender, friendship.ID, model.FriendshipPolicyEnum.SENDER); hasPolicy {
		f.enf.RemovePolicy(friendship.Sender, friendship.ID, model.FriendshipPolicyEnum.SENDER)
	}

	if hasPolicy := f.enf.HasPolicy(friendship.Recipient, friendship.ID, model.FriendshipPolicyEnum.RECIPIENT); hasPolicy {
		f.enf.RemovePolicy(friendship.Recipient, friendship.ID, model.FriendshipPolicyEnum.RECIPIENT)
	}
}
