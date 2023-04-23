// Payload from backend
/**
 * id: message unique id
 * text: user sent text
 * reponse: bot response
 * responseStatusCode for network checking
 * sentTime: user sent time
*/
export interface Message {
  id: number;
  text: string;
  response: string;
  responseStatusCode: number;
  sentTime: string;
}

/**
 * historyId: history unique id
 * topic: this history topic
 * conversation: array of user message for this topic
*/
export interface History {
  historyId: number;
  topic: string;
  conversation: Array<Message>;
}

// Payload from backend
/**
 * messageHistory: all user chat history
*/
export interface MessageHistory {
  messageHistory: Array<History>;
}

// Request sent to backend on user-entered
/**
 * message: current user message
 * method: string match type
 * requestModify: optional, true if user wants to add/remove/change db content, otherwise false
*/
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
