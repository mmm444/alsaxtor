package alsa

type ClientType int

const (
	CLIENT ClientType = 1
	KERNEL            = 2
)

type PortCapType uint

const (
	PORT_CAP_READ       PortCapType = (1 << 0)
	PORT_CAP_WRITE                  = (1 << 1)
	PORT_CAP_SYNC_READ              = (1 << 2)
	PORT_CAP_SYNC_WRITE             = (1 << 3)
	PORT_CAP_DUPLEX                 = (1 << 4)
	PORT_CAP_SUBS_READ              = (1 << 5)
	PORT_CAP_SUBS_WRITE             = (1 << 6)
	PORT_CAP_NO_EXPORT              = (1 << 7)
)

var portCapNames = map[PortCapType]string{
	PORT_CAP_READ:       "READ",
	PORT_CAP_WRITE:      "WRITE",
	PORT_CAP_SYNC_READ:  "SYNC_READ",
	PORT_CAP_SYNC_WRITE: "SYNC_WRITE",
	PORT_CAP_DUPLEX:     "DUPLEX",
	PORT_CAP_SUBS_READ:  "SUBS_READ",
	PORT_CAP_SUBS_WRITE: "SUBS_WRITE",
	PORT_CAP_NO_EXPORT:  "NO_EXPORT",
}

/** Sequencer event type */
type SeqEventType uint8

const (
	/** system status; event data type = #snd_seq_result_t */
	SND_SEQ_EVENT_SYSTEM SeqEventType = 0
	/** returned result status; event data type = #snd_seq_result_t */
	SND_SEQ_EVENT_RESULT = 1

	/** note on and off with duration; event data type = #snd_seq_ev_note_t */
	SND_SEQ_EVENT_NOTE = 5
	/** note on; event data type = #snd_seq_ev_note_t */
	SND_SEQ_EVENT_NOTEON = 6
	/** note off; event data type = #snd_seq_ev_note_t */
	SND_SEQ_EVENT_NOTEOFF = 7
	/** key pressure change (aftertouch); event data type = #snd_seq_ev_note_t */
	SND_SEQ_EVENT_KEYPRESS = 8

	/** controller; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_CONTROLLER = 10
	/** program change; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_PGMCHANGE = 11
	/** channel pressure; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_CHANPRESS = 12
	/** pitchwheel; event data type = #snd_seq_ev_ctrl_t; data is from -8192 to 8191) */
	SND_SEQ_EVENT_PITCHBEND = 13
	/** 14 bit controller value; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_CONTROL14 = 14
	/** 14 bit NRPN;  event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_NONREGPARAM = 15
	/** 14 bit RPN; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_REGPARAM = 16

	/** SPP with LSB and MSB values; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_SONGPOS = 20
	/** Song Select with song ID number; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_SONGSEL = 21
	/** midi time code quarter frame; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_QFRAME = 22
	/** SMF Time Signature event; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_TIMESIGN = 23
	/** SMF Key Signature event; event data type = #snd_seq_ev_ctrl_t */
	SND_SEQ_EVENT_KEYSIGN = 24

	/** MIDI Real Time Start message; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_START = 30
	/** MIDI Real Time Continue message; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_CONTINUE = 31
	/** MIDI Real Time Stop message; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_STOP = 32
	/** Set tick queue position; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_SETPOS_TICK = 33
	/** Set real-time queue position; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_SETPOS_TIME = 34
	/** (SMF) Tempo event; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_TEMPO = 35
	/** MIDI Real Time Clock message; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_CLOCK = 36
	/** MIDI Real Time Tick message; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_TICK = 37
	/** Queue timer skew; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_QUEUE_SKEW = 38
	/** Sync position changed; event data type = #snd_seq_ev_queue_control_t */
	SND_SEQ_EVENT_SYNC_POS = 39

	/** Tune request; event data type = none */
	SND_SEQ_EVENT_TUNE_REQUEST = 40
	/** Reset to power-on state; event data type = none */
	SND_SEQ_EVENT_RESET = 41
	/** Active sensing event; event data type = none */
	SND_SEQ_EVENT_SENSING = 42

	/** Echo-back event; event data type = any type */
	SND_SEQ_EVENT_ECHO = 50
	/** OSS emulation raw event; event data type = any type */
	SND_SEQ_EVENT_OSS = 51

	/** New client has connected; event data type = #snd_seq_addr_t */
	SND_SEQ_EVENT_CLIENT_START = 60
	/** Client has left the system; event data type = #snd_seq_addr_t */
	SND_SEQ_EVENT_CLIENT_EXIT = 61
	/** Client status/info has changed; event data type = #snd_seq_addr_t */
	SND_SEQ_EVENT_CLIENT_CHANGE = 62
	/** New port was created; event data type = #snd_seq_addr_t */
	SND_SEQ_EVENT_PORT_START = 63
	/** Port was deleted from system; event data type = #snd_seq_addr_t */
	SND_SEQ_EVENT_PORT_EXIT = 64
	/** Port status/info has changed; event data type = #snd_seq_addr_t */
	SND_SEQ_EVENT_PORT_CHANGE = 65

	/** Ports connected; event data type = #snd_seq_connect_t */
	SND_SEQ_EVENT_PORT_SUBSCRIBED = 66
	/** Ports disconnected; event data type = #snd_seq_connect_t */
	SND_SEQ_EVENT_PORT_UNSUBSCRIBED = 67

	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR0 = 90
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR1 = 91
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR2 = 92
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR3 = 93
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR4 = 94
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR5 = 95
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR6 = 96
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR7 = 97
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR8 = 98
	/** user-defined event; event data type = any (fixed size) */
	SND_SEQ_EVENT_USR9 = 99

	/** system exclusive data (variable length);  event data type = #snd_seq_ev_ext_t */
	SND_SEQ_EVENT_SYSEX = 130
	/** error event;  event data type = #snd_seq_ev_ext_t */
	SND_SEQ_EVENT_BOUNCE = 131
	/** reserved for user apps;  event data type = #snd_seq_ev_ext_t */
	SND_SEQ_EVENT_USR_VAR0 = 135
	/** reserved for user apps; event data type = #snd_seq_ev_ext_t */
	SND_SEQ_EVENT_USR_VAR1 = 136
	/** reserved for user apps; event data type = #snd_seq_ev_ext_t */
	SND_SEQ_EVENT_USR_VAR2 = 137
	/** reserved for user apps; event data type = #snd_seq_ev_ext_t */
	SND_SEQ_EVENT_USR_VAR3 = 138
	/** reserved for user apps; event data type = #snd_seq_ev_ext_t */
	SND_SEQ_EVENT_USR_VAR4 = 139

	/** NOP; ignored in any case */
	SND_SEQ_EVENT_NONE = 255
)

