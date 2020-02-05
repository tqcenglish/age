package age

import (
	"io"

	"filippo.io/age/internal/age"
)

//EncryptPass 加密
func EncryptPass(pass string, in io.Reader, out io.Writer, armor bool) {
	r, err := age.NewScryptRecipient(pass)
	if err != nil {
		logFatalf("Error: %v", err)
	}
	encrypt([]age.Recipient{r}, in, out, armor)
}
