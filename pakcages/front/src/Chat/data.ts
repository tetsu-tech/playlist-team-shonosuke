export const companies = [
  {
    id: "company",
    name: "gifts japan",
  }
]

export const rooms = [
  {
    id: "room1",
  },
  {
    id: "room2"
  }
]

interface IMessages {
  [key: string]: any[],
}

export const messagesData: IMessages = {
  "room1": [
    {
      text: 'user2 has joined the conversation',
      timestamp: 1578366389250,
      type: 'notification',
    },
    {
      author: {
        username: 'user1',
        id: 1,
        avatarUrl: 'https://image.flaticon.com/icons/svg/2446/2446032.svg',
      },
      text: 'Hi!!!!!!!',
      type: 'text',
      timestamp: 1578366393250,
    },
  ],
  "room2": [
    {
      text: 'user2 has joined the conversation',
      timestamp: 1578366389250,
      type: 'notification',
    },
    {
      author: {
        username: 'user1',
        id: 1,
        avatarUrl: 'https://image.flaticon.com/icons/svg/2446/2446032.svg',
      },
      text: 'Hi!!!!!!!',
      type: 'text',
      timestamp: 1578366393250,
    },
    {
      author: {
        username: 'user1',
        id: 1,
        avatarUrl: 'https://image.flaticon.com/icons/svg/2446/2446032.svg',
      },
      text: 'Hi!!!!!!!',
      type: 'text',
      timestamp: 1578366393250,
    },
    {
      author: {
        username: 'user1',
        id: 1,
        avatarUrl: 'https://image.flaticon.com/icons/svg/2446/2446032.svg',
      },
      text: 'Hi!!!!!!!',
      type: 'text',
      timestamp: 1578366393250,
    },
  ]
}
