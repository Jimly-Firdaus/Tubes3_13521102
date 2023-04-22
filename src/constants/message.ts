export class Message {
  private id: number;
  private sent: boolean;
  private text: string;
  private response: string;
  private responseStatusCode: number;

  constructor(id: number, sent: boolean, text: string) {
    this.id = id;
    this.sent = sent;
    this.text = text;
    this.response = '';
    this.responseStatusCode = 0;
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

  public setResponse(text: string, statusCode: number) {
    this.response = text;
    this.responseStatusCode = statusCode;
  }
}
