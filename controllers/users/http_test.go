package users_test

// import (
// 	"encoding/json"
// 	_userMock "gym-membership/business/users/mocks"
// 	"gym-membership/controllers"
// 	"gym-membership/controllers/users"
// 	"gym-membership/helper/encrypt"
// 	_encryptMock "gym-membership/helper/encrypt/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type Requests struct {
// 	Register				string
// 	RegisterInvalidBind		string
// 	RegisterInvalidStruct	string
// 	Login					string
// 	LoginInvalidBind		string
// 	LoginInvalidStruct		string
// }

// var (
// 	mockUserUC						_userMock.Usecase
// 	mockEncrypt						_encryptMock.Helper
// 	userCtrl						users.UserController
// 	jsonReq							Requests
// 	hashedPassword					string
// )

// func TestMain(m *testing.M){
// 	userCtrl = *users.NewUserController(&mockUserUC)
// 	hashedPassword, _ = encrypt.Hash("testpassword")
// 	jsonReq.Register = `{
// 		"username": "tonotono",
// 		"password": "passwordtono",
// 		"email": "tono@gmail.com",
// 		"fullname": "tono sutono",
// 		"gender": "male",
// 		"telephone": "8123456789",
// 		"address": "jl dr sutomo 106"
// 	}`
// 	jsonReq.RegisterInvalidBind = `{
// 		"username": "tonotono"
// 		"password": "passwordtono",
// 		"email": "tono@gmail.com",
// 		"fullname": "tono sutono",
// 		"gender": "male",
// 		"telephone": "8123456789",
// 		"address": "jl dr sutomo 106"
// 	}`
// 	jsonReq.RegisterInvalidStruct = `{
// 		"username": "tono",
// 		"password": "passwordtono",
// 		"email": "tono@gmail.com",
// 		"fullname": "tono sutono",
// 		"gender": "male",
// 		"telephone": "8123456789",
// 		"address": "jl dr sutomo 106"
// 	}`
// 	jsonReq.Login = `{
// 		"username": "tonotono",
// 		"password": "passwordtono"
// 	}`
// 	jsonReq.LoginInvalidBind = `{
// 		"username": "tonotono"
// 		"password": "passwordtono"
// 	}`
// 	jsonReq.LoginInvalidStruct = `{
// 		"username": "tono",
// 		"password": "passwordtono"
// 	}`
// 	m.Run()
// }

// func TestRegister(t *testing.T){
// 	t.Run("Valid Test", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(jsonReq.Register))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)
// 		mockUserUC.On("Register", mock.Anything).Return("item created", nil).Once()

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusOK
// 		resp.Meta.Message = "Success"
// 		resp.Data = "item created"
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Register(c)){
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// 	t.Run("Invalid Test | Bind Error", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(jsonReq.RegisterInvalidBind))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusBadRequest
// 		resp.Meta.Message = "Error"
// 		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=30, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Register(c)){
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// 	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(jsonReq.RegisterInvalidStruct))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusBadRequest
// 		resp.Meta.Message = "Error"
// 		resp.Meta.Messages = []string{"username: tono does not validate as minstringlength(6)"}
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Register(c)){
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// 	t.Run("Invalid Test | Conflict", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(jsonReq.Register))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)
// 		mockUserUC.On("Register", mock.Anything).Return("", assert.AnError).Once()

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusConflict
// 		resp.Meta.Message = "Error"
// 		resp.Meta.Messages = []string{assert.AnError.Error()}
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Register(c)){
// 			assert.Equal(t, http.StatusConflict, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// }

// func TestLogin(t *testing.T){
// 	t.Run("Valid Test", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(jsonReq.Login))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)
// 		mockUserUC.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("token", nil).Once()

// 		token := struct {
// 			Token string `json:"token"`
// 		}{Token: "token"}

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusOK
// 		resp.Meta.Message = "Success"
// 		resp.Data = token
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Login(c)){
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// 	t.Run("Invalid Test | Bind Error", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(jsonReq.LoginInvalidBind))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusBadRequest
// 		resp.Meta.Message = "Error"
// 		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=30, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Login(c)){
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// 	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(jsonReq.LoginInvalidStruct))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusBadRequest
// 		resp.Meta.Message = "Error"
// 		resp.Meta.Messages = []string{"username: tono does not validate as minstringlength(6)"}
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Login(c)){
// 			assert.Equal(t, http.StatusBadRequest, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// 	t.Run("Invalid Test | Invalid Username/Password", func(t *testing.T){
// 		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(jsonReq.Login))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		rec := httptest.NewRecorder()
// 		e := echo.New()
// 		c := e.NewContext(req,rec)
// 		mockUserUC.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", assert.AnError).Once()

// 		resp := controllers.BaseResponse{}
// 		resp.Meta.Status = http.StatusUnauthorized
// 		resp.Meta.Message = "Error"
// 		resp.Meta.Messages = []string{assert.AnError.Error()}
// 		expected, _ := json.Marshal(resp)

// 		if assert.NoError(t, userCtrl.Login(c)){
// 			assert.Equal(t, http.StatusUnauthorized, rec.Code)
// 			assert.JSONEq(t, string(expected), rec.Body.String())
// 		}
// 	})
// }