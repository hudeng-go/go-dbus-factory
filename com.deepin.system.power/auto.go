// Code generated by "./generator com.deepin.system.power"; DO NOT EDIT.

package power

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Power interface {
	power // interface com.deepin.system.Power
	proxy.Object
}

type objectPower struct {
	interfacePower // interface com.deepin.system.Power
	proxy.ImplObject
}

func NewPower(conn *dbus.Conn) Power {
	obj := new(objectPower)
	obj.ImplObject.Init_(conn, "com.deepin.system.Power", "/com/deepin/system/Power")
	return obj
}

type power interface {
	GoGetBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetBatteries(flags dbus.Flags) ([]dbus.ObjectPath, error)
	GoRefresh(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Refresh(flags dbus.Flags) error
	GoRefreshBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RefreshBatteries(flags dbus.Flags) error
	GoRefreshMains(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RefreshMains(flags dbus.Flags) error
	GoSetCpuGovernor(flags dbus.Flags, ch chan *dbus.Call, governor string) *dbus.Call
	SetCpuGovernor(flags dbus.Flags, governor string) error
	GoSetCpuBoost(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	SetCpuBoost(flags dbus.Flags, enabled bool) error
	GoSetMode(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call
	SetMode(flags dbus.Flags, mode string) error
	GoLockCpuFreq(flags dbus.Flags, ch chan *dbus.Call, governor string, lockTime int32) *dbus.Call
	LockCpuFreq(flags dbus.Flags, governor string, lockTime int32) error
	ConnectBatteryDisplayUpdate(cb func(timestamp int64)) (dbusutil.SignalHandlerId, error)
	ConnectBatteryAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectBatteryRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectLidClosed(cb func()) (dbusutil.SignalHandlerId, error)
	ConnectLidOpened(cb func()) (dbusutil.SignalHandlerId, error)
	ConnectPowerActionCode(cb func(actionCode int32)) (dbusutil.SignalHandlerId, error)
	PowerSavingModeAuto() proxy.PropBool
	OnBattery() proxy.PropBool
	HasLidSwitch() proxy.PropBool
	BatteryPercentage() proxy.PropDouble
	BatteryTimeToEmpty() proxy.PropUint64
	HasBattery() proxy.PropBool
	BatteryStatus() proxy.PropUint32
	BatteryTimeToFull() proxy.PropUint64
	BatteryCapacity() proxy.PropDouble
	PowerSavingModeEnabled() proxy.PropBool
	PowerSavingModeAutoWhenBatteryLow() proxy.PropBool
	PowerSavingModeBrightnessDropPercent() proxy.PropUint32
	PowerSavingModeBrightnessData() proxy.PropString
	CpuGovernor() proxy.PropString
	CpuBoost() proxy.PropBool
	IsBoostSupported() proxy.PropBool
	IsHighPerformanceSupported() proxy.PropBool
	IsBalanceSupported() proxy.PropBool
	IsPowerSaveSupported() proxy.PropBool
	Mode() proxy.PropString
}

type interfacePower struct{}

func (v *interfacePower) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfacePower) GetInterfaceName_() string {
	return "com.deepin.system.Power"
}

// method GetBatteries

func (v *interfacePower) GoGetBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBatteries", flags, ch)
}

func (*interfacePower) StoreGetBatteries(call *dbus.Call) (batteries []dbus.ObjectPath, err error) {
	err = call.Store(&batteries)
	return
}

func (v *interfacePower) GetBatteries(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetBatteries(
		<-v.GoGetBatteries(flags, make(chan *dbus.Call, 1)).Done)
}

// method Refresh

func (v *interfacePower) GoRefresh(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Refresh", flags, ch)
}

func (v *interfacePower) Refresh(flags dbus.Flags) error {
	return (<-v.GoRefresh(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RefreshBatteries

func (v *interfacePower) GoRefreshBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshBatteries", flags, ch)
}

func (v *interfacePower) RefreshBatteries(flags dbus.Flags) error {
	return (<-v.GoRefreshBatteries(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RefreshMains

func (v *interfacePower) GoRefreshMains(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshMains", flags, ch)
}

func (v *interfacePower) RefreshMains(flags dbus.Flags) error {
	return (<-v.GoRefreshMains(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetCpuGovernor

func (v *interfacePower) GoSetCpuGovernor(flags dbus.Flags, ch chan *dbus.Call, governor string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCpuGovernor", flags, ch, governor)
}

func (v *interfacePower) SetCpuGovernor(flags dbus.Flags, governor string) error {
	return (<-v.GoSetCpuGovernor(flags, make(chan *dbus.Call, 1), governor).Done).Err
}

// method SetCpuBoost

func (v *interfacePower) GoSetCpuBoost(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCpuBoost", flags, ch, enabled)
}

func (v *interfacePower) SetCpuBoost(flags dbus.Flags, enabled bool) error {
	return (<-v.GoSetCpuBoost(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method SetMode

func (v *interfacePower) GoSetMode(flags dbus.Flags, ch chan *dbus.Call, mode string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMode", flags, ch, mode)
}

func (v *interfacePower) SetMode(flags dbus.Flags, mode string) error {
	return (<-v.GoSetMode(flags, make(chan *dbus.Call, 1), mode).Done).Err
}

// method LockCpuFreq

func (v *interfacePower) GoLockCpuFreq(flags dbus.Flags, ch chan *dbus.Call, governor string, lockTime int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LockCpuFreq", flags, ch, governor, lockTime)
}

func (v *interfacePower) LockCpuFreq(flags dbus.Flags, governor string, lockTime int32) error {
	return (<-v.GoLockCpuFreq(flags, make(chan *dbus.Call, 1), governor, lockTime).Done).Err
}

// signal BatteryDisplayUpdate

func (v *interfacePower) ConnectBatteryDisplayUpdate(cb func(timestamp int64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BatteryDisplayUpdate", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BatteryDisplayUpdate",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var timestamp int64
		err := dbus.Store(sig.Body, &timestamp)
		if err == nil {
			cb(timestamp)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BatteryAdded

func (v *interfacePower) ConnectBatteryAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BatteryAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BatteryAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BatteryRemoved

func (v *interfacePower) ConnectBatteryRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BatteryRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BatteryRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal LidClosed

func (v *interfacePower) ConnectLidClosed(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LidClosed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LidClosed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal LidOpened

func (v *interfacePower) ConnectLidOpened(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LidOpened", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LidOpened",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PowerActionCode

func (v *interfacePower) ConnectPowerActionCode(cb func(actionCode int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PowerActionCode", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PowerActionCode",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var actionCode int32
		err := dbus.Store(sig.Body, &actionCode)
		if err == nil {
			cb(actionCode)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property PowerSavingModeAuto b

func (v *interfacePower) PowerSavingModeAuto() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "PowerSavingModeAuto",
	}
}

// property OnBattery b

func (v *interfacePower) OnBattery() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "OnBattery",
	}
}

// property HasLidSwitch b

func (v *interfacePower) HasLidSwitch() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasLidSwitch",
	}
}

// property BatteryPercentage d

func (v *interfacePower) BatteryPercentage() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "BatteryPercentage",
	}
}

// property BatteryTimeToEmpty t

func (v *interfacePower) BatteryTimeToEmpty() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "BatteryTimeToEmpty",
	}
}

// property HasBattery b

func (v *interfacePower) HasBattery() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasBattery",
	}
}

// property BatteryStatus u

func (v *interfacePower) BatteryStatus() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "BatteryStatus",
	}
}

// property BatteryTimeToFull t

func (v *interfacePower) BatteryTimeToFull() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "BatteryTimeToFull",
	}
}

// property BatteryCapacity d

func (v *interfacePower) BatteryCapacity() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "BatteryCapacity",
	}
}

// property PowerSavingModeEnabled b

func (v *interfacePower) PowerSavingModeEnabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "PowerSavingModeEnabled",
	}
}

// property PowerSavingModeAutoWhenBatteryLow b

func (v *interfacePower) PowerSavingModeAutoWhenBatteryLow() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "PowerSavingModeAutoWhenBatteryLow",
	}
}

