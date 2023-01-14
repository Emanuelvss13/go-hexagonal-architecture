package cli_test

import (
	"fmt"
	"testing"

	"github.com/emanuelvss13/go-hexagonal/adapters/cli"
	mock_application "github.com/emanuelvss13/go-hexagonal/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)
	ctrl.Finish()

	id := "test"
	name := "Product Test"
	price := 200.20
	status := "disabled"

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)

	serviceMock.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Create(name, price).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", id, name, price, status)

	result, err := cli.Run(serviceMock, "create", "", name, price)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", name)

	result, err = cli.Run(serviceMock, "enable", id, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", name)

	result, err = cli.Run(serviceMock, "disable", id, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", id, name, price, status)

	result, err = cli.Run(serviceMock, "", id, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
