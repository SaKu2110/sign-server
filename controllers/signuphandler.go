package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/sign-server/models"
	"github.com/SaKu2110/sign-server/service/auth"
)

func (ctrl *CCH) SignUpHandler (g *gin.Context) {
	var user, dbuser models.User
	var auth = &auth.Auth{}

	ctrl.ERROR = g.BindJSON(&user)
	// JSON形式のデータの取り込みに失敗
	if ctrl.ERROR != nil {
		return
	}

	// 検索
	ctrl.DB.Find(&dbuser, "id=?", user.Id)

	if dbuser.Id != user.Id {
		ctrl.DB.Create(&user)

		auth.GetToken(user.Id, "signup")
		// トークンの生成に失敗
		if auth.ERROR != nil {
			ctrl.ERROR = auth.ERROR
			return
		}

		g.JSON(201, gin.H{"access_token": auth.Token})
        }else{
                g.JSON(412, gin.H{"message": "Precondition Failed"})
        }

}