// property PowerSavingModeBrightnessDropPercent u

func (v *interfacePower) PowerSavingModeBrightnessDropPercent() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "PowerSavingModeBrightnessDropPercent",
	}
}

// property PowerSavingModeBrightnessData s

func (v *interfacePower) PowerSavingModeBrightnessData() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "PowerSavingModeBrightnessData",
	}
}

// property CpuGovernor s

func (v *interfacePower) CpuGovernor() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "CpuGovernor",
	}
}

// property CpuBoost b

func (v *interfacePower) CpuBoost() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "CpuBoost",
	}
}

// property IsBoostSupported b

func (v *interfacePower) IsBoostSupported() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "IsBoostSupported",
	}
}

// property IsHighPerformanceSupported b

func (v *interfacePower) IsHighPerformanceSupported() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "IsHighPerformanceSupported",
	}
}

// property IsBalanceSupported b

func (v *interfacePower) IsBalanceSupported() proxy.PropBool {
    return &proxy.ImplPropBool{
        Impl: v,
        Name: "IsBalanceSupported",
    }
}

// property IsPowerSaveSupported b

func (v *interfacePower) IsPowerSaveSupported() proxy.PropBool {
    return &proxy.ImplPropBool{
        Impl: v,
        Name: "IsPowerSaveSupported",
    }
}

