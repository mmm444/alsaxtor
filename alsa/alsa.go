package alsa

// #cgo pkg-config: alsa
// #include <alsa/asoundlib.h>
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// Represents the ALSA MIDI sequencer and current program's connection to it.
type Seq struct {
	seq     *C.snd_seq_t
	myId    C.int
	clients []Client
}

// Open a connection to the system MIDI sequencer. The client parameter is our name.
func OpenSeq(client string) (*Seq, error) {
	var seq *C.snd_seq_t

	def := C.CString("default")
	defer C.free(unsafe.Pointer(def))
	if C.snd_seq_open(&seq, def, C.SND_SEQ_OPEN_DUPLEX, 0) < 0 {
		return nil, errors.New("Cannot open sequencer.")
	}

	id := C.snd_seq_client_id(seq)
	if id < 0 {
		return nil, errors.New("Cannot obtain client id.")
	}

	cclient := C.CString(client)
	defer C.free(unsafe.Pointer(cclient))
	if C.snd_seq_set_client_name(seq, cclient) < 0 {
		C.snd_seq_close(seq)
		return nil, errors.New("Cannot set client name.")
	}

	s := &Seq{seq, id, nil}
	s.Refresh()

	return s, nil
}

// Re-read the information about clients and ports from the system.
func (s *Seq) Refresh() {
	clients := make([]Client, 0)

	var cinfo *C.snd_seq_client_info_t
	var pinfo *C.snd_seq_port_info_t

	C.snd_seq_client_info_malloc(&cinfo)
	defer C.snd_seq_client_info_free(cinfo)
	C.snd_seq_port_info_malloc(&pinfo)
	defer C.snd_seq_port_info_free(pinfo)

	C.snd_seq_client_info_set_client(cinfo, -1)
	for C.snd_seq_query_next_client(s.seq, cinfo) >= 0 {
		name := C.GoString(C.snd_seq_client_info_get_name(cinfo))
		id := C.snd_seq_client_info_get_client(cinfo)
		typ := ClientType(C.snd_seq_client_info_get_type(cinfo))
		client := Client{int(id), name, typ, make([]Port, 0), s}

		C.snd_seq_port_info_set_client(pinfo, id)
		C.snd_seq_port_info_set_port(pinfo, -1)
		for C.snd_seq_query_next_port(s.seq, pinfo) >= 0 {
			pname := C.GoString(C.snd_seq_port_info_get_name(pinfo))
			pid := C.snd_seq_port_info_get_port(pinfo)
			capa := C.snd_seq_port_info_get_capability(pinfo)
			port := Port{int(pid), pname, PortCapType(capa), &client}
			client.ports = append(client.ports, port)
		}

		clients = append(clients, client)
	}

	s.clients = clients
}

// Close a connection to the system MIDI sequencer.
func (s *Seq) Close() {
	C.snd_seq_close(s.seq)
}

// Return all clients registered with the system MIDI sequencer.
func (s *Seq) Clients() []Client {
	return s.clients
}

// Make a connection between 2 ports in the system MIDI sequencer.
func (s *Seq) Connect(src, dest *Port) error {
	var subs *C.snd_seq_port_subscribe_t
	C.snd_seq_port_subscribe_malloc(&subs)
	defer C.snd_seq_port_subscribe_free(subs)

	C.snd_seq_port_subscribe_set_sender(subs, src.addr())
	C.snd_seq_port_subscribe_set_dest(subs, dest.addr())
	//C.snd_seq_port_subscribe_set_queue(subs, queue)
	//C.snd_seq_port_subscribe_set_exclusive(subs, exclusive)
	//C.snd_seq_port_subscribe_set_time_update(subs, convert_time)
	//C.snd_seq_port_subscribe_set_time_real(subs, convert_real)

	if C.snd_seq_subscribe_port(s.seq, subs) < 0 {
		return errors.New("Connection failed.") // snd_strerror(errno)
	}

	s.Refresh()

	return nil
}