var sndSeqEventNames = map[SeqEventType]string{
	SND_SEQ_EVENT_SYSTEM:            "EVENT_SYSTEM",
	SND_SEQ_EVENT_RESULT:            "EVENT_RESULT",
	SND_SEQ_EVENT_NOTE:              "EVENT_NOTE",
	SND_SEQ_EVENT_NOTEON:            "EVENT_NOTEON",
	SND_SEQ_EVENT_NOTEOFF:           "EVENT_NOTEOFF",
	SND_SEQ_EVENT_KEYPRESS:          "EVENT_KEYPRESS",
	SND_SEQ_EVENT_CONTROLLER:        "EVENT_CONTROLLER",
	SND_SEQ_EVENT_PGMCHANGE:         "EVENT_PGMCHANGE",
	SND_SEQ_EVENT_CHANPRESS:         "EVENT_CHANPRESS",
	SND_SEQ_EVENT_PITCHBEND:         "EVENT_PITCHBEND",
	SND_SEQ_EVENT_CONTROL14:         "EVENT_CONTROL14",
	SND_SEQ_EVENT_NONREGPARAM:       "EVENT_NONREGPARAM",
	SND_SEQ_EVENT_REGPARAM:          "EVENT_REGPARAM",
	SND_SEQ_EVENT_SONGPOS:           "EVENT_SONGPOS",
	SND_SEQ_EVENT_SONGSEL:           "EVENT_SONGSEL",
	SND_SEQ_EVENT_QFRAME:            "EVENT_QFRAME",
	SND_SEQ_EVENT_TIMESIGN:          "EVENT_TIMESIGN",
	SND_SEQ_EVENT_KEYSIGN:           "EVENT_KEYSIGN",
	SND_SEQ_EVENT_START:             "EVENT_START",
	SND_SEQ_EVENT_CONTINUE:          "EVENT_CONTINUE",
	SND_SEQ_EVENT_STOP:              "EVENT_STOP",
	SND_SEQ_EVENT_SETPOS_TICK:       "EVENT_SETPOS_TICK",
	SND_SEQ_EVENT_SETPOS_TIME:       "EVENT_SETPOS_TIME",
	SND_SEQ_EVENT_TEMPO:             "EVENT_TEMPO",
	SND_SEQ_EVENT_CLOCK:             "EVENT_CLOCK",
	SND_SEQ_EVENT_TICK:              "EVENT_TICK",
	SND_SEQ_EVENT_QUEUE_SKEW:        "EVENT_QUEUE_SKEW",
	SND_SEQ_EVENT_SYNC_POS:          "EVENT_SYNC_POS",
	SND_SEQ_EVENT_TUNE_REQUEST:      "EVENT_TUNE_REQUEST",
	SND_SEQ_EVENT_RESET:             "EVENT_RESET",
	SND_SEQ_EVENT_SENSING:           "EVENT_SENSING",
	SND_SEQ_EVENT_ECHO:              "EVENT_ECHO",
	SND_SEQ_EVENT_OSS:               "EVENT_OSS",
	SND_SEQ_EVENT_CLIENT_START:      "EVENT_CLIENT_START",
	SND_SEQ_EVENT_CLIENT_EXIT:       "EVENT_CLIENT_EXIT",
	SND_SEQ_EVENT_CLIENT_CHANGE:     "EVENT_CLIENT_CHANGE",
	SND_SEQ_EVENT_PORT_START:        "EVENT_PORT_START",
	SND_SEQ_EVENT_PORT_EXIT:         "EVENT_PORT_EXIT",
	SND_SEQ_EVENT_PORT_CHANGE:       "EVENT_PORT_CHANGE",
	SND_SEQ_EVENT_PORT_SUBSCRIBED:   "EVENT_PORT_SUBSCRIBED",
	SND_SEQ_EVENT_PORT_UNSUBSCRIBED: "EVENT_PORT_UNSUBSCRIBED",
	SND_SEQ_EVENT_USR0:              "EVENT_USR0",
	SND_SEQ_EVENT_USR1:              "EVENT_USR1",
	SND_SEQ_EVENT_USR2:              "EVENT_USR2",
	SND_SEQ_EVENT_USR3:              "EVENT_USR3",
	SND_SEQ_EVENT_USR4:              "EVENT_USR4",
	SND_SEQ_EVENT_USR5:              "EVENT_USR5",
	SND_SEQ_EVENT_USR6:              "EVENT_USR6",
	SND_SEQ_EVENT_USR7:              "EVENT_USR7",
	SND_SEQ_EVENT_USR8:              "EVENT_USR8",
	SND_SEQ_EVENT_USR9:              "EVENT_USR9",
	SND_SEQ_EVENT_SYSEX:             "EVENT_SYSEX",
	SND_SEQ_EVENT_BOUNCE:            "EVENT_BOUNCE",
	SND_SEQ_EVENT_USR_VAR0:          "EVENT_USR_VAR0",
	SND_SEQ_EVENT_USR_VAR1:          "EVENT_USR_VAR1",
	SND_SEQ_EVENT_USR_VAR2:          "EVENT_USR_VAR2",
	SND_SEQ_EVENT_USR_VAR3:          "EVENT_USR_VAR3",
	SND_SEQ_EVENT_USR_VAR4:          "EVENT_USR_VAR4",
	SND_SEQ_EVENT_NONE:              "EVENT_NONE",
}

func (set SeqEventType) String() string {
	return sndSeqEventNames[set]
}
