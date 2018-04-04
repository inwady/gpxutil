package context

import (
	"github.com/tkrajina/gpxgo/gpx"
	"errors"
	"log"
	"fmt"
)

type GPXContext struct {
	gpxData []*gpx.GPX
	workIndex uint
}

func InitFromGPXS(gpxs []*gpx.GPX) (*GPXContext, error) {
	if len(gpxs) == 0 {
		return nil, errors.New("bad gpxs")
	}

	return &GPXContext{
		gpxData: gpxs,
		workIndex: 0,
	}, nil
}

func (gctx *GPXContext) GetIndex() uint {
	return gctx.workIndex
}

func (gctx *GPXContext) SetWorkIndex(index uint) error {
	uLength := uint(len(gctx.gpxData))
	if index >= uLength {
		return errors.New("out of range")
	}

	gctx.workIndex = index
	return nil
}

func (gctx *GPXContext) AddGPX(gpx *gpx.GPX) uint {
	gctx.gpxData = append(gctx.gpxData, gpx)
	return uint(len(gctx.gpxData) - 1)
}

func (gctx *GPXContext) RemoveGPX(index uint) error {
	if index >= uint(len(gctx.gpxData)) {
		return fmt.Errorf("out of range")
	}

	gctx.gpxData = append(gctx.gpxData[:index], gctx.gpxData[index + 1:]...)
	return nil
}

func (gctx *GPXContext) AddPoint(lat float64, log float64) error {
	g := gctx.gpxData[gctx.workIndex]

	segment := getSegment(g)
	segment.Points = append(segment.Points, convertGPXPoint(lat, log))
	return nil
}

func (gctx *GPXContext) RemovePoint(index uint) error {
	g := gctx.gpxData[gctx.workIndex]

	segment := getSegment(g)
	if index >= uint(len(segment.Points)) {
		return errors.New("points, out of range")
	}

	segment.Points = append(segment.Points[:index], segment.Points[index + 1:]...)
	return nil
}

func (gctx *GPXContext) ChangePoint(index uint, lat float64, log float64) error {
	g := gctx.gpxData[gctx.workIndex]

	segment := getSegment(g)
	if index >= uint(len(segment.Points)) {
		return errors.New("points, out of range")
	}

	segment.Points[index].Latitude = lat
	segment.Points[index].Longitude = log
	return nil
}

func getSegment(gpx *gpx.GPX) *gpx.GPXTrackSegment {
	if len(gpx.Tracks) != 1 {
		log.Panicf("bad tracks") // TODO make tracks
		return nil
	}

	track := gpx.Tracks[0]
	if len(track.Segments) != 1 {
		log.Panicf("bad segments")
		return nil
	}

	return &track.Segments[0]
}

func (gctx *GPXContext) GetListInfo(index uint) (string, error) {
	uLength := uint(len(gctx.gpxData))
	if index >= uLength {
		return "", errors.New("out of range")
	}

	return gctx.gpxData[index].GetGpxInfo(), nil
}

func convertGPXPoint(lat float64, log float64) gpx.GPXPoint {
	return gpx.GPXPoint{
		Point: gpx.Point{
			Latitude: lat,
			Longitude: log,
		},
	}
}
