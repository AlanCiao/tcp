// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_freebsd.go

package tcp

const (
	sysTCPIOptTimestamps = 0x1
	sysTCPIOptSack       = 0x2
	sysTCPIOptWscale     = 0x4
	sysTCPIOptECN        = 0x8
	sysTCPIOptTOE        = 0x10

	sysSizeofTCPInfo = 0xec
)

type sysTCPInfo struct {
	State                  uint8
	X__tcpi_ca_state       uint8
	X__tcpi_retransmits    uint8
	X__tcpi_probes         uint8
	X__tcpi_backoff        uint8
	Options                uint8
	Pad_cgo_0              [2]byte
	Rto                    uint32
	X__tcpi_ato            uint32
	Snd_mss                uint32
	Rcv_mss                uint32
	X__tcpi_unacked        uint32
	X__tcpi_sacked         uint32
	X__tcpi_lost           uint32
	X__tcpi_retrans        uint32
	X__tcpi_fackets        uint32
	X__tcpi_last_data_sent uint32
	X__tcpi_last_ack_sent  uint32
	Last_data_recv         uint32
	X__tcpi_last_ack_recv  uint32
	X__tcpi_pmtu           uint32
	X__tcpi_rcv_ssthresh   uint32
	Rtt                    uint32
	Rttvar                 uint32
	Snd_ssthresh           uint32
	Snd_cwnd               uint32
	X__tcpi_advmss         uint32
	X__tcpi_reordering     uint32
	X__tcpi_rcv_rtt        uint32
	Rcv_space              uint32
	Snd_wnd                uint32
	Snd_bwnd               uint32
	Snd_nxt                uint32
	Rcv_nxt                uint32
	Toe_tid                uint32
	Snd_rexmitpack         uint32
	Rcv_ooopack            uint32
	Snd_zerowin            uint32
	X__tcpi_pad            [26]uint32
}
