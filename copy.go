package obj

func Copy (wfs , wfd , wft *Wavefront){
	for _,wfd_points := range wfd.CoordinatePoints{
		cur_len := len(wft.CoordinatePoints)
		for _,wfs_points:= range wfs.CoordinatePoints{
			pt := []float64{wfd_points[0]+wfs_points[0],wfd_points[1]+wfs_points[1],wfd_points[2]+wfs_points[2]}
			wft.CoordinatePoints = append(wft.CoordinatePoints,pt)
		}
		fmt.Printf("%v\n",cur_len)
		for _,wfs_faces := range wfs.Faces{
			fc := [][]int{}
			for _,t := range wfs_faces{
				fc = append(fc ,[]int{t[0]+cur_len,t[1]+cur_len})
			}
			wft.Faces = append(wft.Faces , fc)
		}
	}
}

