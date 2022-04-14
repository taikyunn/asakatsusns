package tests

import (
	controller "app/controllers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type users struct {
	Name     string
	Email    string
	Password string
}

type preferResponse struct {
	code int                    //httpステータスコード
	body map[string]interface{} //帰ってくる文字列
}

func TestSignUp(t *testing.T) {
	userTests := []struct {
		users users
		want  preferResponse
	}{
		{
			// 空登録しようとした場合
			users{
				Email:    "",
				Password: "",
				Name:     "",
			},
			preferResponse{
				code: 201,
				body: map[string]interface{}{
					"Email": "*メールアドレスは必須入力です。",
					//"Password": "test password", // hash化されるため省略
					"Name": "*お名前は必須入力です。",
				},
			},
		},
	}
	for i, tt := range userTests {
		requestBody := strings.NewReader("Email=" + tt.users.Email + "&Name=" + tt.users.Name + "&Password=" + tt.users.Password)
		// レスポンス
		response := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/signUp",
			requestBody,
		)

		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// テスト実行
		controller.SignUp(c)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

		// ステータスコードがおかしいもしくは帰ってきたメッセージが想定と違えばダメ
		if response.Code != tt.want.code {
			t.Errorf("%d番目のテストが失敗しました。想定返却コード：%d, 実際の返却コード：%d", i+1, tt.want.code, response.Code)
		} else {
			// 実際に帰ってきたレスポンスの中に想定された値が入っているかどうか
			for key := range tt.want.body {
				// 値の存在チェック
				if _, exist := responseBody[key]; exist {
					// 値の内容チェック
					if responseBody[key] != tt.want.body[key] {
						t.Errorf("%d番目のテストが失敗しました。想定されたキー「%s」の値:%s, 実際に返却された値:%s", i+1, key, tt.want.body[key], responseBody[key])
					}
				} else {
					t.Errorf("%d番目のテストが失敗しました。想定された「%s」がレスポンスボディに入っていません。", i+1, key)
				}
			}
		}
	}
}
