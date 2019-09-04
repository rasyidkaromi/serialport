
package serialport

type Port interface {
	SetMode(mode *Mode) error
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	ResetInputBuffer() error
	ResetOutputBuffer() error
	SetDTR(dtr bool) error
	SetRTS(rts bool) error
	GetModemStatusBits() (*ModemStatusBits, error)
	Close() error
}

type ModemStatusBits struct {
	CTS bool // ClearToSend status
	DSR bool // DataSetReady status
	RI  bool // RingIndicator status
	DCD bool // DataCarrierDetect status
}

func Open(portName string, mode *Mode) (Port, error) {
	return nativeOpen(portName, mode)
}

func GetPortsList() ([]string, error) {
	return nativeGetPortsList()
}

type Mode struct {
	BaudRate int      // The serial port bitrate (aka Baudrate)
	DataBits int      // Size of the character (must be 5, 6, 7 or 8)
	Parity   Parity   // Parity (see Parity type for more info)
	StopBits StopBits // Stop bits (see StopBits type for more info)
}

type Parity int

const (
	NoParity Parity = iota
	OddParity
	EvenParity
	MarkParity
	SpaceParity
)

type StopBits int

const (
	OneStopBit StopBits = iota
	OnePointFiveStopBits
	TwoStopBits
)

type PortError struct {
	code     PortErrorCode
	causedBy error
}

type PortErrorCode int

const (
	PortBusy PortErrorCode = iota
	PortNotFound
	InvalidSerialPort
	PermissionDenied
	InvalidSpeed
	InvalidDataBits
	InvalidParity
	InvalidStopBits
	ErrorEnumeratingPorts
	PortClosed
	FunctionNotImplemented
)

func (e PortError) EncodedErrorString() string {
	switch e.code {
	case PortBusy:
		return "Serial port busy"
	case PortNotFound:
		return "Serial port not found"
	case InvalidSerialPort:
		return "Invalid serial port"
	case PermissionDenied:
		return "Permission denied"
	case InvalidSpeed:
		return "Port speed invalid or not supported"
	case InvalidDataBits:
		return "Port data bits invalid or not supported"
	case InvalidParity:
		return "Port parity invalid or not supported"
	case InvalidStopBits:
		return "Port stop bits invalid or not supported"
	case ErrorEnumeratingPorts:
		return "Could not enumerate serial ports"
	case PortClosed:
		return "Port has been closed"
	case FunctionNotImplemented:
		return "Function not implemented"
	default:
		return "Other error"
	}
}

func (e PortError) Error() string {
	if e.causedBy != nil {
		return e.EncodedErrorString() + ": " + e.causedBy.Error()
	}
	return e.EncodedErrorString()
}

func (e PortError) Code() PortErrorCode {
	return e.code
}
