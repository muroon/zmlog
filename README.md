# zmlog (Zap MarshalLogObject Generator)

`zmlog` generates the file which has MarshalLogObject method from struct.

## Install

```shell
go install github.com/muroon/zmlog/cmd/zmlog@latest
```

## How to generate file

For example, there is a file `target.go` that describes the struct for Log as follows.

```target.go
type TargetLog struct {
	ID   int
	Name string
	Time time.Time
}
```
Run the command below,

```shell
zmlog -f target.go
```

And `zmlog` generates a file `target_zap_obj.go` with MarshalLogObject method.

```target_zap_obj.go
// MarshalLogObject zapcore.ObjectMarshaler interface method
func (l *TargetLog) MarshalLogObject(enc zapcore.ObjectEncoder) error {
		var err error
		enc.AddInt("id", l.ID)
		enc.AddString("name", l.Name)
		enc.AddTime("time", l.Time)
		return err	
}
```

## Key

The key name for `zapcore.ObjectEncoder`'s method like AddInt or AddString... is basically the field name of the target struct with snake case.
However, as an exception, create a tag with `key: "key name"` in the field of the struct.

In the example below, it is originally `enc.AddString("sessison_id", l.SessionID)`, but it becomes `enc.AddString("sess_id", l.SessionID)`

```target.go
type TargetLog struct {
	ID        int
	Name      string
	Time      time.Time
	SessionID string `key:"sess_id"` // custom key
}
```
```target_zap_obj.go
func (l *TargetLog) MarshalLogObject(enc zapcore.ObjectEncoder) error {
		var err error
		enc.AddInt("id", l.ID)
		enc.AddString("name", l.Name)
		enc.AddTime("time", l.Time)
		enc.AddString("sess_id", l.SessionID) // not be `sessison_id`
		return err	
}
```
