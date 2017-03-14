package geth

/*
#include <stddef.h>
#include <stdbool.h>
extern bool StatusServiceSignalEvent(const char *jsonEvent);
*/
import "C"
import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
)

const (
	TestNodeSyncSeconds = 60
	TestNodeHTTPPort    = 8645
	TestNodeWSPort      = 8646
)

var (
	muPrepareTestNode sync.Mutex
	_, basePath, _, _ = runtime.Caller(0)
	RootDir           = filepath.Join(filepath.Dir(basePath), "..")
	DataDir           = filepath.Join(RootDir, "data")
	TestDataDir       = filepath.Join(RootDir, ".ethereumtest")
)

type NodeNotificationHandler func(jsonEvent string)

var notificationHandler NodeNotificationHandler = TriggerDefaultNodeNotificationHandler

// SetDefaultNodeNotificationHandler sets notification handler to invoke on SendSignal
func SetDefaultNodeNotificationHandler(fn NodeNotificationHandler) {
	notificationHandler = fn
}

// TriggerDefaultNodeNotificationHandler triggers default notification handler (helpful in tests)
func TriggerDefaultNodeNotificationHandler(jsonEvent string) {
	glog.V(logger.Info).Infof("notification received (default notification handler): %s\n", jsonEvent)
}

// SendSignal sends application signal (JSON, normally) upwards to application (via default notification handler)
func SendSignal(signal SignalEnvelope) {
	data, _ := json.Marshal(&signal)
	C.StatusServiceSignalEvent(C.CString(string(data)))
}

//export NotifyNode
func NotifyNode(jsonEvent *C.char) {
	notificationHandler(C.GoString(jsonEvent))
}

//export TriggerTestSignal
func TriggerTestSignal() {
	C.StatusServiceSignalEvent(C.CString(`{"answer": 42}`))
}

// TestConfig contains shared (among different test packages) parameters
type TestConfig struct {
	Node struct {
		SyncSeconds int
		HTTPPort    int
		WSPort      int
	}
	Account struct {
		Address  string
		Password string
	}
}

// LoadTestConfig loads test configuration values from disk
func LoadTestConfig() (*TestConfig, error) {
	var testConfig TestConfig

	configData := LoadFromFile(filepath.Join(DataDir, "test-data.json"))
	if err := json.Unmarshal([]byte(configData), &testConfig); err != nil {
		return nil, err
	}

	return &testConfig, nil
}

func CopyFile(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	if _, err := io.Copy(d, s); err != nil {
		return err
	}

	return nil
}

// LoadFromFile is usefull for loading test data, from testdata/filename into a variable
func LoadFromFile(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return ""
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)
	f.Close()

	return string(buf.Bytes())
}

func PrepareTestNode() (err error) {
	muPrepareTestNode.Lock()
	defer muPrepareTestNode.Unlock()

	manager := NodeManagerInstance()
	if manager.NodeInited() {
		return nil
	}

	defer HaltOnPanic()

	syncRequired := false
	if _, err := os.Stat(filepath.Join(TestDataDir, "testnet")); os.IsNotExist(err) {
		syncRequired = true
	}

	// prepare node directory
	if err := os.MkdirAll(filepath.Join(TestDataDir, "testnet", "keystore"), os.ModePerm); err != nil {
		glog.V(logger.Warn).Infoln("make node failed:", err)
		return err
	}

	// import test account (with test ether on it)
	dst := filepath.Join(TestDataDir, "testnet", "keystore", "test-account.pk")
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = CopyFile(dst, filepath.Join(RootDir, "data", "test-account.pk"))
		if err != nil {
			glog.V(logger.Warn).Infof("cannot copy test account PK: %v", err)
			return err
		}
	}

	// start geth node and wait for it to initialize
	err = CreateAndRunNode(&NodeConfig{
		DataDir:    TestDataDir,
		IPCEnabled: false,
		HTTPPort:   TestNodeHTTPPort, // to avoid conflicts with running app, using different port in tests
		WSEnabled:  false,
		WSPort:     TestNodeWSPort, // ditto
		TLSEnabled: false,
	})
	if err != nil {
		panic(err)
	}

	manager = NodeManagerInstance()
	if !manager.NodeInited() {
		panic(ErrInvalidGethNode)
	}
	if service, err := manager.RPCClient(); err != nil || service == nil {
		panic(ErrInvalidGethNode)
	}
	if service, err := manager.WhisperService(); err != nil || service == nil {
		panic(ErrInvalidGethNode)
	}
	if service, err := manager.LightEthereumService(); err != nil || service == nil {
		panic(ErrInvalidGethNode)
	}

	if syncRequired {
		glog.V(logger.Warn).Infof("Sync is required, it will take %d seconds", TestNodeSyncSeconds)
		time.Sleep(TestNodeSyncSeconds * time.Second) // LES syncs headers, so that we are up do date when it is done
	} else {
		time.Sleep(5 * time.Second)
	}

	return nil
}

func RemoveTestNode() {
	err := os.RemoveAll(TestDataDir)
	if err != nil {
		glog.V(logger.Warn).Infof("could not clean up temporary datadir")
	}
}

// PanicAfter throws panic() after waitSeconds, unless abort channel receives notification
func PanicAfter(waitSeconds time.Duration, abort chan struct{}, desc string) {
	go func() {
		select {
		case <-abort:
			return
		case <-time.After(waitSeconds):
			panic("whatever you were doing takes toooo long: " + desc)
		}
	}()
}

func FromAddress(accountAddress string) common.Address {
	from, err := ParseAccountString(accountAddress)
	if err != nil {
		return common.Address{}
	}

	return from.Address
}

func ToAddress(accountAddress string) *common.Address {
	to, err := ParseAccountString(accountAddress)
	if err != nil {
		return nil
	}

	return &to.Address
}

// ParseAccountString parses hex encoded string and returns is as accounts.Account.
func ParseAccountString(account string) (accounts.Account, error) {
	// valid address, convert to account
	if common.IsHexAddress(account) {
		return accounts.Account{Address: common.HexToAddress(account)}, nil
	}

	return accounts.Account{}, ErrInvalidAccountAddressOrKey
}
