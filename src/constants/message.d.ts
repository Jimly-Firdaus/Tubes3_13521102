// Payload from backend
export interface Message {
  id: number;
  text: string;
  response: string;
  responseStatusCode: number;
  sentTime: string;
}

export interface History {
  historyId: number;
  topic: string;
  conversation: Array<Message>;
}

// Payload from backend
export interface MessageHistory {
  messageHistory: Array<History>;
}

// Request sent to backend on user-entered
export interface Request {
  message: Message;
  method: "KMP" | "BoyerMoore";
  requestModify?: boolean;
}

// Future use --ignore this
export interface userData {
  userId: number;
  username?: string;
  messageHistory: MessageHistory;
}

export interface UserMessageHistory extends MessageHistory {
  userId: number;
}
