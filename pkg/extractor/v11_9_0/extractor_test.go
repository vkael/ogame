package v11_9_0

import (
	"github.com/alaingilbert/ogame/pkg/ogame"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExtractProduction1(t *testing.T) {
	pageHTMLBytes, _ := os.ReadFile("../../../samples/v11.9.0/en/shipyard.html")
	prods, secs, _ := NewExtractor().ExtractProduction(pageHTMLBytes)
	assert.Equal(t, 4, len(prods))
	assert.Equal(t, int64(402), secs) //6m40
	assert.Equal(t, ogame.LightFighterID, prods[0].ID)
	assert.Equal(t, int64(1), prods[0].Nbr)
	assert.Equal(t, ogame.SmallCargoID, prods[1].ID)
	assert.Equal(t, int64(1), prods[1].Nbr)
	assert.Equal(t, ogame.RocketLauncherID, prods[2].ID)
	assert.Equal(t, int64(1), prods[2].Nbr)
	assert.Equal(t, ogame.RocketLauncherID, prods[3].ID)
	assert.Equal(t, int64(1), prods[3].Nbr)
}

func TestExtractCombatReportMessages(t *testing.T) {
	pageHTMLBytes, _ := os.ReadFile("../../../samples/v11.11.5/en/combat_reports.html")
	msgs, _, _ := NewExtractor().ExtractCombatReportMessagesSummary(pageHTMLBytes)
	assert.Equal(t, 16, len(msgs))
}

func TestTraderAuctioneer(t *testing.T) {
	pageHTMLBytes, _ := os.ReadFile("../../../samples/v11.15.5/en/traderauctioneer.html")
	auction, _ := NewExtractor().ExtractAuction(pageHTMLBytes)
	assert.Equal(t, "c2bad58fcec374d709099d11d0549e59ea7e233e", auction.Ref)
	assert.Equal(t, "energy booster silver", auction.CurrentItem)
	assert.Equal(t, int64(7), auction.NumBids)
}
