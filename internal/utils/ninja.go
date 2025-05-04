package logger

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"

	"com.github/confusionhill/df/private/server/internal/config"
)

func DecryptNinja(dfConf *config.DragonFableConfig, encrypted string) (string, error) {
	var decrypted string

	textLength := len(encrypted)
	keyLength := len(dfConf.Ninja2Key)

	for i := 0; i < textLength; i += 4 {
		charP1, err := strconv.ParseInt(encrypted[i:i+2], 30, 64)
		if err != nil {
			return "", err
		}
		charP2, err := strconv.ParseInt(encrypted[i+2:i+4], 30, 64)
		if err != nil {
			return "", err
		}
		charP3 := int64(dfConf.Ninja2Key[(i/4)%keyLength])

		decrypted += string(charP1 - charP2 - charP3)
	}

	if !utf8.ValidString(decrypted) {
		return "", errors.New("invalid decrypted text. Verify the ninja2 key")
	}

	return decrypted, nil
}

func EncryptNinja(dfConf *config.DragonFableConfig, text string) string {
	var encrypted string
	textLength := len(text)
	keyLength := len(dfConf.Ninja2Key)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < textLength; i++ {
		random := rand.Intn(66) + 33 // random number between 33 and 98 inclusive
		char := dfConf.Ninja2Key[i%keyLength]
		charValue := int(char)

		// Calculate and convert
		sum := int(text[i]) + random + charValue
		encrypted += strconv.FormatInt(int64(sum), 30) + strconv.FormatInt(int64(random), 30)
	}

	return encrypted
}
