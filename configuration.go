package main

import (
	"fmt"
	"reflect"
)

type BitPosition struct {
	starts []int
	length int
}

func makeRange(start, end int) []int {
	a := make([]int, end-start+1)
	for i := range a {
		a[i] = start + i
	}
	return a
}

func makeRangeReverse(end, start int) []int {
	a := make([]int, end-start+1)
	for i := range a {
		a[i] = end - i
	}
	return a
}

func makeRangeReverseStep(end, start, step int) []int {
	a := make([]int, (end-start)/step+1)
	for i := range a {
		a[i] = end - step*i
	}
	return a
}

func makeRangeStep(start, end, step int) []int {
	a := make([]int, (end-start)/step+1)
	for i := range a {
		a[i] = start + step*i
	}
	return a
}

func updateByteArray(value int64, start, length int, bytearray []byte) {
	fmt.Printf("value: %v\n", value)
	for idx := 0; idx < length; idx++ {
		position := start + idx
		arrayPosition := position / 8
		mask := int64(1 << (length - idx - 1))
		bit := byte((value & mask) >> (length - idx - 1))
		//mask := int64(1 << idx)
		//bit := byte((value & mask) >> (idx))
		bytearray[arrayPosition] = bytearray[arrayPosition] | (bit << (7 - (position % 8)))
		fmt.Printf("position: %v, arrayPosition: %v, bit: %v, bytearray: %v\n", position, arrayPosition, bit, bytearray[arrayPosition])
	}
}

func configurationToByteArray(length int, configuration map[string]any, bitPositions map[string]BitPosition) []byte {
	bit_length := length
	bytearray := make([]byte, bit_length/8)

	for key, values := range configuration {
		fmt.Println(key, values)
		fmt.Println(key, bitPositions[key])

		v := reflect.ValueOf(values)
		fmt.Println(v.Kind())
		fmt.Println(v)

		switch v.Kind() {
		case reflect.Array:
			fmt.Println("hu-array!")
			length := bitPositions[key].length
			for i := 0; i < v.Len(); i++ {
				start := bitPositions[key].starts[i]
				updateByteArray(v.Index(i).Int(), start, length, bytearray)
			}
		case reflect.Int:
			fmt.Println("integer!")
			start := bitPositions[key].starts[0]
			length := bitPositions[key].length
			updateByteArray(v.Int(), start, length, bytearray)
		}
	}
	return bytearray
}

