package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type TestData struct {
	token  string
	exp    int64
	secret string
	uid    string
}

var testDatas []TestData
var testData TestData

func init() {
	testDatas = []TestData{
		{
			uid:    "5ab4d3b5f5f6720005bcdb12",
			token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI1YWI0ZDNiNWY1ZjY3MjAwMDViY2RiMTIiLCJleHAiOjE1ODMwNDEzMzcsImlzcyI6IjVhYjRkM2I1ZjVmNjcyMDAwNWJjZGIxMiJ9.zo-KMsGFedmzAvkhNz1rskcwaWb03WfV4fmXbhv4VrY",
			exp:    1583041337,
			secret: "secret",
		},
		{
			uid:    "5ab4d3b5f5f6720005bcdb12",
			token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI1YWI0ZDNiNWY1ZjY3MjAwMDViY2RiMTIiLCJleHAiOjE1MjIwNDY5MzIsImlzcyI6IjVhYjRkM2I1ZjVmNjcyMDAwNWJjZGIxMiJ9.cdX6oR2_8MzNOOAFFJ3h2MUvEx4Cbrt4RKKYSGAqZf0",
			exp:    1522046932,
			secret: "secret",
		},
		{
			uid:    "5ab4d3b5f5f6720005bcdb12",
			token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI1YWI0YmVmMWM0Y2Q3NDhmMzJjNmRmZjMiLCJleHAiOjE1MjIwNTE1MjYsImlzcyI6IjVhYjRiZWYxYzRjZDc0OGYzMmM2ZGZmMyJ9._2p4YnbxsOzjYUCl7uNeS4GPhC5LCZGAELdiwNA_gX0",
			secret: "6fd64e1644629c16e7c0f994bcd493b5",
		},
	}
	testData = testDatas[0]
}

func Test_base64Encode(t *testing.T) {
	payload, err := json.Marshal(Claims{
		Uid: "5ab4d3b5f5f6720005bcdb12",
	})
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				src: payload,
			},
			want: "eyJ1aWQiOiI1YWI0ZDNiNWY1ZjY3MjAwMDViY2RiMTIifQ==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeB64(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base64Encode() = %v, \nwant %v", string(got), string(tt.want))
			}
		})
	}
}

func TestDecodeB64(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name       string
		args       args
		wantRetour string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				message: "eyJ1aWQiOiI1YWI0ZDNiNWY1ZjY3MjAwMDViY2RiMTIifQ==",
			},
			wantRetour: "5ab4d3b5f5f6720005bcdb12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRetour, err := DecodeB64(tt.args.message)
			if err != nil {
				t.Fatal(err)
			}
			ob := Claims{}
			e := json.Unmarshal(gotRetour, &ob)
			if e != nil {
				t.Fatal(e)
			}
			t.Log(ob.Uid)
			if ob.Uid != tt.wantRetour {
				t.Errorf("DecodeB64() = %v, want %v", ob, tt.wantRetour)
			}
		})
	}
}

func TestEncodeB642(t *testing.T) {
	str := "eyJ1aWQiOiI1YWI0ZDNiNWY1ZjY3MjAwMDViY2RiMTIifQ=="
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	ob := map[string]interface{}{}
	e := json.Unmarshal([]byte(data), &ob)
	fmt.Printf("%q e:%v ob:%v\n", data, e, ob["uid"])
}

func TestToMd5(t *testing.T) {
	type args struct {
		encode string
	}
	tests := []struct {
		name       string
		args       args
		wantDecode string
	}{
		// TODO: Add test cases.
		{
			name: "md5 test",
			args: args{
				encode: "5ab4bef1c4cd748f32c6dff3" + "123",
			},
			wantDecode: "FoNrJ3orKDBD4a31nr4WOg==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDecode := ToMd5(tt.args.encode); gotDecode != tt.wantDecode {
				t.Errorf("ToMd5() = %v, want %v", gotDecode, tt.wantDecode)
			}
		})
	}
}

func TestCreateToken(t *testing.T) {
	type args struct {
		uid    string
		secret string
		exp    int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				uid:    testData.uid,
				secret: testData.secret,
				exp:    testData.exp,
			},
			want: testData.token,
		},
		{
			name: "test2",
			args: args{
				uid:    testDatas[1].uid,
				secret: testDatas[1].secret,
				exp:    testDatas[1].exp,
			},
			want:    testDatas[1].token,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.uid, tt.args.secret, tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, \nwantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateToken() got = \n%v, \nwant %v", got, tt.want)
			}
		})
	}
}

func TestAuthToken(t *testing.T) {
	type args struct {
		signedToken string
		secret      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "authToken ",
			args: args{
				signedToken: testData.token,
				secret:      testData.secret,
			},
			want:    testData.uid,
			wantErr: false,
		},
		{
			name: "authToken err sercret",
			args: args{
				signedToken: testData.token,
				secret:      "123",
			},
			want:    "",
			wantErr: true,
		},
		{ // 过期
			name: "authToken err ",
			args: args{
				signedToken: testDatas[1].token,
				secret:      testDatas[1].secret,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "temp test.",
			args: args{
				signedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI1YWI0YmVmMWM0Y2Q3NDhmMzJjNmRmZjMiLCJleHAiOjE1MjIwNTE5NDYsImlzcyI6IjVhYjRiZWYxYzRjZDc0OGYzMmM2ZGZmMyJ9.T5IU2BRqS0jdILWox95mvc1tXdo5OfuozTGzvqgy29c",
				secret:      "1DGpRhV63WYA0Gk5vKErcg==",
			},
			want:    "5ab4bef1c4cd748f32c6dff3",
			wantErr: false,
		},
	}
	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AuthToken(tt.args.signedToken, tt.args.secret)
			t.Log("index:", index, "err:", err, "got:", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUid(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test uid",
			args: args{
				token: testData.token,
			},
			want:    testData.uid,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUid(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUid() = %v, want %v", got, tt.want)
			}
		})
	}
}
