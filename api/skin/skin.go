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
	Name  string  `json:"name"`
	Icon  string  `json:"icon"`
	Video *string `json:"video"`
}

type skinDataList struct {
	Skins []skinData `json:"skins"`
}

type SkinService struct {
	SkinListURL string
	cache       *skinDataList
	cacheErr    error
}

func NewSkinService(skinListURL string) SkinService {
	s := SkinService{
		SkinListURL: skinListURL,
	}

	s.refreshCache()
	go func() {
		ticker := time.NewTicker(time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			s.refreshCache()
		}
	}()

	return s
}

func (s *SkinService) refreshCache() {
	data, err := s.GetSkinDataListFromAPI()
	s.cache = data
	s.cacheErr = err
}

func (s *SkinService) GetSkinDataList() (*skinDataList, error) {
	return s.cache, s.cacheErr
}

func (m *SkinService) GetSkinDataListFromAPI() (*skinDataList, error) {
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
