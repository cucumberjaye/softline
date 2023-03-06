package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	mock_handler "github.com/cucumberjaye/softline/internal/app/service/mocks"
	"github.com/cucumberjaye/softline/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {

	tests := []struct {
		name     string
		body     io.Reader
		user     models.RegisterUser
		code     int
		response string
		id       int
		err      error
	}{
		{
			name: "up_ok",
			code: http.StatusOK,
			user: models.RegisterUser{
				Login:                "lol",
				Email:                "lol@gmail.com",
				Password:             "123456",
				PasswordConfirmation: "123456",
				PhoneNumber:          "+12234566",
			},
			response: "{\"id\":1}\n",
			id:       1,
			err:      nil,
		},
		{
			name: "up_valid_err",
			code: http.StatusBadRequest,
			user: models.RegisterUser{
				Login:                "",
				Email:                "lol@gmail.com",
				Password:             "123456",
				PasswordConfirmation: "123456",
				PhoneNumber:          "+12234566",
			},
			response: "invalid request body: Key: 'RegisterUser.Login' Error:Field validation for 'Login' failed on the 'required' tag\n",
			id:       0,
			err:      errors.New(""),
		},
		{
			name: "up_service_err",
			code: http.StatusInternalServerError,
			user: models.RegisterUser{
				Login:                "lol",
				Email:                "lol@gmail.com",
				Password:             "123456",
				PasswordConfirmation: "123456",
				PhoneNumber:          "+12234566",
			},
			response: "test\n",
			id:       0,
			err:      errors.New("test"),
		},
	}

	ctrl := gomock.NewController(t)
	as := mock_handler.NewMockAuthService(ctrl)
	h := &Handler{authService: as}

	r := h.InitRoutes()
	ts := httptest.NewServer(r)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := json.Marshal(tt.user)
			require.NoError(t, err)

			tt.body = bytes.NewBuffer(user)
			request := httptest.NewRequest(http.MethodPost, ts.URL+"/auth/sign-up", tt.body)
			request.RequestURI = ""

			if tt.name != "up_valid_err" {
				as.EXPECT().CreateUser(tt.user).Return(tt.id, tt.err)
			}
			resp, err := http.DefaultClient.Do(request)
			require.NoError(t, err)

			require.Equal(t, resp.StatusCode, tt.code)
			defer resp.Body.Close()

			resBody, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, tt.response, string(resBody))

		})
	}
}
