package pointcloud

import (
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

// MergePointCloudsWithColor creates a union of point clouds from the slice of point clouds, giving
// each element of the slice a unique color.
func MergePointCloudsWithColor(clusters []PointCloud) (PointCloud, error) {
	var err error
	palette := colorful.FastWarmPalette(len(clusters))
	colorSegmentation := New()
	for i, cluster := range clusters {
		col := color.NRGBAModel.Convert(palette[i])
		cluster.Iterate(func(pt Point) bool {
			v := pt.Position()
			colorPoint := NewColoredPoint(v.X, v.Y, v.Z, col.(color.NRGBA))
			err = colorSegmentation.Set(colorPoint)
			return err == nil
		})
		if err != nil {
			return nil, err
		}
	}
	return colorSegmentation, nil
}

// CalculateMeanOfPointCloud returns the spatial average center of a given point cloud.
func CalculateMeanOfPointCloud(cloud PointCloud) Vec3 {
	if cloud.Size() == 0 {
		return Vec3{}
	}
	x, y, z := 0.0, 0.0, 0.0
	n := float64(cloud.Size())
	cloud.Iterate(func(pt Point) bool {
		v := pt.Position()
		x += v.X
		y += v.Y
		z += v.Z
		return true
	})
	return Vec3{x / n, y / n, z / n}
}

// CalculateBoundingBoxOfPointCloud returns the dimensions of the bounding box
// formed by finding the dimensions of each axes' extrema.
func CalculateBoundingBoxOfPointCloud(cloud PointCloud) BoxGeometry {
	if cloud.Size() == 0 {
		return BoxGeometry{}
	}
	return BoxGeometry{
		WidthMm:  math.Abs(cloud.MaxX() - cloud.MinX()),
		LengthMm: math.Abs(cloud.MaxY() - cloud.MinY()),
		DepthMm:  math.Abs(cloud.MaxZ() - cloud.MinZ()),
	}
}

// PrunePointClouds removes point clouds from a slice if the point cloud has less than nMin points.
func PrunePointClouds(clouds []PointCloud, nMin int) []PointCloud {
	pruned := make([]PointCloud, 0, len(clouds))
	for _, cloud := range clouds {
		if cloud.Size() >= nMin {
			pruned = append(pruned, cloud)
		}
	}
	return pruned
}
