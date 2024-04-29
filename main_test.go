package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	clientID := 1

	// напиши тест здесь
	client, err := selectClient(db, clientID)
	if err != nil {
		log.Println(err)
		return
	}

	require.Equal(t, client.ID, clientID)
	require.NotEmpty(t, client.FIO, client.Login, client.Birthday, client.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	clientID := -1

	// напиши тест здесь
	client, err := selectClient(db, clientID)
	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, client.ID, client.FIO, client.Login, client.Birthday, client.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	res, err := insertClient(db, cl)
	clientID := res
	require.NotEmpty(t, clientID)
	require.NoError(t, err)

	client, err := selectClient(db, clientID)
	require.NoError(t, err)
	require.Equal(t, client.FIO, cl.FIO)
	require.Equal(t, client.Login, cl.Login)
	require.Equal(t, client.Birthday, cl.Birthday)
	require.Equal(t, client.Email, cl.Email)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	res, err := insertClient(db, cl)
	require.NotEmpty(t, res)
	require.NoError(t, err)

	client, err := selectClient(db, res)
	require.NoError(t, err)

	err = deleteClient(db, client.ID)
	require.NoError(t, err)

	client, err = selectClient(db, client.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
}