// property Mode s

func (v *interfacePower) Mode() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Mode",
	}
}

type Battery interface {
	battery // interface com.deepin.system.Power.Battery
	proxy.Object
}

type objectBattery struct {
	interfaceBattery // interface com.deepin.system.Power.Battery
	proxy.ImplObject
}

func NewBattery(conn *dbus.Conn, path dbus.ObjectPath) (Battery, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectBattery)
	obj.ImplObject.Init_(conn, "com.deepin.system.Power", path)
	return obj, nil
}

type battery interface {
	EnergyFullDesign() proxy.PropDouble
	Capacity() proxy.PropDouble
	Technology() proxy.PropString
	Energy() proxy.PropDouble
	EnergyFull() proxy.PropDouble
	Manufacturer() proxy.PropString
	ModelName() proxy.PropString
	TimeToEmpty() proxy.PropUint64
	IsPresent() proxy.PropBool
	Status() proxy.PropUint32
	EnergyRate() proxy.PropDouble
	Voltage() proxy.PropDouble
	Percentage() proxy.PropDouble
	TimeToFull() proxy.PropUint64
	UpdateTime() proxy.PropInt64
	SysfsPath() proxy.PropString
	SerialNumber() proxy.PropString
	Name() proxy.PropString
}

type interfaceBattery struct{}

func (v *interfaceBattery) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceBattery) GetInterfaceName_() string {
	return "com.deepin.system.Power.Battery"
}

// property EnergyFullDesign d

func (v *interfaceBattery) EnergyFullDesign() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "EnergyFullDesign",
	}
}

// property Capacity d

func (v *interfaceBattery) Capacity() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Capacity",
	}
}

// property Technology s

func (v *interfaceBattery) Technology() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Technology",
	}
}

// property Energy d

func (v *interfaceBattery) Energy() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Energy",
	}
}

// property EnergyFull d

func (v *interfaceBattery) EnergyFull() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "EnergyFull",
	}
}

// property Manufacturer s

func (v *interfaceBattery) Manufacturer() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Manufacturer",
	}
}

// property ModelName s

func (v *interfaceBattery) ModelName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "ModelName",
	}
}

// property TimeToEmpty t

func (v *interfaceBattery) TimeToEmpty() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "TimeToEmpty",
	}
}

// property IsPresent b

func (v *interfaceBattery) IsPresent() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "IsPresent",
	}
}

// property Status u

func (v *interfaceBattery) Status() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "Status",
	}
}

// property EnergyRate d

func (v *interfaceBattery) EnergyRate() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "EnergyRate",
	}
}

// property Voltage d

func (v *interfaceBattery) Voltage() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Voltage",
	}
}

// property Percentage d

func (v *interfaceBattery) Percentage() proxy.PropDouble {
	return &proxy.ImplPropDouble{
		Impl: v,
		Name: "Percentage",
	}
}

// property TimeToFull t

func (v *interfaceBattery) TimeToFull() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "TimeToFull",
	}
}

// property UpdateTime x

func (v *interfaceBattery) UpdateTime() proxy.PropInt64 {
	return &proxy.ImplPropInt64{
		Impl: v,
		Name: "UpdateTime",
	}
}

// property SysfsPath s

func (v *interfaceBattery) SysfsPath() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "SysfsPath",
	}
}

// property SerialNumber s

func (v *interfaceBattery) SerialNumber() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "SerialNumber",
	}
}

// property Name s

func (v *interfaceBattery) Name() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Name",
	}
}
