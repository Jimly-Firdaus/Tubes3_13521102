export class Message {
  private id: number;
  private sent: boolean;
  private text: string;
  private response: string;
  private responseStatusCode: number;
  private sentTime: string;
  private responded: boolean;

  constructor(id: number, sent: boolean, text: string, sentTime: string) {
    this.id = id;
    this.sent = sent;
    this.text = text;
    this.response = "";
    this.responseStatusCode = 0;
    this.sentTime = sentTime;
    this.responded = false;
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
