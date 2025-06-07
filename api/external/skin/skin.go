package skin

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type apiResponse struct {
	Data []apiSkin `json:"data"`
}

type apiSkin struct {
	Name   string     `json:"displayName"`
	Icon   string     `json:"displayIcon"`
	Levels []apiLevel `json:"levels"`
}

type apiLevel struct {
	StreamedVideo *string `json:"streamedVideo"`
}

type skinData struct {
	Name  string
	Icon  string
	Video *string
}

type skinDataList struct {
	Skins []skinData
}

type Model struct {
	SkinListURL string
}

func newModel(skinListURL string) Model {
	return Model{
		SkinListURL: skinListURL,
	}
}

func (m *Model) GetSkinDataList() (*skinDataList, error) {
	// リクエスト
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, m.SkinListURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var apiRes apiResponse
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return nil, err
	}

	// データをマッピング
	var skins []skinData
	for _, s := range apiRes.Data {
		var video *string
		for _, lv := range s.Levels {
			if lv.StreamedVideo != nil {
				video = lv.StreamedVideo
				break
			}
		}

		skins = append(skins, skinData{
			Name:  s.Name,
			Icon:  s.Icon,
			Video: video,
		})
	}

	return &skinDataList{
		Skins: skins,
	}, nil
}
