package hello

import (
	"testing"
)

func TestNewGreeter(t *testing.T) {
	tests := []struct {
		name     string
		language string
		person   string
		want     *Greeter
	}{
		{
			name:     "Create English greeter",
			language: "en",
			person:   "John",
			want:     &Greeter{Language: "en", Name: "John"},
		},
		{
			name:     "Create Spanish greeter",
			language: "es",
			person:   "Juan",
			want:     &Greeter{Language: "es", Name: "Juan"},
		},
		{
			name:     "Create empty name greeter",
			language: "en",
			person:   "",
			want:     &Greeter{Language: "en", Name: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGreeter(tt.language, tt.person)
			if got.Language != tt.want.Language || got.Name != tt.want.Name {
				t.Errorf("NewGreeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreeter_Greet(t *testing.T) {
	tests := []struct {
		name     string
		language string
		person   string
		want     string
	}{
		{
			name:     "English greeting with name",
			language: "en",
			person:   "John",
			want:     "Hello, John![nhoJ]",
		},
		{
			name:     "Spanish greeting with name",
			language: "es",
			person:   "Juan",
			want:     "Hola, Juan![nauJ]",
		},
		{
			name:     "French greeting with name",
			language: "fr",
			person:   "Pierre",
			want:     "Bonjour, Pierre![erreiP]",
		},
		{
			name:     "German greeting with name",
			language: "de",
			person:   "Hans",
			want:     "Hallo, Hans![snaH]",
		},
		{
			name:     "Default greeting with name",
			language: "invalid",
			person:   "Someone",
			want:     "Hello, Someone![enoemoS]",
		},
		{
			name:     "English greeting without name",
			language: "en",
			person:   "",
			want:     "Hello, World!！",
		},
		{
			name:     "Spanish greeting without name",
			language: "es",
			person:   "",
			want:     "Hola, World!！",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGreeter(tt.language, tt.person)
			if got := g.Greet(); got != tt.want {
				t.Errorf("Greeter.Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
