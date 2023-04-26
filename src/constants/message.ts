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

  constructor(id: number, sent: boolean, text: string, sentTime: string, historyId: number) {
    this.id = id;
    this.sent = sent;
    this.text = text;
    this.response = "";
    this.responseStatusCode = 0;
    this.sentTime = sentTime;
    this.responded = false;
    this.historyId = historyId;
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

  public setResponseStatus(status: boolean) {
    this.responded = status;
  }

  public setResponse(text: string, statusCode: number): void {
    this.response = text;
    this.responseStatusCode = statusCode;
  }
}
