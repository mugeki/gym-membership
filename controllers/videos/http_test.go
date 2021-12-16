package videos_test

import (
	"encoding/json"
	_videoBusiness "gym-membership/business/videos"
	_videoMock "gym-membership/business/videos/mocks"
	"gym-membership/controllers"
	"gym-membership/controllers/videos"
	"gym-membership/controllers/videos/response"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Requests struct {
	Video				string
	VideoInvalidBind	string
	VideoInvalidStruct	string
}

var (
	mockVideoUC	 _videoMock.Usecase
	videoCtrl	 videos.VideoController
	videoData	 _videoBusiness.Domain
	jsonReq		 Requests
)

func TestMain(m *testing.M){
	videoCtrl = *videos.NewVideoController(&mockVideoUC)
	jsonReq.Video = `{
		"title": "Test Title",
		"classification": "test class",
		"memberOnly": true,
		"url": "https://www.youtube.com/watch?v=80AjI0hlbf8"
	}`
	jsonReq.VideoInvalidBind = `{
		"title": "Test Title"
		"classification": "test class",
		"memberOnly": true,
		"url": "https://www.youtube.com/watch?v=80AjI0hlbf8"
	}`
	jsonReq.VideoInvalidStruct = `{
		"title": "Test Title",
		"classification": "test class",
		"memberOnly": true,
		"url": "asd"
	}`
	videoData = _videoBusiness.Domain{
		ID                	: 1,
		Title             	: "Test Title",
		ClassificationID  	: 1,
		ClassificationName	: "test class",
		AdminID           	: 1,
		MemberOnly        	: true,
		Url               	: "https://www.youtube.com/watch?v=80AjI0hlbf8",
		CreatedAt         	: time.Date(2021,12,1,0,0,0,0,time.UTC),
	}
	m.Run()
}

func TestGetAll(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/videos", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		data := []response.Videos{}
		copier.Copy(&data, &videoData)
		mockVideoUC.On("GetAll").Return([]_videoBusiness.Domain{videoData}, nil).Once()

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = data
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.GetAll(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/videos", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockVideoUC.On("GetAll").Return([]_videoBusiness.Domain{}, assert.AnError).Once()

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{assert.AnError.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.GetAll(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestInsert(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos", strings.NewReader(jsonReq.Video))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockVideoUC.On("Insert", mock.Anything, mock.AnythingOfType("uint")).Return("item created", nil).Once()

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = "item created"
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.Insert(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos", strings.NewReader(jsonReq.Video))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockVideoUC.On("Insert", mock.Anything, mock.AnythingOfType("uint")).Return("", assert.AnError).Once()

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{assert.AnError.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.Insert(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos", strings.NewReader(jsonReq.VideoInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=29, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.Insert(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos", strings.NewReader(jsonReq.VideoInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"url: asd does not validate as url"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.Insert(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}

func TestUpdateVideoByID(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos/1", strings.NewReader(jsonReq.Video))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockVideoUC.On("UpdateByID", mock.AnythingOfType("uint"), mock.Anything, mock.AnythingOfType("uint")).
					Return("item edited", nil).Once()

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusOK
		resp.Meta.Message = "Success"
		resp.Data = "item edited"
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.UpdateByID(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos/1", strings.NewReader(jsonReq.Video))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)
		mockVideoUC.On("UpdateByID", mock.AnythingOfType("uint"), mock.Anything, mock.AnythingOfType("uint")).
					Return("", assert.AnError).Once()

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusInternalServerError
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{assert.AnError.Error()}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.UpdateByID(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Bind Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos/1", strings.NewReader(jsonReq.VideoInvalidBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"code=400, message=Syntax error: offset=29, error=invalid character '\"' after object key:value pair, internal=invalid character '\"' after object key:value pair"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.UpdateByID(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
	t.Run("Invalid Test | Invalid Struct", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/admins/videos/1", strings.NewReader(jsonReq.VideoInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		resp := controllers.BaseResponse{}
		resp.Meta.Status = http.StatusBadRequest
		resp.Meta.Message = "Error"
		resp.Meta.Messages = []string{"url: asd does not validate as url"}
		expected, _ := json.Marshal(resp)

		if assert.NoError(t, videoCtrl.UpdateByID(c)){
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}