// Return the port with given client ID and port ID.
func (s *Seq) Port(cid, pid int) *Port {
	for _, c := range s.Clients() {
		if c.id == cid {
			for _, p := range c.Ports() {
				if p.id == pid {
					return &p
				}
			}
		}
	}
	return nil
}

// Create a port
func (s *Seq) CreatePort(name string, caps uint32) (*Port, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	res := C.snd_seq_create_simple_port(s.seq, cname, C.uint(caps), C.SND_SEQ_PORT_TYPE_MIDI_GENERIC)
	if res >= 0 {
		s.Refresh()
		return s.Port(int(s.myId), int(res)), nil
	} else {
		return nil, errors.New("Cannot create port.")
	}
}

// Run an event loop and send evets to given channel. Never returns.
func (s *Seq) EventLoop(ch chan<- SeqEventType) {
	for {
		var event *C.snd_seq_event_t
		C.snd_seq_event_input(s.seq, &event)
		ch <- SeqEventType(event._type)
	}
}

// Dump an information about current state of the ALSA sequencer to stdout.
func (seq *Seq) Dump() {
	for _, c := range seq.Clients() {
		fmt.Println(c)
		for _, p := range c.Ports() {
			fmt.Print(" ")
			fmt.Println(p)
		}
	}
}

// Represents the ALSA MIDI sequencer client.
type Client struct {
	id    int
	name  string
	type_ ClientType
	ports []Port
	seq   *Seq
}

// Return the ports that the client has.
func (c *Client) Ports() []Port {
	return c.ports
}

func (c Client) String() string {
	var t string
	switch c.type_ {
	case CLIENT:
		t = "client"
	case KERNEL:
		t = "kernel"
	default:
		t = "unknown"
	}
	return fmt.Sprintf("Client %d %s type=%s", c.id, c.name, t)
}

// Represents the ALSA MIDI sequencer port.
type Port struct {
	id     int
	name   string
	caps   PortCapType
	client *Client
}

// Is port capable to be a source in a connection?
func (p *Port) CanRead() bool {
	return p.caps&PORT_CAP_READ == PORT_CAP_READ && p.caps&PORT_CAP_SUBS_READ == PORT_CAP_SUBS_READ
}

// Is port capable to be a sink in a connection?
func (p *Port) CanWrite() bool {
	return p.caps&PORT_CAP_WRITE == PORT_CAP_WRITE && p.caps&PORT_CAP_SUBS_WRITE == PORT_CAP_SUBS_WRITE
}

// Return the name of the port.
func (p *Port) Name() string {
	return p.name
}

// Is the port connected as a source to some other port?
func (p *Port) HasConnOut() bool {
	return p.hasConn(C.SND_SEQ_QUERY_SUBS_READ)
}

// Has the port some incomming connection?
func (p *Port) HasConnIn() bool {
	return p.hasConn(C.SND_SEQ_QUERY_SUBS_WRITE)
}

func (p *Port) hasConn(dir C.snd_seq_query_subs_type_t) bool {
	var query *C.snd_seq_query_subscribe_t
	C.snd_seq_query_subscribe_malloc(&query)
	defer C.snd_seq_query_subscribe_free(query)

	C.snd_seq_query_subscribe_set_type(query, dir)
	C.snd_seq_query_subscribe_set_client(query, C.int(p.client.id))
	C.snd_seq_query_subscribe_set_port(query, C.int(p.id))
	C.snd_seq_query_port_subscribers(p.client.seq.seq, query)
	cnt := C.snd_seq_query_subscribe_get_num_subs(query)
	return cnt > 0

}

func (p Port) String() string {
	res := fmt.Sprintf("Port %d %s:", p.id, p.name)
	for m := PortCapType(1); m <= PORT_CAP_NO_EXPORT; m <<= 1 {
		if p.caps&m == m {
			res += " " + portCapNames[m]
		}
	}
	if p.HasConnIn() {
		res += " I"
	}
	if p.HasConnOut() {
		res += " O"
	}
	return res
}

func (p *Port) addr() *C.snd_seq_addr_t {
	var a C.snd_seq_addr_t
	a.client = C.uchar(p.client.id)
	a.port = C.uchar(p.id)
	return &a
}
