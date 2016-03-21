// Auto-generated by avdl-compiler v1.2.0 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/ctl.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type ExitCode int

const (
	ExitCode_OK      ExitCode = 0
	ExitCode_NOTOK   ExitCode = 2
	ExitCode_RESTART ExitCode = 4
)

type StopArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	ExitCode  ExitCode `codec:"exitCode" json:"exitCode"`
}

type LogRotateArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ReloadArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type DbNukeArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type CtlInterface interface {
	Stop(context.Context, StopArg) error
	LogRotate(context.Context, int) error
	Reload(context.Context, int) error
	DbNuke(context.Context, int) error
}

func CtlProtocol(i CtlInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.ctl",
		Methods: map[string]rpc.ServeHandlerDescription{
			"stop": {
				MakeArg: func() interface{} {
					ret := make([]StopArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]StopArg)
					if !ok {
						err = rpc.NewTypeError((*[]StopArg)(nil), args)
						return
					}
					err = i.Stop(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"logRotate": {
				MakeArg: func() interface{} {
					ret := make([]LogRotateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LogRotateArg)
					if !ok {
						err = rpc.NewTypeError((*[]LogRotateArg)(nil), args)
						return
					}
					err = i.LogRotate(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"reload": {
				MakeArg: func() interface{} {
					ret := make([]ReloadArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ReloadArg)
					if !ok {
						err = rpc.NewTypeError((*[]ReloadArg)(nil), args)
						return
					}
					err = i.Reload(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"dbNuke": {
				MakeArg: func() interface{} {
					ret := make([]DbNukeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DbNukeArg)
					if !ok {
						err = rpc.NewTypeError((*[]DbNukeArg)(nil), args)
						return
					}
					err = i.DbNuke(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type CtlClient struct {
	Cli rpc.GenericClient
}

func (c CtlClient) Stop(ctx context.Context, __arg StopArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.ctl.stop", []interface{}{__arg}, nil)
	return
}

func (c CtlClient) LogRotate(ctx context.Context, sessionID int) (err error) {
	__arg := LogRotateArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.ctl.logRotate", []interface{}{__arg}, nil)
	return
}

func (c CtlClient) Reload(ctx context.Context, sessionID int) (err error) {
	__arg := ReloadArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.ctl.reload", []interface{}{__arg}, nil)
	return
}

func (c CtlClient) DbNuke(ctx context.Context, sessionID int) (err error) {
	__arg := DbNukeArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.ctl.dbNuke", []interface{}{__arg}, nil)
	return
}
