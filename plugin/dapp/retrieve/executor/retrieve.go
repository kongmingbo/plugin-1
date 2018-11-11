package executor

import (
	"fmt"

	dbm "github.com/33cn/chain33/common/db"
	log "github.com/33cn/chain33/common/log/log15"
	drivers "github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	rt "github.com/33cn/plugin/plugin/dapp/retrieve/types"
)

var (
	minPeriod int64 = 60
	rlog            = log.New("module", "execs.retrieve")
)

var (
	zeroDelay       int64
	zeroPrepareTime int64
	zeroRemainTime  int64
)

var driverName = "retrieve"

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&Retrieve{}))
}

//const maxTimeWeight = 2
func Init(name string, sub []byte) {
	drivers.Register(GetName(), newRetrieve, types.GetDappFork(driverName, "Enable"))
}

func GetName() string {
	return newRetrieve().GetName()
}

type Retrieve struct {
	drivers.DriverBase
}

func newRetrieve() drivers.Driver {
	r := &Retrieve{}
	r.SetChild(r)
	r.SetExecutorType(types.LoadExecutorType(driverName))
	return r
}

func (r *Retrieve) GetDriverName() string {
	return driverName
}

func (r *Retrieve) CheckTx(tx *types.Transaction, index int) error {
	return nil
}

func calcRetrieveKey(backupAddr string, defaultAddr string) []byte {
	key := fmt.Sprintf("LODB-retrieve-backup:%s:%s", backupAddr, defaultAddr)
	return []byte(key)
}

func getRetrieveInfo(db dbm.KVDB, backupAddr string, defaultAddr string) (*rt.RetrieveQuery, error) {
	info := rt.RetrieveQuery{}
	retInfo, err := db.Get(calcRetrieveKey(backupAddr, defaultAddr))
	if err != nil {
		return nil, err
	}

	err = types.Decode(retInfo, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
