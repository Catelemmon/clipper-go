package clipper


type IntPoint struct {
	cIntPointPtr *cgo_IntPoint_T
}

func NewIntPoint(x int64, y int64) *IntPoint {
	return &IntPoint{
		cIntPointPtr: cgo_NewIntPoint_T(x, y),
	}
}

func (point *IntPoint) Delete(){
	cgo_DeleteIntPoint_T(point.cIntPointPtr)
}

type Polygon struct {
	cPolygonPtr *cgo_Polygon_T
}

func NewPolygon(intPointPtrList []*IntPoint) *Polygon {
	cInPointPtrList := make([]*cgo_IntPoint_T,0, len(intPointPtrList))
	for i := 0; i < len(intPointPtrList); i++ {
		cInPointPtrList = append(cInPointPtrList, intPointPtrList[i].cIntPointPtr)
	}
	return &Polygon{
		cPolygonPtr: cgo_NewPolygon_T(cInPointPtrList),
	}
}

func NewEmptyPolygon() *Polygon {
	return &Polygon{
		cPolygonPtr: cgo_NewEmptyPolygon_T(),
	}
}

func (polygon *Polygon) Delete() {
	cgo_DeletePolygon_T(polygon.cPolygonPtr)
}

// AppendIntPoint add point to polygon
func (polygon *Polygon) AppendIntPoint(point *IntPoint){
	cgo_AppendIntPoint(polygon.cPolygonPtr, point.cIntPointPtr)
}

// EdgeNum return edge number(type: unsigned int32) of polygon
func (polygon *Polygon) EdgeNum() uint32{
	return cgo_EdgeNum(polygon.cPolygonPtr)
}

// Area return area(type: float64) of polygon
func (polygon *Polygon) Area() float64{
	return cgo_PolygonArea(polygon.cPolygonPtr)
}

// IntersectionRate intersection of two polygons, (two polygon intersection area) / (polygon2 / area)
func (polygon *Polygon) IntersectionRate(polygon1 *Polygon) float64{
	return cgo_IntersectionRate(polygon.cPolygonPtr, polygon1.cPolygonPtr)
}

