alsaxtor
========
_An automatic ALSA MIDI ports connector_

You give alsaxtor a list of ALSA port names pairs you wish to have connected.
When both ports of any pair become available they get connected.

alsaxtor can be invoked in 3 ways:

1. `alsaxtor -i` prints known clients and ports of ALSA MIDI sequencer (similar
   to `aconnect -iol`)
1. `alsaxtor -f file_of_port_pairs.csv` reads the to be connected port pairs from
   the given CSV file, connects all currently available ports and exits
1. `alsaxtor -r -f file_of_port_pairs.csv` same as 2. but it keeps running
   forever and any time a client is connected or disconnected it connects 
   available ports

alsaxtor was written for connecting multiple USB MIDI instruments via headless
Raspberry Pi to a USB to MIDI interface.

Note
----
It is not a good idea to wrap a unknown API in a language you are still learning.
There is a `cgo` wrapper for parts of the ALSA sequencer API in the `alsa` directory.

Have fun!