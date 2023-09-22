package main

import (
	"fmt"

	"gonum.org/v1/hdf5"
)

type EventDataHDF5 struct {
	//	eventT0    bool
	//	eventT1    bool
	T0         uint32 `t0`
	T1         uint32 `t1`
	LostBuffer uint16 `lostBuffer`
	LostFPGA   uint16 `lostFGPA`
}

func writeData() {
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
