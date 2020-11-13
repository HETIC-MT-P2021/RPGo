package create

import (
	"github.com/HETIC-MT-P2021/RPGo/mock_repository"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/golang/mock/gomock"
	"testing"
)

var char = &repository.Character{
	Name:          "TestChar",
	Class:         "Ranger",
	DiscordUserID: "1234",
}

func TestCharCommandGenerator_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repository.NewMockCharacterRepositoryInterface(ctrl)

	m.EXPECT().GetCharacterByDiscordUserID(char.DiscordUserID).DoAndReturn(func() (*repository.
		Character, error) {
		return nil, nil
	}).MaxTimes(1)

}
