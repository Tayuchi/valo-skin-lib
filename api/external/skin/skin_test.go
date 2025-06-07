package skin

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSkinDataList(t *testing.T) {
	t.Parallel()

	t.Run("データを取得し、マッピングできた場合", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			data, err := os.ReadFile("../../test_case/api_test.json")
			if err != nil {
				t.Fatal("failed to read file")
			}
			w.Header().Set("Content-Type", "application/json")
			if _, err := w.Write(data); err != nil {
				t.Fatal("failed to write response")
			}
		}))

		t.Cleanup(ts.Close)

		m := newModel(ts.URL)

		got, err := m.GetSkinDataList()
		if err != nil {
			t.Fatal("failed to run the method")
		}

		xenoHunterOdinVideo := "https://valorant.dyn.riotcdn.net/x/videos/release-10.10/32f7797f-4491-e21f-e40b-cfb639df3c97_default_universal.mp4"
		neptuneOdinVideo := "https://valorant.dyn.riotcdn.net/x/videos/release-10.10/b794b134-42d6-3138-188d-66a940a66304_default_universal.mp4"

		want := &skinDataList{
			Skins: []skinData{
				{
					Name:  "アルティチュード オーディン",
					Icon:  "https://media.valorant-api.com/weaponskins/89be9866-4807-6235-2a95-499cd23828df/displayicon.png",
					Video: nil,
				},
				{
					Name:  "ゼノハンター オーディン",
					Icon:  "https://media.valorant-api.com/weaponskins/94c085e6-48e1-c879-2552-88bf7850c5a8/displayicon.png",
					Video: &xenoHunterOdinVideo,
				},
				{
					Name:  "ネプチューン オーディン",
					Icon:  "https://media.valorant-api.com/weaponskins/a67c2daa-4f4d-1af0-0ff4-6fafde471776/displayicon.png",
					Video: &neptuneOdinVideo,
				},
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("response mismatch (-want +got):\n%s", diff)
		}
	})
}
