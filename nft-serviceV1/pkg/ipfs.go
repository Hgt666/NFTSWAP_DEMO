package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"nft-service/config"
)

type NFTMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// FetchIPFSMeta 从IPFS网关拉取元数据
func FetchIPFSMeta(cid string) (*NFTMetadata, error) {
	url := config.IPFSGateway + cid
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ipfs fetch failed: %d", resp.StatusCode)
	}

	var meta NFTMetadata
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return nil, err
	}
	return &meta, nil
}