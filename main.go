package main

import (
	"encoding/csv"
	"flag"
	"github.com/mmm444/alsaxtor/alsa"
	"io"
	"log"
	"os"
)

type Connector struct {
	seq           *alsa.Seq
	possiblePairs [][]string
}

func NewConnector(seq *alsa.Seq) *Connector {
	return &Connector{seq, make([][]string, 0)}
}

func (c *Connector) AddPair(from, to string) {
	c.possiblePairs = append(c.possiblePairs, []string{from, to})
}

func (c *Connector) ReadPairsCsv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.FieldsPerRecord = 2

	for {
		pair, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		c.AddPair(pair[0], pair[1])
	}

	return nil
}

func (c *Connector) ConnectPossible() {
	// wheee
	for _, c1 := range c.seq.Clients() {
		for _, p1 := range c1.Ports() {
			if !p1.HasConnOut() {
				for _, poss := range c.possiblePairs {
					if p1.Name() == poss[0] {
						for _, c2 := range c.seq.Clients() {
							for _, p2 := range c2.Ports() {
								if p2.Name() == poss[1] && !p2.HasConnIn() {
									if c.seq.Connect(&p1, &p2) == nil {
										log.Printf("Connected %v to %v", p1.Name(), p2.Name())
									} else {
										log.Printf("Cannot connect %v to %v", p1.Name(), p2.Name())
									}

								}
							}
						}
					}
				}
			}
		}
	}
}

var (
	info        = flag.Bool("i", false, "print ALSA sequencer state")
	csvFile     = flag.String("f", "", "read to be connected port pairs from CSV")
	keepRunning = flag.Bool("r", false, "run and connect ports forever")
	verbose     = flag.Bool("v", false, "be verbose (log received events)")
)

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	seq, err := alsa.OpenSeq("alsaxtor")
	if err != nil {
		log.Fatal(err)
	}
	defer seq.Close()

	port, err := seq.CreatePort("Listener", alsa.PORT_CAP_WRITE|alsa.PORT_CAP_SUBS_WRITE)
	if err != nil {
		log.Fatal(err)
	}

	err = seq.Connect(seq.Port(0, 1), port)
	if err != nil {
		log.Fatal(err)
	}

	if *info {
		seq.Dump()
	}

	if *csvFile != "" {
		connector := NewConnector(seq)
		//connector.AddPair("UM-2G MIDI 1", "UM-2G MIDI 2")
		err = connector.ReadPairsCsv(*csvFile)
		if err != nil {
			log.Fatal("Cannot read csv file: ", err)
		}
		connector.ConnectPossible()

		if *keepRunning {
			ch := make(chan alsa.SeqEventType)
			go seq.EventLoop(ch)

			for t := range ch {
				if *verbose {
					log.Println("Received", t)
				}
				if t == alsa.SND_SEQ_EVENT_PORT_START || t == alsa.SND_SEQ_EVENT_CLIENT_EXIT {
					seq.Refresh()
					connector.ConnectPossible()
				}
			}
		}
	} else {
		log.Fatal("No to be connected port pairs specified.")
	}
}
