import { MessageInterface } from "src/constants/index";

export class Message implements MessageInterface {
  id: number;
  sent: boolean;
  text: string;
  response: string;
  responseStatusCode: number;
  sentTime: string;
  responded: boolean;
  // This history id must be determined by fetching backend total history
  historyId: number;
  // This is filled when this chat marks a new conversation
  historyTimestamp: string;

  constructor(
    id: number,
    sent: boolean,
    text: string,
    sentTime: string,
    historyId: number
  ) {
    this.id = id;
    this.sent = sent;
    this.text = text;
    this.response = "";
    this.responseStatusCode = 0;
    this.sentTime = sentTime;
    this.responded = false;
    this.historyId = historyId;
    this.historyTimestamp = "";
  }

  public getId(): number {
    return this.id;
  }

  public getStatus(): boolean {
    return this.sent;
  }

  public getText(): string {
    return this.text;
  }

  public getResponseMsg(): string {
    return this.response;
  }

  public getResponseCode(): number {
    return this.responseStatusCode;
  }

  public getSentTime(): string {
    return this.sentTime;
  }

  public getResponseStatus(): boolean {
    return this.responded;
  }

  public getHistoryId(): number {
    return this.historyId;
  }

  public setResponseStatus(status: boolean): void {
    this.responded = status;
  }

  public setHistoryTimestamp(timestamp: string): void {
    this.historyTimestamp = timestamp;
  }

  public getHistoryTimestamp(): string {
    return this.historyTimestamp;
  }

  public setResponse(text: string, statusCode: number): void {
    this.response = text;
    this.responseStatusCode = statusCode;
  }
}
