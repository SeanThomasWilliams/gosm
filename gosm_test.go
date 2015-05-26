package gosm

import (
	"math"
	"testing"
)

func Test_GetTilesInBBoxForZoom(t *testing.T) {
	results, err := GetTilesInBBoxForZoom(-180, 85, 179.99, -85, 0)
	if err != nil {
		t.Errorf("Error getting tiles: %v", err)
	}
	for _, res := range results {
		t.Logf("Tile: %#v", res)
	}
}

func Test_GetAllTilesForZoomLevel(t *testing.T) {
	for i := 0; i < 10; i++ {
		results := GetAllTilesForZoomLevel(i)
		if len(results) != int(math.Pow(2, float64(2*i))) {
			t.Errorf("Wrong number of results for z: %d is %d", i, len(results))
		}
	}
}

func Test_Tile_GetBoundingBox(t *testing.T) {
	tile := GetAllTilesForZoomLevel(0)[0]
	t.Logf("Tile: %v", tile)
	ulx, uly, lrx, lry := tile.GetBoundingBox()
	t.Logf("BBOx: %f, %f, %f, %f", ulx, uly, lrx, lry)
}
