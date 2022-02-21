package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	someMessage := emoji.Sprint("Hello :world_map:!")
	return someMessage
}
