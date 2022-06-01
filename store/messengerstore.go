package store

import (
	"time"

	"telegraph/model"

	"github.com/satori/go.uuid"
)

func (m *MessengerStore) Create(message *model.Message) (*model.Message, error) {
	message.ID = uuid.NewV4().String()
	message.Status = model.MessageStatusEnum.SENT
	err := m.db.Create(message).Error

	return message, err
}

func (m *MessengerStore) Delete(messageID string) (message *model.Message, err error) {
	message, err = m.Get(messageID)
	if err != nil {
		return
	}

	message.DeletedAt = time.Now()
	err = m.db.Unscoped().Where("id = ?", messageID).Delete(message).Error

	return
}

func (m *MessengerStore) Read(messageID string) (message *model.Message, err error) {
	message, err = m.Get(messageID)
	if err != nil {
		return
	}

	err = m.db.Model(message).Where("id = ?", messageID).Update("status = ?", model.MessageStatusEnum.READ).Error

	return
}

func (m *MessengerStore) Get(messageID string) (message *model.Message, err error) {
	err = m.db.Where("id = ?", messageID).First(&message).Error

	return
}

func (m *MessengerStore) GetMany(conversationID string) (messages []*model.Message, err error) {
	err = m.db.Where("conversation_id = ?", conversationID).Find(&messages).Error

	return
}
