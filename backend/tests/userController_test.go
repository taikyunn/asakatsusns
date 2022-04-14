package tests

import (
	controller "app/controllers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
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

// ユーザー登録
func TestSignUp(t *testing.T) {
	signUpTests := []struct {
		users users
		want  preferResponse
	}{
		{
			// 名前を空登録しようとした場合
			users{
				Email:    "testtest@example.com",
				Password: "testtest",
				Name:     "",
			},
			preferResponse{
				code: 201,
				body: map[string]interface{}{
					//"Password": "test password", // hash化されるため省略
					"Name": "*お名前は必須入力です。",
				},
			},
		},
		{
			// メールアドレスを空登録しようとした場合
			users{
				Email:    "",
				Password: "testtest",
				Name:     "testuser",
			},
			preferResponse{
				code: 201,
				body: map[string]interface{}{
					//"Password": "test password",
					"Email": "*メールアドレスは必須入力です。",
				},
			},
		},
		{
			// メールアドレスの形式ミス
			users{
				Email:    "testtest",
				Password: "testtest",
				Name:     "testuser",
			},
			preferResponse{
				code: 201,
				body: map[string]interface{}{
					//"Password": "test password",
					"Email": "*正しいメールアドレス形式で入力してください。",
				},
			},
		},
		{
			// パスワードを空で登録しようとした場合
			users{
				Email:    "testtest@example.com",
				Password: "",
				Name:     "testuser",
			},
			preferResponse{
				code: 201,
				body: map[string]interface{}{
					"Password": "*パスワードは必須入力です。",
				},
			},
		},
	}
	for i, tt := range signUpTests {
		form := url.Values{}
		form.Add("name", tt.users.Name)
		form.Add("email", tt.users.Email)
		form.Add("password", tt.users.Password)
		requestBody := strings.NewReader(form.Encode())
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

		// ステータスコードチェック
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

// ログイン処理
func TestLogin(t *testing.T) {
	loginTests := []struct {
		users users
		want  preferResponse
	}{
		{
			// 名前を空入力した場合
			users{
				Email:    "testtest@example.com",
				Password: "testtest",
				Name:     "",
			},
			preferResponse{
				code: 201,
				body: map[string]interface{}{
					"Name": "*お名前は必須入力です。",
				},
			},
		},
		{
			// パスワードを空入力した場合
			users{
				Email:    "testtest@example.com",
				Password: "",
				Name:     "testuser",
			},
			preferResponse{
				code: 201,
				body: map[string]interface{}{
					"Password": "*パスワードは必須入力です。",
				},
			},
		},
	}

	for i, tt := range loginTests {
		form := url.Values{}
		form.Add("name", tt.users.Name)
		form.Add("email", tt.users.Email)
		requestBody := strings.NewReader(form.Encode())
		// レスポンス
		response := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/login",
			requestBody,
		)

		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// テスト実行
		controller.SignUp(c)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

		// ステータスコードチェック
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
