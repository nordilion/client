// Auto-generated by avdl-compiler v1.3.20 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/metadata_update.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type RekeyRequest struct {
	FolderID string `codec:"folderID" json:"folderID"`
	Revision int64  `codec:"revision" json:"revision"`
}

func (o RekeyRequest) DeepCopy() RekeyRequest {
	return RekeyRequest{
		FolderID: o.FolderID,
		Revision: o.Revision,
	}
}

type MetadataUpdateArg struct {
	FolderID string `codec:"folderID" json:"folderID"`
	Revision int64  `codec:"revision" json:"revision"`
}

func (o MetadataUpdateArg) DeepCopy() MetadataUpdateArg {
	return MetadataUpdateArg{
		FolderID: o.FolderID,
		Revision: o.Revision,
	}
}

type FolderNeedsRekeyArg struct {
	FolderID string `codec:"folderID" json:"folderID"`
	Revision int64  `codec:"revision" json:"revision"`
}

func (o FolderNeedsRekeyArg) DeepCopy() FolderNeedsRekeyArg {
	return FolderNeedsRekeyArg{
		FolderID: o.FolderID,
		Revision: o.Revision,
	}
}

type FoldersNeedRekeyArg struct {
	Requests []RekeyRequest `codec:"requests" json:"requests"`
}

func (o FoldersNeedRekeyArg) DeepCopy() FoldersNeedRekeyArg {
	return FoldersNeedRekeyArg{
		Requests: (func(x []RekeyRequest) []RekeyRequest {
			if x == nil {
				return nil
			}
			var ret []RekeyRequest
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Requests),
	}
}

type MetadataUpdateInterface interface {
	MetadataUpdate(context.Context, MetadataUpdateArg) error
	FolderNeedsRekey(context.Context, FolderNeedsRekeyArg) error
	FoldersNeedRekey(context.Context, []RekeyRequest) error
}

func MetadataUpdateProtocol(i MetadataUpdateInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.metadataUpdate",
		Methods: map[string]rpc.ServeHandlerDescription{
			"metadataUpdate": {
				MakeArg: func() interface{} {
					ret := make([]MetadataUpdateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]MetadataUpdateArg)
					if !ok {
						err = rpc.NewTypeError((*[]MetadataUpdateArg)(nil), args)
						return
					}
					err = i.MetadataUpdate(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"folderNeedsRekey": {
				MakeArg: func() interface{} {
					ret := make([]FolderNeedsRekeyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FolderNeedsRekeyArg)
					if !ok {
						err = rpc.NewTypeError((*[]FolderNeedsRekeyArg)(nil), args)
						return
					}
					err = i.FolderNeedsRekey(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"foldersNeedRekey": {
				MakeArg: func() interface{} {
					ret := make([]FoldersNeedRekeyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FoldersNeedRekeyArg)
					if !ok {
						err = rpc.NewTypeError((*[]FoldersNeedRekeyArg)(nil), args)
						return
					}
					err = i.FoldersNeedRekey(ctx, (*typedArgs)[0].Requests)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type MetadataUpdateClient struct {
	Cli rpc.GenericClient
}

func (c MetadataUpdateClient) MetadataUpdate(ctx context.Context, __arg MetadataUpdateArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.metadataUpdate.metadataUpdate", []interface{}{__arg}, nil)
	return
}

func (c MetadataUpdateClient) FolderNeedsRekey(ctx context.Context, __arg FolderNeedsRekeyArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.metadataUpdate.folderNeedsRekey", []interface{}{__arg}, nil)
	return
}

func (c MetadataUpdateClient) FoldersNeedRekey(ctx context.Context, requests []RekeyRequest) (err error) {
	__arg := FoldersNeedRekeyArg{Requests: requests}
	err = c.Cli.Call(ctx, "keybase.1.metadataUpdate.foldersNeedRekey", []interface{}{__arg}, nil)
	return
}