var citirocSlowControlBitPosition = map[string]BitPosition{
	"dac_t":                  {starts: makeRangeStep(0, 127, 4), length: 4},
	"dac":                    {starts: makeRangeStep(128, 255, 4), length: 4},
	"enableDiscriminator":    {starts: makeRange(256, 256), length: 1},
	"discriminator":          {starts: makeRange(257, 257), length: 1},
	"RS_or_discriminator":    {starts: makeRange(258, 258), length: 1},
	"enableDiscriminatorT":   {starts: makeRange(259, 259), length: 1},
	"discriminatorT":         {starts: makeRange(260, 260), length: 1},
	"enable_4b_dac":          {starts: makeRange(261, 261), length: 1},
	"dac4b":                  {starts: makeRange(262, 262), length: 1},
	"enable_4b_dac_t":        {starts: makeRange(263, 263), length: 1},
	"dac4b_t":                {starts: makeRange(264, 264), length: 1},
	"discriminatorMask":      {starts: makeRange(265, 296), length: 1},
	"HG_trackHold":           {starts: makeRange(297, 297), length: 1},
	"enable_HG_trackHold":    {starts: makeRange(298, 298), length: 1},
	"LH_trackHold":           {starts: makeRange(299, 299), length: 1},
	"enable_LG_trackHold":    {starts: makeRange(300, 300), length: 1},
	"scaBias":                {starts: makeRange(301, 301), length: 1},
	"HG_peakDetector":        {starts: makeRange(302, 302), length: 1},
	"enable_HG_peakDetector": {starts: makeRange(303, 303), length: 1},
	"LG_peakDetector":        {starts: makeRange(304, 304), length: 1},
	"enable_LG_peakDetector": {starts: makeRange(305, 305), length: 1},
	"sel_SCA_Pdet_HG":        {starts: makeRange(306, 306), length: 1},
	"sel_SCA_Pdet_LG":        {starts: makeRange(307, 307), length: 1},
	"bypass_PSC":             {starts: makeRange(308, 308), length: 1},
	"sel_trig_ext_PSC":       {starts: makeRange(309, 309), length: 1},
	"fastShaperFollower":     {starts: makeRange(310, 310), length: 1},
	"enableFastShaper":       {starts: makeRange(311, 311), length: 1},
	"fastShaper":             {starts: makeRange(312, 312), length: 1},
	"LG_slowShaper":          {starts: makeRange(313, 313), length: 1},
	"enable_LG_slowShaper":   {starts: makeRange(314, 314), length: 1},
	"timeConstant_LG_shaper": {starts: makeRange(315, 317), length: 3},
	"HG_slowShaper":          {starts: makeRange(318, 318), length: 1},
	"enable_HG_slowShaper":   {starts: makeRange(319, 319), length: 1},
	"timeConstant_HG_shaper": {starts: makeRange(320, 322), length: 3},
	"LG_preamp_bias":         {starts: makeRange(323, 323), length: 1},
	"HG_preamp":              {starts: makeRange(324, 324), length: 1},
	"enable_HG_praemp":       {starts: makeRange(325, 325), length: 1},
	"LG_preamp":              {starts: makeRange(326, 326), length: 1},
	"enable_LG_preamp":       {starts: makeRange(327, 327), length: 1},
	"fastShaperLG":           {starts: makeRange(328, 328), length: 1},
	"enable_input_DAC":       {starts: makeRange(329, 329), length: 1},
	"dac8_reference":         {starts: makeRange(330, 330), length: 1},
	"input_dac":              {starts: makeRangeStep(331, 618, 9), length: 8},
	"input_dac_on":           {starts: makeRangeStep(339, 618, 9), length: 1},
	"channel_preamp_HG":      {starts: makeRangeStep(619, 1098, 15), length: 6},
	"channel_preamp_LG":      {starts: makeRangeStep(625, 1098, 15), length: 6},
	"channel_Ctest_HG":       {starts: makeRangeStep(631, 1098, 15), length: 1},
	"channel_Ctest_LG":       {starts: makeRangeStep(632, 1098, 15), length: 1},
	"channel_preamp_disable": {starts: makeRangeStep(633, 1098, 15), length: 1},
	"temperature":            {starts: makeRange(1099, 1099), length: 1},
	"enableTemperature":      {starts: makeRange(1100, 1100), length: 1},
	"bandgap":                {starts: makeRange(1101, 1101), length: 1},
	"enable_bandgap":         {starts: makeRange(1102, 1102), length: 1},
	"enable_dac1":            {starts: makeRange(1103, 1103), length: 1},
	"dac1":                   {starts: makeRange(1104, 1104), length: 1},
	"enable_dac2":            {starts: makeRange(1105, 1105), length: 1},
	"dac2":                   {starts: makeRange(1106, 1106), length: 1},
	"dac1_code":              {starts: makeRange(1107, 1116), length: 10},
	"dac2_code":              {starts: makeRange(1117, 1126), length: 10},
	"enable_HG_OTA":          {starts: makeRange(1127, 1127), length: 1},
	"HG_OTA":                 {starts: makeRange(1128, 1128), length: 1},
	"enable_LG_OTA":          {starts: makeRange(1129, 1129), length: 1},
	"LG_OTA":                 {starts: makeRange(1130, 1130), length: 1},
	"enable_probe_OTA":       {starts: makeRange(1131, 1131), length: 1},
	"probe_OTA":              {starts: makeRange(1132, 1132), length: 1},
	"testb_OTA":              {starts: makeRange(1133, 1133), length: 1},
	"enable_val_evt":         {starts: makeRange(1134, 1134), length: 1},
	"val_evt":                {starts: makeRange(1135, 1135), length: 1},
	"enable_raz_chn":         {starts: makeRange(1136, 1136), length: 1},
	"raz_chn":                {starts: makeRange(1137, 1137), length: 1},
	"enable_out_dig":         {starts: makeRange(1138, 1138), length: 1},
	"enable_or32":            {starts: makeRange(1139, 1139), length: 1},
	"enable_nor32_oc":        {starts: makeRange(1140, 1140), length: 1},
	"trigger_polarity":       {starts: makeRange(1141, 1141), length: 1},
	"enable_bor32t_oc":       {starts: makeRange(1142, 1142), length: 1},
	"enable_32_triggers":     {starts: makeRange(1143, 1143), length: 1},
}

