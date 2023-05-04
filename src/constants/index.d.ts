// Payload from backend
/**
 * id: message unique id
 * text: user sent text
 * response: bot response
 * sentTime: user sent time
 * historyId: message related history
 */
export interface Message {
  id: number;
  text: string;
  response: string;
  sent: boolean;
  sentTime: string;
  historyId: number;
  historyTimestamp: string;
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
  id: number;
  text: string;
  response: string;
  sentTime: string;
  historyId: number;
  historyTimestamp: string;
  method: 'KMP' | 'BoyerMoore' | 'GPT';
}

export interface MessageInterface extends Message {
  responseStatusCode: number;
  responded: boolean;

  getId(): number;
  getStatus(): boolean;
  getText(): string;
  getResponseMsg(): string;
  getResponseCode(): number;
  getSentTime(): string;
  getResponseStatus(): boolean;
  getHistoryTimestamp(): string;
  getHistoryId(): number;
  setResponseStatus(status: boolean): void;
  setResponse(text: string, statusCode: number): void;
  setHistoryTimestamp(timestamp: string): void;
}

export interface HistoryRequest {
  historyId: number;
  historyTopic: string;
}

export interface HistoryPayload {
  historyCollection: Array<HistoryRequest>;
}

export interface FetchedHistoryTopic {
  HistoryID: number;
  HistoryTopic: string;
}
