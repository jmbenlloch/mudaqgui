package main

import (
	"fmt"

	"gonum.org/v1/hdf5"
)

type WriterData struct {
	file    *hdf5.File
	data    *hdf5.Dataset
	charges *hdf5.Dataset
}

type EventDataHDF5 struct {
	card       byte
	eventT0    bool
	eventT1    bool
	overflowT0 bool
	overflowT1 bool
	T0         uint32 `t0`
	T1         uint32 `t1`
	LostBuffer uint16 `lostBuffer`
	LostFPGA   uint16 `lostFGPA`
}

func writeDataTest() {
	const length = 10
	s1 := [length]EventDataHDF5{}
	for i := 0; i < int(length); i++ {
		s1[i] = EventDataHDF5{
			T0:         123,
			T1:         345,
			LostBuffer: 346,
			LostFPGA:   326,
			//Charges:    [32]uint16{},
		}
	}
	fmt.Printf(":: data: %v\n", s1)

	// create data space
	dims := []uint{length}
	space, err := hdf5.CreateSimpleDataspace(dims, nil)
	if err != nil {
		panic(err)
	}

	// create the file
	fname := "test.h5"
	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf(":: file [%s] created (id=%d)\n", fname, f.ID())

	// create the memory data type
	dtype, err := hdf5.NewDatatypeFromValue(s1[0])
	if err != nil {
		panic("could not create a dtype")
	}

	// create the dataset
	dsname := "data"
	dset, err := f.CreateDataset(dsname, dtype, space)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: dset (id=%d)\n", dset.ID())

	// write data to the dataset
	fmt.Printf(":: dset.Write...\n")
	err = dset.Write(&s1)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: dset.Write... [ok]\n")

	// write data to the dataset
	fmt.Printf(":: dset.Write...\n")
	err = dset.Write(&s1)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: dset.Write... [ok]\n")

	// write data to the dataset
	fmt.Printf(":: dset.Write...\n")
	err = dset.Write(&s1)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: dset.Write... [ok]\n")

	// release resources
	dset.Close()
	f.Close()
}

func openFile(fname string) *hdf5.File {
	// create the file
	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		panic(err)
	}
	//defer f.Close()
	fmt.Printf(":: file [%s] created (id=%d)\n", fname, f.ID())
	return f
}

func createTable(file *hdf5.File) *hdf5.Dataset {
	dims := []uint{0}
	unlimitedDims := -1 // H5S_UNLIMITED is -1L
	maxDims := []uint{uint(unlimitedDims)}
	file_space, err := hdf5.CreateSimpleDataspace(dims, maxDims)
	if err != nil {
		fmt.Println("space")
		panic(err)
	}
	fmt.Println(file_space)

	// create property list
	plist, err := hdf5.NewPropList(hdf5.P_DATASET_CREATE)
	if err != nil {
		fmt.Println("plist")
		panic(err)
	}
	chunks := []uint{32768}
	plist.SetChunk(chunks)
	// Set compression level
	plist.SetDeflate(4)

	// create the memory data type
	dtype, err := hdf5.NewDatatypeFromValue(EventDataHDF5{})
	if err != nil {
		fmt.Println("datatype")
		panic("could not create a dtype")
	}

	// create the dataset
	dsname := "data"
	dset, err := file.CreateDatasetWith(dsname, dtype, file_space, plist)
	if err != nil {
		fmt.Println("dataset")
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf(":: dset (id=%d)\n", dset.ID())
	return dset
}

func writeData(dataset *hdf5.Dataset, events *[]EventData) {
	length := uint(len(*events))
	s2 := make([]EventDataHDF5, length)
	for i := 0; i < int(length); i++ {
		s2[i] = EventDataHDF5{
			card:       (*events)[i].card,
			eventT0:    (*events)[i].eventT0,
			eventT1:    (*events)[i].eventT1,
			overflowT0: (*events)[i].overflowT0,
			overflowT1: (*events)[i].overflowT1,
			T0:         (*events)[i].T0,
			T1:         (*events)[i].T1,
			LostBuffer: (*events)[i].LostBuffer,
			LostFPGA:   (*events)[i].LostFPGA,
		}
	}
	fmt.Printf(":: data: %v\n", s2)

	dims := []uint{length}
	dataspace, err := hdf5.CreateSimpleDataspace(dims, nil)
	if err != nil {
		fmt.Println("space")
		panic(err)
	}

	// extend
	dimsGot, maxdimsGot, err := dataset.Space().SimpleExtentDims()
	eventsInFile := dimsGot[0]
	fmt.Println("Size: ", dimsGot, maxdimsGot)
	newsize := []uint{eventsInFile + length}
	dataset.Resize(newsize)
	filespace := dataset.Space()
	fmt.Println(filespace)

	start := []uint{eventsInFile}
	count := []uint{length}
	filespace.SelectHyperslab(start, nil, count, nil)

	// write data to the dataset
	fmt.Printf(":: dset.Write...\n")
	//err = dset.Write(&s2)
	err = dataset.WriteSubset(&s2, dataspace, filespace)
	if err != nil {
		fmt.Println("final write")
		panic(err)
	}
	fmt.Printf(":: dset.Write... [ok]\n")
}
