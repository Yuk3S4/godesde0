package users

import (
	"fmt"
	"time"

	"github.com/Yuk3S4/godesde0/models"
)

func AltaUsuario() {
	u := new(models.User)
	u.AddUser(10, "Daniel", time.Now(), true)
	fmt.Println(u)
}
