package presence

import (
	"errors"
	"fmt"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func Presence(name, imageLarge, nameAnime, state, smallImage string) {
	err := client.Login("1075841986923352079")

	if err != nil {
		err := errors.New("Aviso: Não foi possivel ativar a conexao com o discord\n")
		fmt.Println(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      state,
		Details:    nameAnime,
		LargeImage: imageLarge,
		LargeText:  nameAnime,
		SmallImage: smallImage,
		SmallText:  "Gophimation",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			{
				Label: "GitHub",
				Url:   "https://github.com/caixetadev",
			},
		},
	})

	if err != nil {
		panic(err)
	}
}
