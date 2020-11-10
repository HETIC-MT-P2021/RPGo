package create

import (
	"github.com/HETIC-MT-P2021/RPGo/database"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	. "github.com/onsi/gomega"
	"log"
	"testing"
)

func TestMakeCreateCommand(t *testing.T) {
	g := NewGomegaWithT(t)

	db := database.DBCon
	repo := repository.Repository{Conn: db}

	userId := "1"

	char, err := repo.GetCharacterByDiscordUserID("1")
	if err != nil {
		log.Fatalf("Couldn't get character with userID: %s, %v", userId, err)
	}

	character := repository.Character{}

	// Create a character if none found in DB
	if char == nil {
		character := repository.Character{
			Name:          "Tynyndil",
			Class:         "Ranger",
			DiscordUserID: userId,
		}
		err := repo.CreateACharacter(&character)
		if err != nil {
			log.Fatalf("Couldn't create character: %v", err)
		}
	}

	g.Expect(character).ShouldNot(BeNil(), "Character should have been created")

	char, err = repo.GetCharacterByDiscordUserID(string("1"))

	if err != nil {
		log.Fatalf("Couldn't get character with userID: %s, %v", "1", err)
	}

	g.Expect(char).ShouldNot(BeNil(), "User can no longer create a character")

}

