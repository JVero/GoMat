package matrix

import ("testing")


func BenchmarkMult(b *testing.B){
	data := LoadCSV("data/sampledata")
	println(data.numRows, data.numCols)
	//var bd Matrix
	for i:=0; i<b.N; i++{
		_ = data.multiply(data)
	}

}
