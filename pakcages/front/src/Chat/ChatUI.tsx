import React, { useEffect, useMemo, useState } from 'react';
// @ts-expect-error
import ChatBox from 'react-chat-plugin';
import './Chat.css';
import { messagesData, rooms } from "./data";

export function ChatUI() {

  const [rommId, setRoomId] = useState("");
  const [messages, setMessages] = useState(messagesData[rommId]);

  const selectRoom = (roomId: string) => {
    setMessages(messagesData[roomId])
    setRoomId(roomId)
  }

  const isRoomSelected = useMemo(() => {
    return rommId !== ""
  }, [rommId]);

  const ChatRooms = rooms.map((room) =>
    <div
      key={room.id}
      className="Chat_Left_Item"
      onClick={() => { selectRoom(room.id) }}
    >
      {room.id}
    </div>
  );

  const handleOnSendMessage = (message: any) => {
    const newMessage = {
      author: {
        username: 'user1',
        id: 1,
        avatarUrl: 'https://image.flaticon.com/icons/svg/2446/2446032.svg',
      },
      text: message,
      type: 'text',
      timestamp: +new Date(),
    };
    const newMessages = messages.concat(newMessage)
    setMessages(newMessages);
    messagesData[rommId] = newMessages;
  };

  return (
    <div className="Chat_Container">
      <div
        className="Chat_Left"
      >
        {
          ChatRooms
        }
      </div>
      <div>
        <ChatBox
          onSendMessage={handleOnSendMessage}
          userId={1}
          messages={messages}
          showTypingIndicator={false}
          disableInput={!isRoomSelected}
          activeAuthor={{ username: 'user2', id: 2, avatarUrl: null }}
        />
      </div>
    </div>
  );
}
