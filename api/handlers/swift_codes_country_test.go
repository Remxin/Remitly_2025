package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "example.com/m/v2/db/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestSwiftCodesCountry(t *testing.T) {
	swift := RandomSwift()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().
		AddNewSwiftCode(gomock.Any(), gomock.Eq(swift)).
		Times(1)

	server := SetupMockRouter(store)
	recorder := httptest.NewRecorder()

	url := "/v1/swift-codes"
	jsonData, err := json.Marshal(swift)
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
	require.NoError(t, err)
	server.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
}
