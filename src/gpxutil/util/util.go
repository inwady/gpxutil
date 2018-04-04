package util

import (
	"github.com/twpayne/go-polyline"
	"github.com/tkrajina/gpxgo/gpx"
)

type HistoryCommand struct {
	canUndo bool
}

func ImportFromPolyline(name string, polylineBytes []byte) (*gpx.GPX, error) {
	coords, _, err := polyline.DecodeCoords(polylineBytes)
	if err != nil {
		return nil, err
	}

	gpxTrack := gpx.GPXTrack{}

	gpxSegment := gpx.GPXTrackSegment{}
	gpxSegment.Points = make([]gpx.GPXPoint, len(coords))
	for i, coord := range coords {
		gpxSegment.Points[i].Point = gpx.Point{
			Latitude: coord[0],
			Longitude: coord[1],
		}
	}

	gpxTrack.Segments = []gpx.GPXTrackSegment{gpxSegment}

	newGPX := &gpx.GPX{}
	newGPX.Name = name
	newGPX.Tracks = []gpx.GPXTrack{gpxTrack}
	return newGPX, nil
}
