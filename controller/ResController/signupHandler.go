package ResController

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/sign_server/model/LoginModel"
)

func (rctl *ResCnt) SignupHandler (g *gin.Context) {
	var key, keyhole LoginModel.Login

	g.BindJSON(&key)

	rctl.DB.Find(&keyhole, "id=?", key.Id)
	if keyhole.Id == key.Id {
		g.JSON(412, gin.H{"message": "Precondition Failed"})
	}else{
		rctl.DB.Create(&key)
		rctl.AuthCnt.MakeToken(key.Id, "signup")
		g.JSON(201, gin.H{"token": rctl.AuthCnt.Token})
	}
}
