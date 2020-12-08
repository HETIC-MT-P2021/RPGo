// Code generated by MockGen. DO NOT EDIT.
// Source: commands/discord.go

// Package commands is a generated GoMock package.
package commands

import (
	discordgo "github.com/bwmarrin/discordgo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDiscordConnector is a mock of DiscordConnector interface
type MockDiscordConnector struct {
	ctrl     *gomock.Controller
	recorder *MockDiscordConnectorMockRecorder
}

// MockDiscordConnectorMockRecorder is the mock recorder for MockDiscordConnector
type MockDiscordConnectorMockRecorder struct {
	mock *MockDiscordConnector
}

// NewMockDiscordConnector creates a new mock instance
func NewMockDiscordConnector(ctrl *gomock.Controller) *MockDiscordConnector {
	mock := &MockDiscordConnector{ctrl: ctrl}
	mock.recorder = &MockDiscordConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscordConnector) EXPECT() *MockDiscordConnectorMockRecorder {
	return m.recorder
}

// ChannelMessageSend mocks base method
func (m *MockDiscordConnector) ChannelMessageSend(arg0, arg1 string) (*discordgo.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelMessageSend", arg0, arg1)
	ret0, _ := ret[0].(*discordgo.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChannelMessageSend indicates an expected call of ChannelMessageSend
func (mr *MockDiscordConnectorMockRecorder) ChannelMessageSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelMessageSend", reflect.TypeOf((*MockDiscordConnector)(nil).ChannelMessageSend), arg0, arg1)
}

// ChannelMessageSendEmbed mocks base method
func (m *MockDiscordConnector) ChannelMessageSendEmbed(arg0 string, arg1 *discordgo.MessageEmbed) (*discordgo.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelMessageSendEmbed", arg0, arg1)
	ret0, _ := ret[0].(*discordgo.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChannelMessageSendEmbed indicates an expected call of ChannelMessageSendEmbed
func (mr *MockDiscordConnectorMockRecorder) ChannelMessageSendEmbed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelMessageSendEmbed", reflect.TypeOf((*MockDiscordConnector)(nil).ChannelMessageSendEmbed), arg0, arg1)
}
