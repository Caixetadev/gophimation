package presence

import (
	"errors"
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func Presence(name, imageLarge, nameAnime, state string) {
	err := client.Login("1075841986923352079")

	if err != nil {
		err := errors.New("nao foi possivel ativar a conexao com o discord")
		fmt.Println(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      state,
		Details:    nameAnime,
		LargeImage: imageLarge,
		LargeText:  nameAnime,
		SmallImage: "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png",
		SmallText:  "Gophimation",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})

	if err != nil {
		panic(err)
	}
}
