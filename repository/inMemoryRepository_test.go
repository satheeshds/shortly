package repository

import (
	"reflect"
	"testing"
)

func TestNewInMemoryRepository(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Generate New InMemory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemoryRepository(); got.isRepoInitialized() != nil {
				t.Errorf("NewInMemoryRepository() = %v is not initialized", got)
			}
		})
	}
}

func TestInMemoryRepository_isRepoInitialized(t *testing.T) {
	type fields struct {
		urlDictionary map[string]string
		domainTracker map[string]int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Url dictionary not created",
			fields: fields{
				urlDictionary: nil,
				domainTracker: make(map[string]int),
			},
			wantErr: true,
		},
		{
			name: "domain tracker not created",
			fields: fields{
				urlDictionary: make(map[string]string),
				domainTracker: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRepository{
				urlDictionary: tt.fields.urlDictionary,
				domainTracker: tt.fields.domainTracker,
			}
			if err := r.isRepoInitialized(); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.isRepoInitialized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryRepository_Store(t *testing.T) {
	type args struct {
		shortUrl string
		original string
	}
	r := NewInMemoryRepository()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "first store",
			args: args{
				shortUrl: "short.ly/test",
				original: "domain/long",
			},
			wantErr: false,
		},
		{
			name: "second store",
			args: args{
				shortUrl: "short.ly/test",
				original: "domain/long",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.Store(tt.args.shortUrl, tt.args.original); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryRepository_Get(t *testing.T) {
	r := NewInMemoryRepository()
	r.Store("short.ly/abc", "long.com/test")
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Key not exists",
			args: args{
				shortUrl: "short.ly/def",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Key exists",
			args: args{
				shortUrl: "short.ly/abc",
			},
			want:    "long.com/test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Get(tt.args.shortUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InMemoryRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryRepository_GetTopShortedDomains(t *testing.T) {
	r := NewInMemoryRepository()
	r.AddDomain("test")
	r.AddDomain("test")
	r.AddDomain("test1")
	r.AddDomain("test1")
	r.AddDomain("test1")
	r.AddDomain("test2")
	r.AddDomain("test2")
	r.AddDomain("test2")
	r.AddDomain("test2")
	r.AddDomain("test3")
	r.AddDomain("test4")
	r.AddDomain("test5")
	tests := []struct {
		name    string
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "sample",
			want: map[string]int{
				"test2": 4,
				"test1": 3,
				"test":  2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetTopShortedDomains()
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.GetTopShortedDomains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryRepository.GetTopShortedDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
