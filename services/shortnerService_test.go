package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/satheeshds/shortly/mock"
)

func TestShortnerService_ShortURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockIShortnerRepository(ctrl)
	tests := []struct {
		name     string
		original string
		want     string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     "short valid",
			original: "https://test.com/example",
			want:     "",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortnerService{
				Repo: mockRepo,
			}
			mockRepo.EXPECT().GetPreviousShortenedIfExist(gomock.Any()).Return("", fmt.Errorf("No previous exist"))
			mockRepo.EXPECT().Store(gomock.Any(), gomock.Eq(tt.original))
			mockRepo.EXPECT().AddDomain(gomock.Any())
			_, err := s.ShortURL(tt.original)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortnerService.ShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestShortnerService_GetRedirectURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockIShortnerRepository(ctrl)
	tests := []struct {
		name    string
		short   string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "get redirect",
			short:   "1237",
			want:    "result",
			wantErr: false,
		},
		{
			name:    "get redirect",
			short:   "1237",
			want:    "result",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortnerService{
				Repo: mockRepo,
			}
			var err error
			if tt.wantErr {
				err = fmt.Errorf("Sample error")
			}
			mockRepo.EXPECT().Get(gomock.Eq(tt.short)).Return(tt.want, err)
			got, err := s.GetRedirectURL(tt.short)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortnerService.GetRedirectURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShortnerService.GetRedirectURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortnerService_GetTopShortedDomains(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockIShortnerRepository(ctrl)
	tests := []struct {
		name    string
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "no error",
			want:    map[string]int{"google": 1},
			wantErr: false,
		},
		{
			name:    "no error",
			want:    map[string]int{"google": 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortnerService{
				Repo: mockRepo,
			}
			var err error
			if tt.wantErr {
				err = fmt.Errorf("Sample error")
			}
			mockRepo.EXPECT().GetTopShortedDomains().Return(tt.want, err)
			got, err := s.GetTopShortedDomains()
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortnerService.GetTopShortedDomains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortnerService.GetTopShortedDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