var slowControlRegister = map[string]any{
	"dac_t":                  [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"dac":                    [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"enableDiscriminator":    1,
	"discriminator":          1,
	"RS_or_discriminator":    0,
	"enableDiscriminatorT":   1,
	"discriminatorT":         1,
	"enable_4b_dac":          1,
	"dac4b":                  1,
	"enable_4b_dac_t":        1,
	"dac4b_t":                1,
	"discriminatorMask":      [32]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	"HG_trackHold":           1,
	"enable_HG_trackHold":    1,
	"LH_trackHold":           1,
	"enable_LG_trackHold":    1,
	"scaBias":                0,
	"HG_peakDetector":        0,
	"enable_HG_peakDetector": 0,
	"LG_peakDetector":        0,
	"enable_LG_peakDetector": 0,
	"sel_SCA_Pdet_HG":        1,
	"sel_SCA_Pdet_LG":        1,
	"bypass_PSC":             1,
	"sel_trig_ext_PSC":       0,
	"fastShaperFollower":     1,
	"enableFastShaper":       1,
	"fastShaper":             1,
	"LG_slowShaper":          1,
	"enable_LG_slowShaper":   1,
	"timeConstant_LG_shaper": 6,
	"HG_slowShaper":          1,
	"enable_HG_slowShaper":   1,
	"timeConstant_HG_shaper": 6,
	"LG_preamp_bias":         0,
	"HG_preamp":              1,
	"enable_HG_praemp":       1,
	"LG_preamp":              1,
	"enable_LG_preamp":       1,
	"fastShaperLG":           0,
	"enable_input_DAC":       1,
	"dac8_reference":         1,
	"input_dac":              [32]int{198, 203, 188, 202, 192, 203, 191, 202, 203, 203, 193, 197, 191, 201, 193, 205, 204, 190, 198, 195, 191, 193, 197, 192, 194, 185, 187, 192, 191, 195, 188, 198},
	"input_dac_on":           [32]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	"channel_preamp_HG":      [32]int{51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51},
	"channel_preamp_LG":      [32]int{47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47, 47},
	"channel_Ctest_HG":       [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"channel_Ctest_LG":       [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"channel_preamp_disable": [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"temperature":            1,
	"enableTemperature":      1,
	"bandgap":                1,
	"enable_bandgap":         1,
	"enable_dac1":            1,
	"dac1":                   1,
	"enable_dac2":            1,
	"dac2":                   1,
	"dac1_code":              250,
	"dac2_code":              250,
	"enable_HG_OTA":          1,
	"HG_OTA":                 1,
	"enable_LG_OTA":          1,
	"LG_OTA":                 1,
	"enable_probe_OTA":       1,
	"probe_OTA":              1,
	"testb_OTA":              0,
	"enable_val_evt":         0,
	"val_evt":                0,
	"enable_raz_chn":         0,
	"raz_chn":                0,
	"enable_out_dig":         0,
	"enable_or32":            0,
	"enable_nor32_oc":        0,
	"trigger_polarity":       0,
	"enable_bor32t_oc":       0,
	"enable_32_triggers":     1,
}

var citirocProbeBitPosition = map[string]BitPosition{
	"fastShaperOuput":    BitPosition{starts: makeRange(0, 31), length: 1},
	"slowShaperLGOutput": BitPosition{starts: makeRange(32, 64), length: 1},
	"peakSensingLG":      BitPosition{starts: makeRange(64, 95), length: 1},
	"slowShaperHGOutput": BitPosition{starts: makeRange(96, 127), length: 1},
	"peakSensingHG":      BitPosition{starts: makeRange(128, 159), length: 1},
	"preampHGOutput":     BitPosition{starts: makeRangeStep(160, 223, 2), length: 1},
	"preampLGOutput":     BitPosition{starts: makeRangeStep(161, 223, 2), length: 1},
	"inputDAC":           BitPosition{starts: makeRange(224, 255), length: 1},
}

var probeRegister = map[string]any{
	"fastShaperOuput":    [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"slowShaperLGOutput": [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"peakSensingLG":      [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"slowShaperHGOutput": [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"peakSensingHG":      [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	"preampHGOutput":     [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"preampLGOutput":     [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	"inputDAC":           [32]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}
