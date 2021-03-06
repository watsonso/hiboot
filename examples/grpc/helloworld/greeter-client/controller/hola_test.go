package controller

import (
	"github.com/golang/mock/gomock"
	"github.com/hidevopsio/hiboot/examples/grpc/helloworld/mock"
	"github.com/hidevopsio/hiboot/examples/grpc/helloworld/protobuf"
	"github.com/hidevopsio/hiboot/pkg/app"
	"github.com/hidevopsio/hiboot/pkg/app/web"
	grpcmock "github.com/hidevopsio/hiboot/pkg/starter/grpc/mock"
	"net/http"
	"testing"
)

func TestHolaClient(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHolaClient := mock.NewMockHolaServiceClient(ctrl)
	app.Register("protobuf.holaServiceClient", mockHolaClient)

	testApp := web.NewTestApplication(t, newHolaController)

	req := &protobuf.HolaRequest{Name: "Steve"}

	mockHolaClient.EXPECT().SayHola(
		gomock.Any(),
		&grpcmock.RPCMsg{Message: req},
	).Return(&protobuf.HolaReply{Message: "Hola " + req.Name}, nil)

	testApp.Get("/hola/name/{name}").
		WithPath("name", req.Name).
		Expect().Status(http.StatusOK).
		Body().Contains(req.Name)
}
