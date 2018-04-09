// gomuks - A terminal Matrix client written in Go.
// Copyright (C) 2018 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package ifc

import (
	"time"

	"github.com/gdamore/tcell"
	"maunium.net/go/gomatrix"
	"maunium.net/go/gomuks/matrix/pushrules"
	"maunium.net/go/gomuks/matrix/rooms"
	"maunium.net/go/tview"
)

type View string

// Allowed views in GomuksUI
const (
	ViewLogin View = "login"
	ViewMain  View = "main"
)

type GomuksUI interface {
	Render()
	SetView(name View)
	InitViews() tview.Primitive
	MainView() MainView
	LoginView() LoginView
}

type MainView interface {
	GetRoom(roomID string) RoomView
	HasRoom(roomID string) bool
	AddRoom(roomID string)
	RemoveRoom(roomID string)
	SetRooms(roomIDs []string)
	SaveAllHistory()

	SetTyping(roomID string, users []string)
	ProcessMessageEvent(roomView RoomView, evt *gomatrix.Event) Message
	ProcessMembershipEvent(roomView RoomView, evt *gomatrix.Event) Message
	NotifyMessage(room *rooms.Room, message Message, should pushrules.PushActionArrayShould)
}

type LoginView interface {
}

type MessageDirection int

const (
	AppendMessage  MessageDirection = iota
	PrependMessage
	IgnoreMessage
)

type RoomView interface {
	MxRoom() *rooms.Room
	SaveHistory(dir string) error
	LoadHistory(dir string) (int, error)

	SetStatus(status string)
	SetTyping(users []string)
	UpdateUserList()

	NewMessage(id, sender, msgtype, text string, timestamp time.Time) Message
	NewTempMessage(msgtype, text string) Message
	AddMessage(message Message, direction MessageDirection)
	AddServiceMessage(message string)
}

type MessageMeta interface {
	Sender() string
	SenderColor() tcell.Color
	TextColor() tcell.Color
	TimestampColor() tcell.Color
	Timestamp() string
	Date() string
	CopyFrom(from MessageMeta)
}

// MessageState is an enum to specify if a Message is being sent, failed to send or was successfully sent.
type MessageState int

// Allowed MessageStates.
const (
	MessageStateSending MessageState = iota
	MessageStateDefault
	MessageStateFailed
)

type Message interface {
	MessageMeta

	SetIsHighlight(isHighlight bool)
	IsHighlight() bool

	SetIsService(isService bool)
	IsService() bool

	SetID(id string)
	ID() string

	SetType(msgtype string)
	Type() string

	SetText(text string)
	Text() string

	SetState(state MessageState)
	State() MessageState
}
