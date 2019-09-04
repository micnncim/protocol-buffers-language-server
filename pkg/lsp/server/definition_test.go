package server

import (
	"context"
	"testing"
	"text/scanner"

	protobuf "github.com/emicklei/proto"
	"github.com/go-language-server/protocol"
	"github.com/go-language-server/uri"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry/registrytest"
)

func TestDefinition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cases := []struct {
		name       string
		protoSet   registry.ProtoSet
		params     *protocol.TextDocumentPositionParams
		wantResult []protocol.Location
		wantErr    bool
	}{
		{
			name: "success",
			protoSet: func() registry.ProtoSet {
				protoSet := registrytest.NewMockProtoSet(ctrl)
				protoSet.
					EXPECT().
					GetProtoByFilename("proto/test.proto").
					Return(func() registry.Proto {
						proto := registrytest.NewMockProto(ctrl)
						proto.EXPECT().
							GetMessageFieldByLine(10).
							Return(&registry.MessageField{
								ProtoField: &protobuf.NormalField{
									Field: &protobuf.Field{
										Type: "Message",
									},
								},
							}, true)
						proto.EXPECT().
							GetMessageByName("Message").
							Return(func() registry.Message {
								message := registrytest.NewMockMessage(ctrl)
								message.EXPECT().
									Protobuf().
									Return(&protobuf.Message{
										Position: scanner.Position{
											Line:   14,
											Column: 1,
										},
									})
								return message
							}(), true)
						return proto
					}())

				return protoSet
			}(),
			params: &protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{
					URI: uri.File("proto/test.proto"),
				},
				Position: protocol.Position{
					Line:      10,
					Character: 6,
				},
			},
			wantResult: []protocol.Location{
				{
					URI: uri.File("proto/test.proto"),
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      14,
							Character: 1,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := &Server{
				protoSet: tc.protoSet,
				logger:   zap.NewNop(),
			}
			gotResult, err := s.definition(context.Background(), tc.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("Server.definition() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if diff := cmp.Diff(gotResult, tc.wantResult); diff != "" {
				t.Errorf("Server.definition: (-got, +want)\n%s", diff)
			}
		})
	}
}
