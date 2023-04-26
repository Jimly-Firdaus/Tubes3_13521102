// Payload from backend
/**
 * id: message unique id
 * text: user sent text
 * sentTime: user sent time
 */
export interface Message {
  id: number;
  text: string;
  sentTime: string;
  historyId: number;
}

/**
 * historyId: history unique id
 * topic: this history topic
 * conversation: array of user message for this topic
 */
export interface History {
  historyId: number;
  topic: string;
  conversation: Array<MessageInterface>;
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
  messageData: Message;
  method: 'KMP' | 'BoyerMoore';
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

export interface MessageInterface extends Message {
  response: string;
  responseStatusCode: number;

  getId(): number;
  getStatus(): boolean;
  getText(): string;
  getResponseMsg(): string;
  getResponseCode(): number;
  getSentTime(): string;
  getResponseStatus(): boolean;
  setResponseStatus(status: boolean): void;
  setResponse(text: string, statusCode: number): void;
}
