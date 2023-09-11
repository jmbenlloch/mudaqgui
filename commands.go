package main

type Cmd int

const DaqEtherType = 0x0801

const (
	// Ethernet switch commands and responses
	FEB_RD_SR      Cmd = 0x0001
	FEB_WR_SR          = 0x0002
	FEB_RD_SR_SRFF     = 0x0003
	FEB_WR_SR_SRFF     = 0x0004
	FEB_OK_SR          = 0x0000
	FEB_ERR_SR         = 0x00FF

	// Board configuration and data acquisition
	FEB_SET_RECV  Cmd = 0x0101
	FEB_GEN_INIT      = 0x0102
	FEB_GEN_HVON      = 0x0103
	FEB_GEN_HVOFF     = 0x0104
	FEB_GET_RATE      = 0x0105
	FEB_OK            = 0x0100
	FEB_ERR           = 0x01FF

	// CITIROC slow control register
	FEB_RD_SCR  Cmd = 0x0201
	FEB_WR_SCR      = 0x0202
	FEB_OK_SCR      = 0x0200
	FEB_ERR_SCR     = 0x02FF

	// CITIROC data control register
	FEB_RD_CDR   Cmd = 0x0301
	FEB_DATA_CDR     = 0x0300
	FEB_EOF_CDR      = 0x0303
	FEB_ERR_CDR      = 0x03FF

	// CITIROC probe register
	FEB_RD_PMR  Cmd = 0x0401
	FEB_WR_PMR      = 0x0402
	FEB_OK_PMR      = 0x0400
	FEB_ERR_PMR     = 0x04FF

	// Firmware read-write transmission
	FEB_RD_FW   Cmd = 0x0501
	FEB_WR_FW       = 0x0502
	FEB_OK_FW       = 0x0500
	FEB_ERR_FW      = 0x05FF
	FEB_EOF_FW      = 0x0503
	FEB_DATA_FW     = 0x0504

	// FPGA input logic configuration register
	FEB_RD_FIL  Cmd = 0x0601
	FEB_WR_FIL      = 0x0602
	FEB_OK_FIL      = 0x0600
	FEB_ERR_FIL     = 0x06FF
)
