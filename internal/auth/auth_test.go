package auth

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestGetAPIKey(t *testing.T){
    req1, _ := http.NewRequest(http.MethodGet, "/", nil)
    req2, _ := http.NewRequest(http.MethodGet, "/", nil)
    newUUID := uuid.New()
    stringUUID := newUUID.String()
    header1 := req1.Header
    header2 := req2.Header
    // Add the API Key
    header1.Add("Authorization", fmt.Sprintf("ApiKey %s", stringUUID))

    tests := []struct{
        name    string
        header  http.Header
        want    string
        wantErr bool
    }{
        {
            name: "Good Header",
            header: header1,
            want: stringUUID,
            wantErr: false,
        },
        {
            name: "Bad Header",
            header: header2,
            want: "no authorization header included",
            wantErr: true,
        },
    }
    for _, tt := range tests{
        t.Run(tt.name, func(t * testing.T){
            gotKey, err := GetAPIKey(tt.header)
            if (err != nil) != tt.wantErr{
                t.Errorf("GetAPIKey() error = %v , wantErr = %v", err, tt.wantErr)
                return
            }
            if err != nil && tt.wantErr{
                gotKey = fmt.Sprintf("%v", err)
            }
            if gotKey != tt.want {
                t.Errorf("GetAPIKey() key = %v , want = %v", gotKey, tt.want)
                return
            }
        })
    }
}
