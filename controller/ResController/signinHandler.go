package ResController

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/sign_server/model/LoginModel"
)

func (rctl *ResCnt) SigninHandler (g *gin.Context) {
	var key, keyhole LoginModel.Login

	g.BindJSON(&key)

	rctl.DB.Find(&keyhole, "id=? and password=?", key.Id, key.Password)
	if keyhole.Id == key.Id && keyhole.Password == key.Password {
		rctl.AuthCnt.MakeToken(key.Id, "signin")
                g.JSON(200, gin.H{"access_token": rctl.AuthCnt.Token})
        }else{
                g.JSON(400, gin.H{"message": "Bad Request"})
        }
}
