package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wizlif/tempmee_assignment/util"
)

func TestListOrders(t *testing.T) {
	limit := 5
	var createdOrders []VOrder
	for i := 0; i < limit*2; i++ {
		createdOrders = append(createdOrders, createRandomOrder(t))
	}

	arg := ListOrdersParams{
		Limit:  int64(limit),
		Offset: 0,
	}

	orders, err := testStore.ListOrders(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.Equal(t, limit, len(orders))

	for index, order := range orders {
		require.NotEmpty(t, order)

		createdOrder := createdOrders[index]
		require.Equal(t, createdOrder.Price, order.Price)
		require.Equal(t, createdOrder.Total, order.Total)
		require.Equal(t, createdOrder.Status, order.Status)
	}
}
func TestListUserOrders(t *testing.T) {
	limit := 5
	user, _ := createRandomOrdersForUser(t, limit)

	arg := ListUserOrdersParams{
		Limit:  int64(limit),
		Offset: 0,
		Email:  user.Email,
	}

	orders, err := testStore.ListUserOrders(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.Equal(t, limit, len(orders))

	for _, order := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, user.Email, order.Email)
	}
}

func TestUpdateOrderStatus(t *testing.T) {
	order := createRandomOrder(t)

	newStatus := "COMPLETE"
	arg := UpdateOrderStatusParams{
		Status: newStatus,
		ID:     order.ID,
	}
	err := testStore.UpdateOrderStatus(context.Background(), arg)
	require.NoError(t, err)

	vOrder, err := testStore.GetOrderById(context.Background(), order.ID)
	require.NoError(t, err)
	require.NotEmpty(t, vOrder)
	require.Equal(t, vOrder.Status, newStatus)
}

func TestCreateRandomOrder(t *testing.T) {
	createRandomOrder(t)
}

func createRandomOrder(t *testing.T) VOrder {
	book := createRandomBook(t)
	user := createRandomUser(t)

	arg := CreateOrderParams{
		BookID: book.ID,
		UserID: user.ID,
		Price:  util.RandomFloat(30.0, 80.0),
		Total:  util.RandomInt(1, 10),
		Status: "PENDING",
	}

	orderId, err := testStore.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderId)

	order, err := testStore.GetOrderById(context.Background(), orderId)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, book.Title, order.Title)
	require.Equal(t, user.Email, order.Email)
	require.Equal(t, arg.Status, order.Status)
	require.Equal(t, arg.Total, order.Total)
	require.NotZero(t, order.CreatedAt)

	return order
}
func createRandomOrdersForUser(t *testing.T, limit int) (User, Book) {
	book := createRandomBook(t)
	user := createRandomUser(t)

	arg := CreateOrderParams{
		BookID: book.ID,
		UserID: user.ID,
		Price:  util.RandomFloat(30.0, 80.0),
		Total:  util.RandomInt(1, 10),
		Status: "PENDING",
	}

	for i := 0; i < limit; i++ {
		orderId, err := testStore.CreateOrder(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, orderId)
	}

	return user, book
}
