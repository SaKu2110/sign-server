package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/sign-server/models"
	"github.com/SaKu2110/sign-server/service/auth"
)

func (ctrl *CCH) SignInHandler (g *gin.Context) {
	var user, dbuser models.User
	var auth = &auth.Auth{}
	ctrl.ERROR = g.BindJSON(&user)
	// JSONデータの取り込みに失敗
	if ctrl.ERROR != nil {
		return
	}
	// 検索
	ctrl.DB.Find(&dbuser, "id=? and password=?", user.Id, user.Password)
	if dbuser.Id == user.Id && dbuser.Password == user.Password {
		auth.GetToken(user.Id, "signin")
		// トークンの生成に失敗
		if auth.ERROR != nil {
			ctrl.ERROR = auth.ERROR
			return
		}
                g.JSON(200, gin.H{"access_token": auth.Token})
        }else{
                g.JSON(400, gin.H{"message": "Bad Request"})
        }

}
